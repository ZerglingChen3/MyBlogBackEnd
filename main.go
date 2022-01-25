package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"myBlog/common"
	"os"
	"strconv"
)

func main() {
	InitConfig()

	common.InitDB()

	r := gin.Default()
	r = CollectRoute(r)

	port := strconv.Itoa(viper.Get("server.port").(int))

	fmt.Println(port)
	if port != "" {
		panic(r.Run(":" + port))
	} else {
		panic(r.Run()) // listen and serve on 0.0.0.0:8080
	}
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Read config Error!")
	}
}
