package main
import (
	"fmt"
	"os"
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

	
