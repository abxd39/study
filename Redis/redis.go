package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	//"github.com/garyburd/redigo/redis"
	redis1 "github.com/gomodule/redigo/redis"
)

const checkMark = "\u2713" // 对号 %v
const ballotX = "\u2717"   // 错号

var RedisPool *redis1.Pool

const CouponBuyPreventAttack = "prevent_attack:%v"

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	//添加防刷功能
	err := OpenRedis()
	defer RedisPool.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	group := sync.WaitGroup{}
	key := fmt.Sprintf(CouponBuyPreventAttack, 1)
	CouponBuyTotlExist(key, 10)
	for i := 0; i <= 12; i++ {
		//fmt.Println(i)
		group.Add(1)
		//go1(group)
		go go2(key, &group)

	}
	group.Wait()
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<-quit

	fmt.Println("Shutdown Server ...")
}

//把购券总数写缓存
//key的过期时间为一小时
func CouponBuyTotlExist(key string, total int) {
	redisConn := RedisPool.Get()
	defer redisConn.Close()
	redis1.Int(redisConn.Do("SETNX", key, total))
	redisConn.Do("EXPIRE", key, 60)
}

func go2(key string, group *sync.WaitGroup) {
	CouponBuyTotlExist(key, 10)
	count := CouponBuyTotalUpdate(key)
	if count < 0 {
		fmt.Println(ballotX, "-----------", count)
		group.Done()
		return
	}
	if count == 0 {
		CouponBuyDeleteKey(key)
	}
	group.Done()
	fmt.Println(checkMark, "    ", count)
}

func go1(group sync.WaitGroup) {
	go func() {
		key := fmt.Sprintf(CouponBuyPreventAttack, 1011)
		bl := CouponBuyLock(key)
	Lock:
		if bl {
			fmt.Println("加锁失败------------------------")
			for true {
				count := rand.Intn(10)
				fmt.Println(count)
				time.Sleep(time.Duration(count) * time.Nanosecond)
				bl = CouponBuyLock(key)
				if !bl {
					goto Lock
				} else {
					fmt.Println("加锁失败")
				}
			}

		} else {
			fmt.Println("加锁成功！！！")
			CouponBuyDeleteLock(key)
			group.Done()
		}
	}()
}

//
func CouponBuyLock(key string) bool {
	redisConn := RedisPool.Get()
	defer redisConn.Close()
	res, _ := redis1.Int(redisConn.Do("SETNX", key, "lock"))
	redisConn.Do("EXPIRE", key, 1)
	if res == 1 {
		return false
	} else {
		return true
	}
}

//解锁
func CouponBuyDeleteLock(key string) {
	redisConn := RedisPool.Get()
	defer redisConn.Close()
	redisConn.Do("DEL", key)
}

func OpenRedis() error {
	var err error
	RedisPool = &redis1.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis1.Conn, error) {
			conn, err := redis1.Dial("tcp", "10.20.20.82:6379", redis1.DialDatabase(3), redis1.DialPassword(""))
			//if Config.Redis.ShowCommand {
			//	conn = redis1.NewLoggingConn(conn, Log.info, "[redis]")
			//}
			return conn, err
		},
	}
	return err
}

func CouponBuyTotalUpdate(key string) int {
	redisConn := RedisPool.Get()
	defer redisConn.Close()
	res, err := redis1.Int(redisConn.Do("DECRBY", key, 1))
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return res
}

func CouponBuyDeleteKey(key string) {
	redisConn := RedisPool.Get()
	defer redisConn.Close()
	r, err := redisConn.Do("DEL", key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", r)
}

func SetNxValueWithDuration(key string, val interface{}, duration time.Duration) error {
	redisConn := RedisPool.Get()
	defer redisConn.Close()
	res, err := redis1.Int(redisConn.Do("SETNX", key, val))
	redisConn.Do("EXPIRE", key, duration)
	if res == 1 {
		return err
	} else {
		return err
	}
}
