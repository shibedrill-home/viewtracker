package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func get_view(c *gin.Context) {
	articleId := c.Param("id")

	// Fetch the value (will be 0 if not exist)
	views, err := rdb.Get(ctx, articleId).Int()
	if err != nil && err != redis.Nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Increment the views, since someone just viewed this article
	views++
	err = rdb.Set(ctx, articleId, views, 0).Err()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"viewcounter": views,
	})
}
