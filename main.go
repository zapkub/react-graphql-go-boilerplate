package main

import (
	"log"
	"net/http/httputil"
	"net/url"
	"react-graphql-go-boilerplate/pkg/server"
	"time"

	"github.com/gin-gonic/gin"
)

func reverseProxy(target string) gin.HandlerFunc {

	return func(c *gin.Context) {
		url, _ := url.Parse(target)
		handler := httputil.NewSingleHostReverseProxy(url)
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

// entry point
func main() {

	r := server.NewAPIRoutes()

	// Create reverse proxy fallback
	// to serve client application
	// or you can setup from NGINX
	target := "http://localhost:3001"
	url, _ := url.Parse(target)
	handler := httputil.NewSingleHostReverseProxy(url)

	// line below is only for
	// development with nextjs hot reload
	// webpack hmr use Http event stream to
	// update hot loader status, if not provide this
	// flush interval, event stream will not flushing
	// message until request is close
	handler.FlushInterval = 100 * time.Millisecond
	r.NoRoute(func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	})

	err := r.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
