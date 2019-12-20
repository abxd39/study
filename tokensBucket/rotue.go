package main

import (
	"github.com/abxd39/myproject/tokensBucket/bucket"
	"github.com/gin-gonic/gin"
)

func httpRouter(r *gin.Engine) {

	v1 := r.Group("/v1",Middleware(B))
	{
		v1.POST("/test",new(bucket.RateBucket).Test)
	}
}
