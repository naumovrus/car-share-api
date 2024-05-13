package balancetracker

import (
	"math/rand"
	"sync"
	"time"
)

type Client struct {
	cache map[string]float64
	mu    *sync.Mutex
}

func New() *Client {
	cache := map[string]float64{}
	mu := &sync.Mutex{}
	return &Client{cache: cache, mu: mu}
}

func (c *Client) GetBalance(clientUUID string) float64 {
	rand.Seed(time.Now().UnixNano())
	randomFloat := rand.Float64()
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.cache[clientUUID]; !ok {
		c.cache[clientUUID] = randomFloat * 10000
		return c.cache[clientUUID]
	}
	c.cache[clientUUID] = randomFloat * 10000
	return c.cache[clientUUID]
}
