package redis_cache_test

import (
	"strconv"
	"testing"

	"github.com/Devisree146/Go_project-library.git/redis_cache"
)

const (
	benchmarkKey   = "benchmark_key"
	benchmarkValue = 12345
)

func BenchmarkRedisCache_Set(b *testing.B) {
	cache := redis_cache.NewRedisCache("localhost:6379", "", 0, 1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Set(benchmarkKey, benchmarkValue, redis_cache.StandardTTL)
	}
}

func BenchmarkRedisCache_Get(b *testing.B) {
	cache := redis_cache.NewRedisCache("localhost:6379", "", 0, 1000)
	cache.Set(benchmarkKey, benchmarkValue, redis_cache.StandardTTL)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Get(benchmarkKey)
	}
}

func BenchmarkRedisCache_Delete(b *testing.B) {
	cache := redis_cache.NewRedisCache("localhost:6379", "", 0, 1000)
	cache.Set(benchmarkKey, benchmarkValue, redis_cache.StandardTTL)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Delete(benchmarkKey)
	}
}

func BenchmarkRedisCache_DeleteAll(b *testing.B) {
	cache := redis_cache.NewRedisCache("localhost:6379", "", 0, 1000)
	for i := 0; i < 100; i++ {
		cache.Set(benchmarkKey+strconv.Itoa(i), benchmarkValue, redis_cache.StandardTTL)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.DeleteAll()
	}
}

func BenchmarkRedisCache_GetAllKeys(b *testing.B) {
	cache := redis_cache.NewRedisCache("localhost:6379", "", 0, 1000)
	for i := 0; i < 100; i++ {
		cache.Set(benchmarkKey+strconv.Itoa(i), benchmarkValue, redis_cache.StandardTTL)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.GetAllKeys()
	}
}

// Negative Benchmarking

func BenchmarkRedisCache_Set_Failure(b *testing.B) {
	cache := redis_cache.NewRedisCache("non_existing_host:6379", "", 0, 1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Set(benchmarkKey, benchmarkValue, redis_cache.StandardTTL)
	}
}

func BenchmarkRedisCache_Get_Failure(b *testing.B) {
	cache := redis_cache.NewRedisCache("non_existing_host:6379", "", 0, 1000)
	cache.Set(benchmarkKey, benchmarkValue, redis_cache.StandardTTL)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Get(benchmarkKey)
	}
}

func BenchmarkRedisCache_Delete_Failure(b *testing.B) {
	cache := redis_cache.NewRedisCache("non_existing_host:6379", "", 0, 1000)
	cache.Set(benchmarkKey, benchmarkValue, redis_cache.StandardTTL)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Delete(benchmarkKey)
	}
}

func BenchmarkRedisCache_DeleteAll_Failure(b *testing.B) {
	cache := redis_cache.NewRedisCache("non_existing_host:6379", "", 0, 1000)
	for i := 0; i < 100; i++ {
		cache.Set(benchmarkKey+strconv.Itoa(i), benchmarkValue, redis_cache.StandardTTL)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.DeleteAll()
	}
}

func BenchmarkRedisCache_GetAllKeys_Failure(b *testing.B) {
	cache := redis_cache.NewRedisCache("non_existing_host:6379", "", 0, 1000)
	for i := 0; i < 100; i++ {
		cache.Set(benchmarkKey+strconv.Itoa(i), benchmarkValue, redis_cache.StandardTTL)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.GetAllKeys()
	}
}
