package api_handler

import (
	"net/http"
	"time"

	"github.com/Devisree146/Go_project-library.git/redis_cache"
	"github.com/gin-gonic/gin"
)

const (
	standardTTL = 5 * time.Minute // Standard TTL of 5 minutes
)

func SetupRedisCacheRouter() *gin.Engine {
	// Initialize your Redis cache instance with maxSize of 3
	cache := redis_cache.NewRedisCache("localhost:6379", "", 0, 3)

	// Setup Gin router
	router := gin.Default()

	// POST endpoint to set cache with standard TTL
	router.POST("/cache", func(c *gin.Context) {
		var data CacheEntry
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		err := cache.Set(data.Key, data.Value, standardTTL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set key"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Key set successfully"})
	})

	// GET endpoint to retrieve cache
	router.GET("/cache", func(c *gin.Context) {
		key := c.Query("key")
		if key == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Key not provided"})
			return
		}

		value, err := cache.Get(key)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
	})

	// DELETE endpoint to delete cache by key
	router.DELETE("/cache", func(c *gin.Context) {
		key := c.Query("key")
		if key == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Key not provided"})
			return
		}

		err := cache.Delete(key)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Key deleted successfully"})
	})

	// DELETE endpoint to delete all cache keys
	router.DELETE("/cache/all", func(c *gin.Context) {
		cache.DeleteAll()

		c.JSON(http.StatusOK, gin.H{"message": "All keys deleted successfully"})
	})

	// GET endpoint to retrieve all cache keys
	router.GET("/cache/all", func(c *gin.Context) {
		cachedKeys, err := cache.GetAllKeys()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get keys"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"keys": cachedKeys})
	})

	return router
}
