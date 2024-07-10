package api_handler

import (
	"net/http"
	"time"

	"github.com/Devisree146/Go_project-library.git/in_memory"
	"github.com/gin-gonic/gin"
)

const TTL = 5 * time.Minute

func SetupInMemoryRouter() *gin.Engine {
	cache := in_memory.NewInMemoryCache(3, TTL)
	router := gin.Default()

	router.POST("/cache", func(c *gin.Context) {
		var data CacheEntry
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		if err := cache.Set(data.Key, data.Value); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Key set successfully"})
	})

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

	router.DELETE("/cache", func(c *gin.Context) {
		key := c.Query("key")
		if key == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Key not provided"})
			return
		}

		if err := cache.Delete(key); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Key deleted successfully"})
	})

	router.DELETE("/cache/all", func(c *gin.Context) {
		cache.DeleteAll()
		c.JSON(http.StatusOK, gin.H{"message": "All keys deleted successfully"})
	})

	router.GET("/cache/all", func(c *gin.Context) {
		cachedKeys := cache.GetAllKeys()
		c.JSON(http.StatusOK, gin.H{"keys": cachedKeys})
	})

	return router
}
