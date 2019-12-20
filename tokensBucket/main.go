package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	B = ratelimit.NewBucket(time.Second*1, 5)
}

var B *ratelimit.Bucket

func Middleware(b *ratelimit.Bucket) gin.HandlerFunc {
	return func(context *gin.Context) {
		v := b.TakeAvailable(1)
		if v == 0 {
			context.JSON(http.StatusBadRequest, gin.H{
				"msg": "限流",
				"code":400,
			})
			context.Abort()
			return

		}
		//b.Wait(1)
		context.Next()
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	//r.Use(middleware.Logger(), gin.Recovery())
	httpRouter(r)
	srv := &http.Server{
		Addr:    "localhost:6061",
		Handler: r,
	}
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
	// service connections
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("listen: %s\n", err)
	}
	defer srv.Shutdown(context.TODO())

}

func Begin() error {

	rq,err:=http.NewRequest("POST","http://localhost:6061",nil)
	if err!=nil{
		log.Println(err)
		return err
	}
	rq.Header.Set("Content-type","application/json")
	res,err:=http.DefaultClient.Do(rq)
	if err!=nil{
		log.Println(err)
		return err
	}
	var result = struct {
		Code int64
		Msg string
	}{}
	bts,err:=ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err!=nil{
		log.Println(err)
		return err
	}
	err=json.Unmarshal(bts,&result)
	if err!=nil{
		log.Println(err)
		return err
	}
	if result.Code!=200{
		return fmt.Errorf("bad request !! %v",runtime.NumGoroutine())
	}else {
		return nil
	}


}
