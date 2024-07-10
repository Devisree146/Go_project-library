package redis_cache_test

import (
	"testing"

	"github.com/Devisree146/Go_project-library.git/redis_cache"
)

func TestRedisCache_SetGetDelete(t *testing.T) {
	cache := redis_cache.NewRedisCache("localhost:6379", "", 0, 3)

	// Positive test case: Set and Get
	key := "test_key"
	value := 42
	err := cache.Set(key, value, redis_cache.StandardTTL)
	if err != nil {
		t.Errorf("Set() error = %v, want nil", err)
	}

	retValue, err := cache.Get(key)
	if err != nil {
		t.Errorf("Get() error = %v, want nil", err)
	}
	if retValue != value {
		t.Errorf("Get() got = %v, want %v", retValue, value)
	}

	// Positive test case: Delete key
	err = cache.Delete(key)
	if err != nil {
		t.Errorf("Delete() error = %v, want nil", err)
	}

	// Verify key is deleted
	retValue, err = cache.Get(key)
	if err != redis_cache.ErrCacheMiss {
		t.Errorf("Get() error = %v, want ErrCacheMiss", err)
	}
	if retValue != 0 { // Change comparison from nil to 0
		t.Errorf("Get() got = %v, want %v", retValue, 0)
	}

	// Negative test case: Delete non-existent key
	err = cache.Delete("non_existing_key")
	if err != nil {
		t.Errorf("Delete() error = %v, want nil", err)
	}
}

func TestRedisCache_GetAllKeys(t *testing.T) {
	cache := redis_cache.NewRedisCache("localhost:6379", "", 0, 3)

	// Positive test case: Get all keys
	keys, err := cache.GetAllKeys()
	if err != nil {
		t.Errorf("GetAllKeys() error = %v, want nil", err)
	}
	if len(keys) != 0 {
		t.Errorf("GetAllKeys() got = %v keys, want 0", len(keys))
	}

	// Negative test case: Get all keys when Redis is not reachable
	// Simulate a scenario where Redis is not accessible
	cache = redis_cache.NewRedisCache("non_existing_host:6379", "", 0, 3)
	keys, err = cache.GetAllKeys()
	if err == nil {
		t.Errorf("GetAllKeys() expected error, got nil")
	}
}

func TestRedisCache_DeleteAll(t *testing.T) {
	cache := redis_cache.NewRedisCache("localhost:6379", "", 0, 3)

	// Set some keys first for testing
	cache.Set("key1", 1, redis_cache.StandardTTL)
	cache.Set("key2", 2, redis_cache.StandardTTL)

	// Delete all keys
	err := cache.DeleteAll()
	if err != nil {
		t.Errorf("DeleteAll() error = %v, want nil", err)
	}

	// Verify all keys are deleted
	keys, err := cache.GetAllKeys()
	if err != nil {
		t.Errorf("GetAllKeys() error after DeleteAll() = %v, want nil", err)
	}
	if len(keys) != 0 {
		t.Errorf("GetAllKeys() got = %v keys after DeleteAll(), want 0", len(keys))
	}

	// Negative test case: Delete all keys when Redis is not reachable
	// Simulate a scenario where Redis is not accessible
	cache = redis_cache.NewRedisCache("non_existing_host:6379", "", 0, 3)
	err = cache.DeleteAll()
	if err == nil {
		t.Errorf("DeleteAll() expected error, got nil")
	}
}
