package cache

import (
	"testing"
	"time"
)

func TestCacheSetAndGet(t *testing.T) {
	cache := New()

	key := "testKey"
	value := "testValue"
	ttl := time.Microsecond * 200

	cache.Set(key, value, ttl)

	// Check if the value is retrieved correctly
	result := cache.Get(key)
	if result != value {
		t.Errorf("Expected value %v, but got %v", value, result)
	}
}

func TestCacheSetWithTTL(t *testing.T) {
	cache := New()

	key := "testKey"
	value := "testValue"
	ttl := time.Second * 2

	cache.Set(key, value, ttl)

	// Wait for the TTL to expire
	time.Sleep(ttl + time.Second)

	// Check if the value is expired and not retrievable
	result := cache.Get(key)
	if result != nil {
		t.Errorf("Expected nil value, but got %v", result)
	}
}

func TestCacheDelete(t *testing.T) {
	cache := New()

	key := "testKey"
	value := "testValue"

	cache.Set(key, value, time.Second*2)

	// Check if the value is initially present
	result := cache.Get(key)
	if result != value {
		t.Errorf("Expected value %v, but got %v", value, result)
	}

	// Delete the value and check if it's removed
	cache.Delete(key)
	result = cache.Get(key)
	if result != nil {
		t.Errorf("Expected nil value after deletion, but got %v", result)
	}
}
