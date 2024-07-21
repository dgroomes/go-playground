package main

import (
	"encoding/json"
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(`
Here is one of my first Go programs. It does the following:
    * It takes some quotes (sayings) from the "rsc.io/quote" package' 
    * Puts them in a slice
    * Marshals (serializes) the slice into a JSON document
    * Unmarshalls (deserializes) it back to a slice and prints it!`, "\n")

	quotes := []string{
		quote.Hello(),
		quote.Go(),
		quote.Glass(),
		quote.Opt(),
	}

	// Marshall the sayings into a JSON document
	marshalled, _ := json.MarshalIndent(quotes, "", "  ")
	fmt.Printf("Here are the sayings in an indented JSON document:\n%s\n\n", marshalled)

	// Unmarshall the JSON document back into a string slice
	var unmarshalled []string
	_ = json.Unmarshal(marshalled, &unmarshalled)
	fmt.Printf("Here are the sayings as a string slice formatted by the 'fmt' package and the '%%q' format verb:\n%q\n\n", unmarshalled)
}
