package mq

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"sctek.com/typhoon/th-platform-gateway/common"
	"time"
)

//var url = flag.String("url", "amqp:///", "AMQP url for both the publisher and subscriber")

// exchange binds the publishers to the subscribers

// message is the application type for a message.  This can contain identity,
// or a reference to the recevier chan for further demuxing.
var ChanMessage = make(chan []byte, 100)

// session composes an amqp.Connection with an amqp.Channel
type session struct {
	*amqp.Connection
	*amqp.Channel
}

// Close tears the connection down, taking the channel with it.
func (s session) Close() error {
	if s.Connection == nil {
		return nil
	}
	return s.Connection.Close()
}

// redial continually connects to the URL, exiting the program when no longer possible
func redial(ctx context.Context, url, exchange string) chan chan session {
	sessions := make(chan chan session)

	go func() {
		sess := make(chan session)
		defer close(sessions)
		defer close(sess)
		for {
			select {
			case sessions <- sess:
			case <-ctx.Done():
				common.Log.Traceln("shutting down session factory")
				return
			}
			conn, err := amqp.Dial(url)
			if err != nil {
				errInfo:=fmt.Sprintf("cannot (re)dial: %v: %q", err, url)
				common.Log.Errorln(errInfo)
				log.Fatal(errInfo)
			}

			ch, err := conn.Channel()
			if err != nil {
				errInfo:=fmt.Sprintf("cannot create channel: %v", err)
				common.Log.Errorln(errInfo)
				log.Fatal(errInfo)
			}

			if err := ch.ExchangeDeclare(exchange, "fanout", true, false, false, false, nil); err != nil {
				errInfo:=fmt.Sprintf("cannot declare fanout exchange: %v", err)
				common.Log.Errorln(errInfo)
				log.Fatal(errInfo)
			}
			//common.Log.Tracef("%v %v",url,exchange)
			select {
			case sess <- session{conn, ch}:
			case <-ctx.Done():
				common.Log.Infoln("shutting down new session")
				return
			}
		}
	}()

	return sessions
}

// publish publishes messages to a reconnecting session to a fanout exchange.
// It receives from the application specific source of messages.
func publish(sessions chan chan session, exchange string, messages <-chan []byte) {
	var (
		running bool
		reading = messages
		pending = make(chan []byte, 1)
		confirm = make(chan amqp.Confirmation, 1)
	)
	defer close(confirm)
	defer close(pending)

	for session := range sessions {
		pub := <-session
		// publisher confirms for this channel/connection
		if err := pub.Confirm(false); err != nil {
			common.Log.Errorln("publisher confirms not supported")
			close(confirm) // confirms not supported, simulate by always nacking
		} else {
			pub.NotifyPublish(confirm)
		}

		//common.Log.Traceln("publishing...")
	Publish:
		for {
			var body []byte
			select {
			case confirmed, ok := <-confirm:
				if !ok {
					break Publish
				}
				if !confirmed.Ack {
					common.Log.Errorf("nack message %d, body: %q", confirmed.DeliveryTag, string(body))
				}
				reading = messages

			case body = <-pending:
				//routingKey := "ignored for fanout exchanges, application dependent for other exchanges"
				err := pub.Publish(exchange, "", false, false, amqp.Publishing{
					ContentType: "text/plain",
					Body:        body,
				})
				// Retry failed delivery on the next session
				if err != nil {
					pending <- body
					pub.Close()
					break Publish
				}

			case body, running = <-reading:
				// all messages consumed
				if !running {
					return
				}
				// work on pending delivery until ack'd
				pending <- body
				reading = nil
			}
		}
	}
}

// identity returns the same host/process unique string for the lifetime of
// this process so that subscriber reconnections reuse the same queue name.
func identity() string {
	hostname, err := os.Hostname()
	h := sha1.New()
	fmt.Fprint(h, hostname)
	fmt.Fprint(h, err)
	fmt.Fprint(h, os.Getpid())
	return fmt.Sprintf("%x", h.Sum(nil))
}

// subscribe consumes deliveries from an exclusive queue from a fanout exchange and sends to the application specific messages chan.
func subscribe(sessions chan chan session, queueName, exchange string, callBack func([]byte)) {

	for session := range sessions {
		sub := <-session

		if _, err := sub.QueueDeclare(queueName, true, false, false, false, nil); err != nil {
			str := fmt.Sprintf("cannot consume from exclusive queue: %q, %v", queueName, err)
			common.Log.Traceln(str)
			panic(str)
			return
		}

		//routingKey := "application specific routing key for fancy toplogies"
		if err := sub.QueueBind(queueName, "", exchange, false, nil); err != nil {
			str := fmt.Sprintf("cannot consume without a binding to exchange: %q, %v", exchange, err)
			common.Log.Traceln(str)
			panic(str)
			return
		}

		deliveries, err := sub.Consume(queueName, "", false, false, false, false, nil)
		if err != nil {
			str := fmt.Sprintf("cannot consume from: %q, %v", queueName, err)
			common.Log.Traceln(str)
			panic(str)
			return
		}

		common.Log.Tracef("  subscribed 【%v】...",queueName)
		for msg := range deliveries {
			callBack(msg.Body)
			sub.Ack(msg.DeliveryTag, false)
			time.Sleep(time.Nanosecond)
		}

	}
}



func (s *session) amQpDial(url, exchange string) error {
	conn, err := amqp.Dial(url)
	if err != nil {
		common.Log.Errorln("cannot (re)dial: %v: %q", err, url)
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		common.Log.Errorln("cannot create channel: %v", err)
		return err
	}

	if exchange!=""{
		if err := ch.ExchangeDeclare(exchange, "fanout", true, false, false, false, nil); err != nil {
			common.Log.Errorln("cannot declare fanout exchange: %v", err)
			return err
		}
	}
	s.Connection = conn
	s.Channel = ch
	return nil
}

func (s *session) publish(ctx context.Context, exchange string, messages <-chan []byte) {
	defer common.TryE()
	var (
		running bool
		reading = messages
		pending = make(chan []byte, 1)
		confirm = make(chan amqp.Confirmation, 1)
	)

	if err := s.Confirm(false); err != nil {
		common.Log.Errorln("publisher confirms not supported")
		close(confirm) // confirms not supported, simulate by always nacking
	} else {
		s.NotifyPublish(confirm)
	}
	defer close(confirm)
	defer close(pending)
	//common.Log.Traceln("publishing...")
Publish:
	for {
		var body []byte
		select {
		case confirmed, ok := <-confirm:
			if !ok {
				common.Log.Traceln("confirm chan is closed")
				break Publish
			}
			if !confirmed.Ack {
				common.Log.Errorf("nack message %d, body: %q", confirmed.DeliveryTag, string(body))
			}
			reading = messages

		case body = <-pending:
			//routingKey := "ignored for fanout exchanges, application dependent for other exchanges"
			err := s.Publish(exchange, "", false, false, amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})
			// Retry failed delivery on the next session
			if err != nil {
				pending <- body
				s.Close()
				common.Log.Errorln("Publish failed!!")
				break Publish
			}

		case body, running = <-reading:
			// all messages consumed
			if !running {
				return
			}
			// work on pending delivery until ack'd
			pending <- body
			reading = nil
		case <-ctx.Done():
			common.Log.Infoln("shutting down  publishing")
			return
		}
	}

}

func (s *session) subscribe(queueName, exchange string, callBack func([]byte)) {
	defer common.TryE()
	if _, err := s.QueueDeclare(queueName, true, false, false, false, nil); err != nil {
		str := fmt.Sprintf("cannot consume from exclusive queue: %q, %v", queueName, err)
		common.Log.Traceln(str)
		panic(str)
		return
	}

	if exchange!=""{
		//routingKey := "application specific routing key for fancy toplogies"
		if err := s.QueueBind(queueName, "", exchange, false, nil); err != nil {
			str := fmt.Sprintf("cannot consume without a binding to exchange: %q, %v", exchange, err)
			common.Log.Traceln(str)
			panic(str)
			return
		}
	}

	deliveries, err := s.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		str := fmt.Sprintf("cannot consume from: %q, %v", queueName, err)
		common.Log.Traceln(str)
		panic(str)
		return
	}

	//common.Log.Tracef(" subscribed【%v】...",queueName)
	for msg := range deliveries {
		callBack(msg.Body)
		s.Ack(msg.DeliveryTag, false)
		//common.Log.Tracef("msg.DeliveryTag=%v", msg.DeliveryTag)
		time.Sleep(time.Nanosecond)
	}
	return
}
