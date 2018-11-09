package main

import (
	"fmt"
	"github.com/ghost0036/go-hello/web/gin-blog/pkg/setting"
	"github.com/ghost0036/go-hello/web/gin-blog/routers"
	"net/http"
)

func main() {
	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:           routers.InitRouter(),
		ReadHeaderTimeout: setting.ReadTimeOut,
		WriteTimeout:      setting.WriteTimeOut,
		MaxHeaderBytes:    1 << 20,
	}
	s.ListenAndServe()
}
