package main

import (
	"github.com/gin-gonic/gin"
	"github.com/herrluk/test/common"
)

func main() {
	db := common.InitDB()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run(":8080"))
}
