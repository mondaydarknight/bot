package main

import (
	"github.com/gin-gonic/gin"
	"github.com/molpadia/molpastream/configs"
)

func main() {
	configs.Init()
	r := gin.Default()
	r.Run()
}
