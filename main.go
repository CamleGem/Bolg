package main

import (
	"ginEssential/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)



func main()  {

	db := common.GetDb()
	defer db.Close()

	router := gin.Default()
	router = CollectRouter(router)

	// 指定地址和端口号
	panic(router.Run())
}






