// Package main implements a plaintext to morse-code encoder
package main

// import required packages
import (
	"fmt"
	"errors"
	"flag"
	"os"
	"strings"
	"bytes"
)

// errorHandler prints out(and terminates) an error according to the error-log argument
func errorHandler(errorLog string) {
	err := errors.New(errorLog)
	fmt.Println(err)
	os.Exit(1)
}

// translateInput parses the STDIN data and translates it into morse code
func translateInput(args []string) {
	// map that holds character to morse code mapping
	reverseMap := make(map[string]string)

	// Alphabates
	reverseMap["A"] = ".-"
	reverseMap["B"] = "-..."
	reverseMap["C"] = "-.-."
	reverseMap["D"] = "-.."
	reverseMap["E"] = "."
	reverseMap["F"] = "..-."
	reverseMap["G"] = "--."
	reverseMap["H"] = "...."
	reverseMap["I"] = ".."
	reverseMap["J"] = ".---"
	reverseMap["K"] = "-.-"
	reverseMap["L"] = ".-.."
	reverseMap["M"] = "--"
	reverseMap["N"] = "-."
	reverseMap["O"] = "---"
	reverseMap["P"] = ".--."
	reverseMap["Q"] = "--.-"
	reverseMap["R"] = ".-."
	reverseMap["S"] = "..."
	reverseMap["T"] = "-"
	reverseMap["U"] = "..-"
	reverseMap["V"] = "...-"
	reverseMap["W"] = ".--"
	reverseMap["X"] = "-..-"
	reverseMap["Y"] = "-.--"
	reverseMap["Z"] = "--.."

	// Decimals
	reverseMap["1"] = ".----"
	reverseMap["2"] = "..---"
	reverseMap["3"] = "...--"
	reverseMap["4"] = "....-"
	reverseMap["5"] = "....."
	reverseMap["6"] = "-...."
	reverseMap["7"] = "--..."
	reverseMap["8"] = "---.."
	reverseMap["9"] = "----."
	reverseMap["0"] = "-----"

	// Punctuation marks and miscellaneous signs
	reverseMap["."] = ".-.-.-"
	reverseMap[","] = "--..--"
	reverseMap[":"] = "---..."
	reverseMap["?"] = "..--.."
	reverseMap["'"] = ".----."
	reverseMap["-"] = "-....-"
	reverseMap["/"] = "-..-."
	reverseMap["("] = "-.--."
	reverseMap[")"] = "-.--.-"
	reverseMap["+"] = ".-.-."
	reverseMap["×"] = "-..-"
	reverseMap["@"] = ".--.-."

	// buffer holding resultant string
	var result bytes.Buffer

	// no text provided to translate
	if len(args) == 0 {
		errorHandler("Usage: morser -encode PLAINTEXT")
	}

	// iterate over each word of the input text
	for i := 0; i < len(args); i++ {

		// separate word representations in morse code by a backslace
		if i >= 1 {
			result.WriteString(" / ")
		}

		// iterate over each character of the word
		for j := 0; j < len(args[i]); j++ {

			// separate character representations in morse code by a space
			if j >= 1 {
				result.WriteString(" ")
			}

			// validates the input text, character by character
			_, isPresent := reverseMap[strings.ToUpper(string(args[i][j]))]

			// invalid input provided
			if !isPresent {
				errorHandler("Invalid input. Only alphabates, decimals and punctuation marks are supported.")
			}

			result.WriteString(reverseMap[strings.ToUpper(string(args[i][j]))])
		}
	}

	fmt.Println(result.String())
}

func main() {
	// defines a command line flag that looks for encoding permission
	encodeFlag := flag.Bool("encode", false, "morser -encode [arg] : Encodes a plaintext string (argument) to morse code")

	// parse the command line arguments
	flag.Parse()

	// process ahead if the 'encode' flag is available in arguments
	if *encodeFlag {
		translateInput(flag.Args())
	} else {
		errorHandler("Usage: morser -encode PLAINTEXT")
	}
}
