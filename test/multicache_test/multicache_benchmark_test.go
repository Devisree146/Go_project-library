package multicache_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// BenchmarkCacheOperations benchmarks various cache operations.
func BenchmarkCacheOperations(b *testing.B) {
	// Initialize router and cache for benchmarks
	router := setupRouter()
	initializeCache()

	// Benchmark SET operation
	b.Run("BenchmarkSetOperation", func(b *testing.B) {
		b.ResetTimer() // Reset timer before each benchmark
		for i := 0; i < b.N; i++ {
			key := "testkey"
			value := "testvalue"

			req, _ := http.NewRequest("POST", "/cache/"+key+"?value="+value, nil)
			rr := httptest.NewRecorder()

			b.StartTimer()
			router.ServeHTTP(rr, req)
			b.StopTimer()
		}
		b.ReportMetric(float64(b.N), "operations") // Report number of operations
	})

	// Benchmark GET operation
	b.Run("BenchmarkGetOperation", func(b *testing.B) {
		b.ResetTimer() // Reset timer before each benchmark
		for i := 0; i < b.N; i++ {
			key := "testkey"

			req, _ := http.NewRequest("GET", "/cache/"+key, nil)
			rr := httptest.NewRecorder()

			b.StartTimer()
			router.ServeHTTP(rr, req)
			b.StopTimer()
		}
		b.ReportMetric(float64(b.N), "operations") // Report number of operations
	})

	// Benchmark DELETE operation
	b.Run("BenchmarkDeleteOperation", func(b *testing.B) {
		b.ResetTimer() // Reset timer before each benchmark
		for i := 0; i < b.N; i++ {
			key := "testkey"

			req, _ := http.NewRequest("DELETE", "/cache/"+key, nil)
			rr := httptest.NewRecorder()

			b.StartTimer()
			router.ServeHTTP(rr, req)
			b.StopTimer()
		}
		b.ReportMetric(float64(b.N), "operations") // Report number of operations
	})

	// Benchmark GET ALL operation
	b.Run("BenchmarkGetAllOperation", func(b *testing.B) {
		b.ResetTimer() // Reset timer before each benchmark
		for i := 0; i < b.N; i++ {
			req, _ := http.NewRequest("GET", "/cache", nil)
			rr := httptest.NewRecorder()

			b.StartTimer()
			router.ServeHTTP(rr, req)
			b.StopTimer()
		}
		b.ReportMetric(float64(b.N), "operations") // Report number of operations
	})

	// Benchmark DELETE ALL operation
	b.Run("BenchmarkDeleteAllOperation", func(b *testing.B) {
		b.ResetTimer() // Reset timer before each benchmark
		for i := 0; i < b.N; i++ {
			req, _ := http.NewRequest("DELETE", "/cache", nil)
			rr := httptest.NewRecorder()

			b.StartTimer()
			router.ServeHTTP(rr, req)
			b.StopTimer()
		}
		b.ReportMetric(float64(b.N), "operations") // Report number of operations
	})
}

// Ensure that the BenchmarkCacheOperations function is recognized as a benchmark
var _ = BenchmarkCacheOperations
