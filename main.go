package main

import (
	"time"
	"github.com/JuanasoKsKs/pokedex/internal/pokeapi"
)

func main(){
	pokeClient := pokeapi.NewClient(5 * time.Second)
	conf := &config{
		myclient: pokeClient,
	}
	startRepl(conf)
}

