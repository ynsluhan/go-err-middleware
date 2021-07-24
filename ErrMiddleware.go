package ErrMiddleware

import (
	"github.com/gin-gonic/gin"
	R "github.com/ynsluhan/go-r"
	"log"
	"runtime/debug"
)

/**
* @Author: yNsLuHan
* @Description:
* @File: ErrMiddleware
* @Version: 1.0.0
* @Date: 2021/6/8 5:14 下午
 */
func ErrMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// 打印错误堆栈信息
				log.Printf("panic: %v\n  ", r.(error).Error())
				// 打印详细错误
				debug.PrintStack()
				//封装通用json返回
				R.Ok(c, 0, r.(error).Error(), nil)
				//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
				c.Abort()
			}
		}()
		//加载完 defer recover，继续后续接口调用
		c.Next()
	}
}
