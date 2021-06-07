package main

import (
	_ "PrintCloud/model"
	"PrintCloud/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Http(r)
	_ = r.Run(":8090")
}
