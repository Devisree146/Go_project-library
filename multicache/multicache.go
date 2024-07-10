package multicache

import (
	"net/http"
	"time"
	"github.com/Devisree146/Go_project-library.git/in_memory"
	"github.com/Devisree146/Go_project-library.git/redis_cache"

	"github.com/gin-gonic/gin"
)

var cacheInMemory *in_memory.InMemoryCache
var cacheRedis *redis_cache.Cache

func init() {
	cacheInMemory = in_memory.NewInMemoryCache(3, 5*time.Minute) // Initialize with TTL
	cacheRedis = redis_cache.NewRedisCache("localhost:6379", "", 0, 3)
}

// Example usage in your API handlers or other functions
func ExampleHandler(c *gin.Context) {
	value, err := cacheInMemory.Get("some_key")
	if err != nil {
		if err == in_memory.ErrCacheMiss {
			// Handle cache miss for in-memory cache
			// Example response
			c.JSON(http.StatusNotFound, gin.H{"error": "Key not found in in-memory cache"})
			return
		}
		// Handle other errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving key from in-memory cache"})
		return
	}

	// Continue processing with the retrieved value
	c.JSON(http.StatusOK, gin.H{"key": "some_key", "value": value})
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Example middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/cache", func(c *gin.Context) {
		var data struct {
			Key   string `json:"key"`
			Value int    `json:"value"`
			TTL   string `json:"ttl"` // Assuming you may receive it, but not mandatory
		}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Handle TTL logic
		var ttlDuration time.Duration
		if data.TTL != "" {
			var err error
			ttlDuration, err = time.ParseDuration(data.TTL)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid TTL format"})
				return
			}
		} else {
			// Use default TTL here if TTL is not provided in the request
			ttlDuration = 5 * time.Minute
		}

		// Convert value to interface{} type before setting in cache
		value := interface{}(data.Value)

		err := cacheInMemory.Set(data.Key, value)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set key in in-memory cache"})
			return
		}

		// Handling TTL logic separately if needed
		time.AfterFunc(ttlDuration, func() {
			cacheInMemory.Delete(data.Key)
		})

		// Perform type assertion to int before setting in Redis cache
		intValue, ok := value.(int)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assert value to int"})
			return
		}

		err = cacheRedis.Set(data.Key, intValue, ttlDuration)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set key in Redis cache"})
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

		value, err := cacheInMemory.Get(key)
		if err != nil {
			if err == in_memory.ErrCacheMiss {
				value, err = cacheRedis.Get(key)
				if err != nil {
					if err == redis_cache.ErrCacheMiss {
						c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
					} else {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving key from Redis cache"})
					}
					return
				}
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving key from in-memory cache"})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
	})

	router.DELETE("/cache", func(c *gin.Context) {
		key := c.Query("key")
		if key == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Key not provided"})
			return
		}

		err := cacheInMemory.Delete(key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting key from in-memory cache"})
			return
		}

		err = cacheRedis.Delete(key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting key from Redis cache"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Key deleted successfully"})
	})

	router.GET("/cache/all", func(c *gin.Context) {
		cachedKeysInMemory := cacheInMemory.GetAllKeys() // Assuming GetAllKeys() retrieves all keys
		cachedKeysRedis, err := cacheRedis.GetAllKeys()  // Assuming GetAllKeys() retrieves all keys
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get keys from Redis cache"})
			return
		}

		allCachedKeys := append(cachedKeysInMemory, cachedKeysRedis...)

		c.JSON(http.StatusOK, gin.H{"keys": allCachedKeys})
	})

	router.DELETE("/cache/all", func(c *gin.Context) {
		cacheInMemory.DeleteAll()
		cacheRedis.DeleteAll()
		c.JSON(http.StatusOK, gin.H{"message": "All keys deleted successfully"})
	})

	return router
}
