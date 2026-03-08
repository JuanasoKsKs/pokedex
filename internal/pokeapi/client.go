package pokeapi

import (
	"time"
	"net/http"
	"github.com/JuanasoKsKs/pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	mycache pokecache.Cache
	Mypokedex pokedex
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		mycache: pokecache.NewCache(cacheInterval),
		Mypokedex: NewPokedex(),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}	
}