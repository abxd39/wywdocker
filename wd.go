package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init(){
	log.SetFlags(log.Lshortfile|log.Ldate|log.Ltime)
}

func main() {
	log.Println("Hello from Seelog!")
	log.Println("Hello Docker 世界！！")
	_=Auto("Hello World !")
	r:=gin.Default()
	r.GET("wyw/docker", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"name":"docker",
			"target":"study",
		})
	})
	err:=r.Run()
	if err!=nil{
		log.Fatal(err)
	}
}

func Auto(param string)error{
	if param=="Hello World !"{
		log.Printf("param=%s",param)
		return nil
	}else {
		err:=errors.New("字符串输入错误！！")
		log.Println(err)
		return err
	}
}