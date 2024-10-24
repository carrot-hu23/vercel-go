package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	app *gin.Engine
)

// init gin app
func init() {
	app = gin.New()

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
	r.GET("/api/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello")
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
