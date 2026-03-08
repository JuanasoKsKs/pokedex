package main
import (
	"fmt"
	"os"
	"math/rand"
)

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(conf *config) error {
	locationsResp, err := conf.myclient.ListLocations(conf.next)
	if err != nil {
		return err
	}
	conf.next = locationsResp.Next
	conf.previous = locationsResp.Previous

	for _, result := range locationsResp.Results {
		println(result.Name)
	}

	return nil
}

func commandMapb(conf *config) error {
	if conf.previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}
	locationsResp, err := conf.myclient.ListLocations(conf.previous)
	if err != nil {
		return err
	}
	conf.next = locationsResp.Next
	conf.previous = locationsResp.Previous

	for _, result := range locationsResp.Results {
		println(result.Name)
	}

	return nil
}

func commandExplore(conf *config) error {
	location, err := conf.myclient.GetLocation(conf.argument)
	if err != nil {
		return err
	}
	fmt.Println("----------Pokemons in the area: ---------")
	for _, encounter := range location.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(conf *config) error {
	pokemon, err := conf.myclient.GetPokemon(conf.argument)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	randomResult := rand.Intn(10)
	if randomResult > 5 {
		fmt.Printf("%s was caught!  *******\n", pokemon.Name)
		conf.myclient.Mypokedex.Collection[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped! xxxxxx\n", pokemon.Name)
	}




	return nil
}

func commandInspect(conf *config) error {
	pokemon, ok := conf.myclient.Mypokedex.Collection[*conf.argument]
	if !ok {
		return fmt.Errorf("------you have not caught that pokemon yet----")
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, ty := range pokemon.Types {
		fmt.Printf("- %v\n", ty.Type.Name)
	}
	return nil

}

func commandPokedex(conf *config) error {
	fmt.Println("------ Your Pokedex ------")
	for _, pokemon := range conf.myclient.Mypokedex.Collection {
		fmt.Println(pokemon.Name)
	}
	return nil
}