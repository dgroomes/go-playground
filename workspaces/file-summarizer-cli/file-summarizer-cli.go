// A CLI runner for the core file summarizer library. This CLI sprinkles in some JSON syntax highlighting. This is an
// example of "application code" that uses "library code". It's useful to understand that contrast, and we can use Go
// workspaces to enforce it.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"os"
)

func main() {
	fileSummaries, err := file_summarizer_lib.SummarizeFiles()
	if err != nil {
		fmt.Println("Error summarizing files:", err)
		return
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(fileSummaries, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Syntax highlight with chrome
	lexer := lexers.Get("json")
	if lexer == nil {
		fmt.Printf("Lexer not found\n")
		return
	}

	style := styles.Get("github")
	if style == nil {
		fmt.Printf("Style not found\n")
		return
	}

	formatter := formatters.Get("terminal")
	if formatter == nil {
		fmt.Printf("Formatter not found\n")
		return
	}

	// Tokenize and format the JSON data
	iterator, err := lexer.Tokenise(nil, string(jsonData))
	if err != nil {
		fmt.Println("Error tokenizing JSON:", err)
		return
	}

	err = formatter.Format(os.Stdout, style, iterator)
	if err != nil {
		fmt.Println("Error formatting output:", err)
		return
	}
}
