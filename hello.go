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
    * Marshals (serializes) the slice into JSON (as bytes, how do I get a JSON string?)
    * Unmarshalls (deserializes) it back to a slice and prints it!`)

	quotes := []string{
		quote.Hello(),
		quote.Go(),
		quote.Glass(),
		quote.Opt(),
	}

	marshalled, _ := json.Marshal(quotes)

	var unmarshalled []string
	_ = json.Unmarshal(marshalled, &unmarshalled)

	fmt.Printf("\nHere are the sayings: %q", unmarshalled)
}
