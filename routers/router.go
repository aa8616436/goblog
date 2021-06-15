package routers

import (
	"fmt"
	"github.com/aa8616436/goblog/pkg/setting"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type GinConfig struct {
	setting.Gin
}

func InitHttp() {
	g := GinConfig{
		Gin: setting.Setting.Gin,
	}
	//启动gin路由
	var router *gin.Engine
	router = gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	r := router.Group(fmt.Sprintf("%s/api", "gobolg"))

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"test": "This is a test"})
	})

	server := http.Server{
		Addr:           fmt.Sprintf(":%d", g.HttpPort),
		Handler:        router,
		ReadTimeout:    time.Duration(g.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(g.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("http port : ", g.HttpPort)
	if e := server.ListenAndServe(); e != nil {
		if e == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}
	log.Printf("connect gin success")
}
