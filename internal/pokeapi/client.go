package pokeapi

import (
	"net/http"
	"time"

	"github.com/bpeter340/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(requestTTL time.Duration, cacheTTL time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: requestTTL,
		},
		cache: pokecache.NewCache(cacheTTL),
	}
}
