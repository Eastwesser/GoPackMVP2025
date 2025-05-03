package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

// User represents a user with personal information
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
}

// Cacher defines the interfaces for cache operations
type Cacher interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	GetUser(key string) (*User, error)
}

// cache implements the Cacher interfaces using Redis
type cache struct {
	client *redis.Client
}

// NewCache creates a new cache instance
func NewCache(client *redis.Client) Cacher {
	return &cache{
		client: client,
	}
}

// Set stores a value in cache with the given key
func (c *cache) Set(key string, value interface{}) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("serialization error: %w", err)
	}
	return c.client.Set(key, jsonData, 24*time.Hour).Err() // Set TTL to 24 hours
}

// Get retrieves a value from cache by key
func (c *cache) Get(key string) (interface{}, error) {
	val, err := c.client.Get(key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("key not found: %s", key)
	}
	if err != nil {
		return nil, err
	}

	var result interface{}
	if err := json.Unmarshal([]byte(val), &result); err != nil {
		return nil, fmt.Errorf("deserialization error: %w", err)
	}
	return result, nil
}

// GetUser retrieves a user from cache and unmarshals it properly
func (c *cache) GetUser(key string) (*User, error) {
	val, err := c.client.Get(key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("user not found: %s", key)
	}
	if err != nil {
		return nil, err
	}

	var user User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, fmt.Errorf("user deserialization error: %w", err)
	}
	return &user, nil
}

// waitForRedis attempts to connect to Redis with retries
func waitForRedis(client *redis.Client) error {
	var err error
	maxAttempts := 5
	for i := 0; i < maxAttempts; i++ {
		_, err = client.Ping().Result()
		if err == nil {
			return nil
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("failed to connect to Redis after %d attempts: %v", maxAttempts, err)
}

// generateSampleUsers creates a list of sample users
func generateSampleUsers() []*User {
	now := time.Now().Format(time.RFC3339)
	return []*User{
		{
			ID:        1,
			Name:      "John Smith",
			Age:       30,
			Email:     "john.smith@example.com",
			Phone:     "+1234567890",
			CreatedAt: now,
		},
		{
			ID:        2,
			Name:      "Emily Johnson",
			Age:       28,
			Email:     "emily.j@example.com",
			Phone:     "+1987654321",
			CreatedAt: now,
		},
		{
			ID:        3,
			Name:      "Michael Brown",
			Age:       35,
			Email:     "michael.b@example.com",
			Phone:     "+1122334455",
			CreatedAt: now,
		},
		{
			ID:        4,
			Name:      "Sarah Davis",
			Age:       26,
			Email:     "sarah.d@example.com",
			Phone:     "+5566778899",
			CreatedAt: now,
		},
		{
			ID:        5,
			Name:      "David Wilson",
			Age:       42,
			Email:     "david.w@example.com",
			Phone:     "+4433221100",
			CreatedAt: now,
		},
		{
			ID:        6,
			Name:      "Jennifer Miller",
			Age:       31,
			Email:     "jennifer.m@example.com",
			Phone:     "+7788990011",
			CreatedAt: now,
		},
		{
			ID:        7,
			Name:      "Robert Taylor",
			Age:       29,
			Email:     "robert.t@example.com",
			Phone:     "+9900112233",
			CreatedAt: now,
		},
		{
			ID:        8,
			Name:      "Lisa Anderson",
			Age:       33,
			Email:     "lisa.a@example.com",
			Phone:     "+3344556677",
			CreatedAt: now,
		},
		{
			ID:        9,
			Name:      "Thomas Martinez",
			Age:       27,
			Email:     "thomas.m@example.com",
			Phone:     "+1122334455",
			CreatedAt: now,
		},
		{
			ID:        10,
			Name:      "Jessica Lee",
			Age:       24,
			Email:     "jessica.l@example.com",
			Phone:     "+6677889900",
			CreatedAt: now,
		},
	}
}

func main() {
	// Initialize Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Wait for Redis connection
	if err := waitForRedis(client); err != nil {
		log.Fatalf("Redis connection error: %v", err)
		os.Exit(1)
	}
	log.Println("Successfully connected to Redis")

	cache := NewCache(client)

	// Test basic cache operation
	if err := cache.Set("test:key", "sample value"); err != nil {
		log.Fatalf("Error setting value: %v", err)
	}

	// Store sample users in cache
	users := generateSampleUsers()
	for _, user := range users {
		userKey := fmt.Sprintf("user:%d", user.ID)
		if err := cache.Set(userKey, user); err != nil {
			log.Fatalf("Error saving user %d: %v", user.ID, err)
		}
		log.Printf("Saved user: %s (ID: %d)", user.Name, user.ID)
	}

	// Retrieve and display a user
	userID := 3
	userKey := fmt.Sprintf("user:%d", userID)
	user, err := cache.GetUser(userKey)
	if err != nil {
		log.Fatalf("Error getting user %d: %v", userID, err)
	}

	fmt.Printf("\nRetrieved user:\n")
	fmt.Printf("ID: %d\nName: %s\nAge: %d\nEmail: %s\nPhone: %s\nCreated At: %s\n",
		user.ID, user.Name, user.Age, user.Email, user.Phone, user.CreatedAt)
}
