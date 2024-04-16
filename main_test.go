package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "capitalization",
			input:    "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			expected: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{
			name:     "number conversion",
			input:    "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			expected: "Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			name:     "articles",
			input:    "There is no greater agony than bearing a untold story inside you.",
			expected: "There is no greater agony than bearing an untold story inside you.",
		},
		{
			name:     "punctuations",
			input:    "Punctuation tests are ... kinda boring ,don't you think !?",
			expected: "Punctuation tests are... kinda boring, don't you think!?",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary input and output files
			inputFile, err := ioutil.TempFile("", "input-*.txt")
			if err != nil {
				t.Errorf("Failed to create temporary input file: %v", err)
				return
			}
			defer os.Remove(inputFile.Name())

			outputFile, err := ioutil.TempFile("", "output-*.txt")
			if err != nil {
				t.Errorf("Failed to create temporary output file: %v", err)
				return
			}
			defer os.Remove(outputFile.Name())

			// Write input data to the temporary input file
			_, err = inputFile.WriteString(tt.input)
			if err != nil {
				t.Errorf("Failed to write input data to temporary file: %v", err)
				return
			}

			// Run the main function with the temporary files
			os.Args = []string{"test", inputFile.Name(), outputFile.Name()}
			main()

			// Read the output file and compare with the expected output
			outputData, err := ioutil.ReadFile(outputFile.Name())
			if err != nil {
				t.Errorf("Failed to read output data from temporary file: %v", err)
				return
			}

			if string(outputData) != tt.expected+"\n" {
				t.Errorf("Output mismatch: got %q, want %q", string(outputData), tt.expected+"\n")
			}
		})
	}
}
