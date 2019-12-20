package bucket

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RateBucket struct {

}

func (r*RateBucket)Test(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK,gin.H{
		"msg":"hello 世界！！",
		"code":200,
	})
}
