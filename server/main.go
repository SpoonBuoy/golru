package main

import (
	"fmt"
	"lru/lru"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Configure CORS middleware options
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	}))

	var (
		PORT = ":9000"
		SIZE = 10
	)
	var lru = lru.NewLRU(uint(SIZE))
	//cleans up expired
	go lru.CleanUpExpired()

	r.GET("/:key", func(ctx *gin.Context) {
		keyStr := ctx.Param("key")
		key, _ := strconv.Atoi(keyStr)
		fmt.Println("KEY : ", key)
		val := lru.Get(key)
		lru.Print()
		if val == -1 {
			ctx.JSON(http.StatusOK, gin.H{"error": "Key Does Not Exist"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"value": val})

	})
	// r.GET("/", func(ctx *gin.Context) {
	// 	lru = LRU.NewLRU(uint(SIZE))
	// 	ctx.JSON(http.StatusOK, gin.H{"message": "New LRU Created"})
	// })
	r.POST("/", func(ctx *gin.Context) {
		type Body struct {
			Key    int `json:"key"`
			Value  int `json:"value"`
			Expiry int `json:"expiry"`
		}
		var body Body
		err := ctx.Bind(&body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unexpected Body " + err.Error()})
			return
		}
		exp := time.Second * time.Duration(body.Expiry)
		lru.Set(body.Key, body.Value, exp)
		lru.Print()
		ctx.JSON(http.StatusOK, gin.H{"value": body.Value})

	})
	r.GET("/top", func(ctx *gin.Context) {
		res := lru.Top10()
		ctx.JSON(http.StatusOK, gin.H{"entries": res})
	})

	r.Run(PORT)
}
