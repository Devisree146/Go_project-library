package in_memory_test

import (
	"testing"
	"time"

	"github.com/Devisree146/Go_project-library.git/in_memory"
)

func TestNewInMemoryCache(t *testing.T) {
	cache := in_memory.NewInMemoryCache(3, 5*time.Minute)
	if cache == nil {
		t.Fatal("expected cache to be created")
	}
}

func TestSetGet(t *testing.T) {
	cache := in_memory.NewInMemoryCache(3, 5*time.Minute)

	// Positive Test Case: Set and Get
	err := cache.Set("key1", 100)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	value, err := cache.Get("key1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if value != 100 {
		t.Errorf("expected value 100, got %v", value)
	}

	// Negative Test Case: Get non-existing key
	_, err = cache.Get("nonexistent")
	if err == nil {
		t.Error("expected error for non-existing key, got nil")
	}

	// Negative Test Case: Set with empty key
	err = cache.Set("", 200)
	if err == nil {
		t.Error("expected error for empty key, got nil")
	}

	// Negative Test Case: Set with nil value
	err = cache.Set("key2", nil)
	if err == nil {
		t.Error("expected error for nil value, got nil")
	}
}

func TestDelete(t *testing.T) {
	cache := in_memory.NewInMemoryCache(3, 5*time.Minute)

	cache.Set("key1", 100)
	err := cache.Delete("key1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Negative Test Case: Delete non-existing key
	err = cache.Delete("nonexistent")
	if err == nil {
		t.Error("expected error for non-existing key, got nil")
	}
}

func TestDeleteAll(t *testing.T) {
	cache := in_memory.NewInMemoryCache(3, 5*time.Minute)

	cache.Set("key1", 100)
	cache.Set("key2", 200)
	cache.DeleteAll()

	if len(cache.GetAllKeys()) != 0 {
		t.Error("expected all keys to be deleted")
	}
}

func TestTTLExpiration(t *testing.T) {
	cache := in_memory.NewInMemoryCache(3, 1*time.Second)

	cache.Set("key1", 100)
	time.Sleep(2 * time.Second)

	_, err := cache.Get("key1")
	if err == nil {
		t.Error("expected error for expired key, got nil")
	}
}

func TestLRUEviction(t *testing.T) {
	cache := in_memory.NewInMemoryCache(2, 5*time.Minute)

	cache.Set("key1", 100)
	cache.Set("key2", 200)
	cache.Set("key3", 300)

	_, err := cache.Get("key1")
	if err == nil {
		t.Error("expected error for evicted key, got nil")
	}
}

func TestExists(t *testing.T) {
	cache := in_memory.NewInMemoryCache(3, 5*time.Minute)

	cache.Set("key1", 100)
	if !cache.Exists("key1") {
		t.Error("expected key1 to exist")
	}

	if cache.Exists("nonexistent") {
		t.Error("expected nonexistent key to not exist")
	}
}

func TestGetAllKeys(t *testing.T) {
	cache := in_memory.NewInMemoryCache(3, 5*time.Minute)

	cache.Set("key1", 100)
	cache.Set("key2", 200)

	keys := cache.GetAllKeys()
	expectedKeys := []string{"key1", "key2"}

	if len(keys) != len(expectedKeys) {
		t.Errorf("expected %v, got %v", expectedKeys, keys)
	}
}

func TestSetAndGet(t *testing.T) {
	cache := in_memory.NewInMemoryCache(2, 5*time.Minute)

	// Positive Test Case: Set and Get a key-value pair
	err := cache.Set("key1", "value1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	value, err := cache.Get("key1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if value != "value1" {
		t.Fatalf("Expected value1, got %v", value)
	}

	// Negative Test Case: Get a non-existent key
	_, err = cache.Get("key2")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestEvictionPolicy(t *testing.T) {
	cache := in_memory.NewInMemoryCache(2, 5*time.Minute)

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3") // This should evict "key1"

	_, err := cache.Get("key1")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

	value, err := cache.Get("key2")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if value != "value2" {
		t.Fatalf("Expected value2, got %v", value)
	}

	value, err = cache.Get("key3")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if value != "value3" {
		t.Fatalf("Expected value3, got %v", value)
	}
}
