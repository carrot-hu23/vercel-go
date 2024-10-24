package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"vercel-go/middleware"
)

var (
	app  *gin.Engine
	list []string
)

type Response struct {
	Code int         `json:"code"` //提示代码
	Msg  string      `json:"msg"`  //提示信息
	Data interface{} `json:"data"` //数据
}

// init gin app
func init() {

	app = gin.New()
	app.Use(middleware.Recover)

	// Handling routing errors
	app.NoRoute(func(c *gin.Context) {
		sb := &strings.Builder{}
		sb.WriteString("routing err: no route, try this:\n")
		for _, v := range app.Routes() {
			sb.WriteString(fmt.Sprintf("%s %s\n", v.Method, v.Path))
		}
		c.String(http.StatusBadRequest, sb.String())
	})

	r := app.Group("/")

	// register route
	registerRouter(r)
}

func registerRouter(r *gin.RouterGroup) {
	r.GET("/api/add/:path", func(ctx *gin.Context) {
		path := ctx.Param("path")
		list = append(list)
		ctx.String(http.StatusOK, path)
	})
	r.GET("/api/list", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, Response{
			Code: 200,
			Msg:  "success",
			Data: list,
		})
	})
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello vercel go")
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
