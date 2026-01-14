package main


import (

	"fmt"
	"os"
	"go-reloaded/internal/tokenizer"
	"go-reloaded/internal/pipeline"

)

func main(){

	// Checks for the correct number of args
	if len(os.Args) =! 3 (
		fmt.Println("Error: Input the <input_File> <output_File>")
		os.Exit(1)
	)
	input_File := os.Args[1]
	output_File := os.Args[2]

	inputbytesrep, err:= os.ReadFile(input_File)

	if err != nill (
		fmt.Printf("Error: in the file as described: %v\n", err)
		os.Exit[1]
	)

	input_text := string(inputbytesrep)

	tokens, err := tokenizer.Tokenize(input_text)

	if err != nill {

		fmt.Printf("Erro in the input text as diplayed: %v\n", err)
		os.Exit[1]

	}

		

	resultToken := pipeline.Run(tokens)

	output_Text := tokenizer.Join(resultToken)

	// writes the new updates text into the output file
	err := os.WriteFile(output_File, []byte(output_Text), 0644)
	if err != nill {

		fmt.Printf("Error Writing File: %v\n", err)
	}

}