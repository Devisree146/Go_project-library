package in_memory_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Devisree146/Go_project-library.git/in_memory"
)

var cache *in_memory.InMemoryCache

func BenchmarkSet(b *testing.B) {
	cache = in_memory.NewInMemoryCache(3, 5*time.Minute)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Set(fmt.Sprintf("key%d", i), "value")
	}
}

func BenchmarkGet(b *testing.B) {
	cache.Set("key", "value") // Set up initial state for Get benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get("key")
	}
}

func BenchmarkDelete(b *testing.B) {
	cache.Set("key", "value") // Set up initial state for Delete benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Delete("key")
	}
}

func BenchmarkNegativeGet(b *testing.B) {
	cache = in_memory.NewInMemoryCache(3, 5*time.Minute)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Attempting to get a non-existent key
		cache.Get("nonexistent_key")
	}
}

func BenchmarkNegativeDelete(b *testing.B) {
	cache = in_memory.NewInMemoryCache(3, 5*time.Minute)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Attempting to delete a non-existent key
		cache.Delete("nonexistent_key")
	}
}

func BenchmarkTTLExpiration(b *testing.B) {
	cache = in_memory.NewInMemoryCache(3, 2*time.Second) // TTL set to 2 seconds
	cache.Set("key", "value")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get("key")
	}
}

func BenchmarkLRUEviction(b *testing.B) {
	cache = in_memory.NewInMemoryCache(3, 5*time.Minute)
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")
	cache.Get("key1")           // Access key1 to make it least recently used
	cache.Set("key4", "value4") // This should evict key2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get("key2")
	}
}

func BenchmarkExists(b *testing.B) {
	cache.Set("key", "value")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Exists("key")
	}
}

func BenchmarkGetAllKeys(b *testing.B) {
	keys := []string{"key1", "key2", "key3"}
	for _, key := range keys {
		cache.Set(key, "value")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.GetAllKeys()
	}
}
