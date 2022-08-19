package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gmcli/client"
	"gmcli/service"
	"net/http"
)

func main() {
	//ServTest1()
	ServTest2()
}

func httpClinet(){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		s:=service.Ping(context.TODO())
		c.JSON(http.StatusOK, gin.H{
			"message": s,
		})
		return
	})
	r.Run(":8090") //使用http://127.0.0.1:8090/ping 可以获取数据
}

func ServTest1(){//ok
	mic,_:=client.NewService("dev","myclient",9900)
	client.Register(mic)
	go httpClinet()
	if err := mic.Run(); err != nil {
		fmt.Println(err)
	}
}

func ServTest2(){
	mic,_:=client.NewService("dev","myclient",9900)
	client.Register(mic)
	httpClinet()
}