package multicache_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	cache  map[string]string
	router *gin.Engine
)

func init() {
	// Initialize cache
	initializeCache()
	// Initialize router instance
	router = setupRouter()
}

func initializeCache() {
	cache = make(map[string]string)
}

// Handlers for cache operations
func handleSet(c *gin.Context) {
	key := c.Param("key")
	value := c.Query("value")
	cache[key] = value
	c.String(http.StatusOK, "Set operation successful")
}

func handleGet(c *gin.Context) {
	key := c.Param("key")
	value, found := cache[key]
	if !found {
		c.String(http.StatusNotFound, "Key not found")
		return
	}
	c.String(http.StatusOK, value)
}

func handleDelete(c *gin.Context) {
	key := c.Param("key")

	// Check if key exists in cache
	_, found := cache[key]
	if !found {
		c.String(http.StatusNotFound, "Key not found")
		return
	}

	// Delete operation
	delete(cache, key)
	c.String(http.StatusOK, "Delete operation successful")
}

func handleGetAll(c *gin.Context) {
	keys := make([]string, 0, len(cache))
	for k := range cache {
		keys = append(keys, k)
	}
	c.JSON(http.StatusOK, keys)
}

func handleDeleteAll(c *gin.Context) {
	initializeCache()
	c.String(http.StatusOK, "All keys deleted")
}

func TestCacheOperations(t *testing.T) {
	// Positive test cases
	t.Run("TestSetAndGet", func(t *testing.T) {
		key := "testkey"
		value := "testvalue"

		// Test Set operation
		w := performRequest("POST", "/cache/"+key+"?value="+value, router)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
		}

		// Test Get operation
		w = performRequest("GET", "/cache/"+key, router)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
		}
		if body := w.Body.String(); body != value {
			t.Errorf("Expected response body %s but got %s", value, body)
		}
	})

	// Negative test cases
	t.Run("TestGetNonExistingKey", func(t *testing.T) {
		key := "nonexistent"

		// Test Get operation on a non-existing key
		w := performRequest("GET", "/cache/"+key, router)
		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status code %d but got %d", http.StatusNotFound, w.Code)
		}
	})

	t.Run("TestDeleteNonExistingKey", func(t *testing.T) {
		key := "nonexistent"

		// Test Delete operation on a non-existing key
		w := performRequest("DELETE", "/cache/"+key, router)
		if w.Code != http.StatusNotFound { // Update expected status code here
			t.Errorf("Expected status code %d but got %d", http.StatusNotFound, w.Code)
		}

		// Verify that the response message indicates key not found
		expectedMsg := "Key not found"
		if body := w.Body.String(); body != expectedMsg {
			t.Errorf("Expected response body '%s' but got '%s'", expectedMsg, body)
		}
	})

}

// Helper function to perform HTTP requests
func performRequest(method, path string, r http.Handler) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// HTTP routes setup
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/cache/:key", handleSet)
	r.GET("/cache/:key", handleGet)
	r.DELETE("/cache/:key", handleDelete)
	r.GET("/cache", handleGetAll)
	r.DELETE("/cache", handleDeleteAll)

	return r
}
