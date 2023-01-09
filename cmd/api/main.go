package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/molpadia/molpastream/configs"
	"github.com/molpadia/molpastream/internal/app"
)

func main() {
	ctx := context.Background()
	configs.Init()
	r := gin.Default()
	app.SetupRoutes(r, ctx)
	r.Run()
}
