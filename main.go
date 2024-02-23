package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kylerequez/go-gin-dependency-injection-v1/api"
	"github.com/kylerequez/go-gin-dependency-injection-v1/common"
)

func main() {
	fmt.Println("Hello")
	run()
}

func run() {
	common.LoadEnvVariables()

	err := common.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer common.CloseDB()

	err = common.MigrateDB()
	if err != nil {
		panic(err)
	}
	defer common.CloseDB()

	// err = common.DropAllTables()
	// if err != nil {
	// 	panic(err)
	// }
	// defer common.CloseDB()

	r := gin.New()
	api.InitRoutes(r, common.GetDB())
	r.Run()
}
