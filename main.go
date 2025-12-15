package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("ADDR"),
	Password: os.Getenv("PASS"),
	DB:       0,
	Protocol: 2,
})

func main() {
	r := gin.Default()
	r.GET("/api/view/:id", get_view)
	r.Run()
}
