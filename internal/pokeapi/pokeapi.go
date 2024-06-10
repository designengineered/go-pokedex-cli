package pokeapi

import (
	"github.com/laztaxon/go-pokedex-cli/internal/pokecache"
	"net/http"
	"time"
)

// main api url
const baseURL = "https://pokeapi.co/api/v2/"

type Client struct {
	// implementing a pointer to the cache interface
	cache      *pokecache.Cache
	httpClient http.Client
}

// pass in cacheInterval to the constructor
func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
