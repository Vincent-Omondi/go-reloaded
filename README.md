# Table of Contents

- [go-reloaded](#go-reloaded)
- [Features](#features)
- [Installation](#installation)
  - [Cloning the Repository](#cloningthe-repository)
  - [Navigate to the Project Directory](#navigate-to-the-project-directory)
- [Usage](#usage)
- [File Structure](#file-structure)
- [Test Cases](#test-cases)

## go-reloaded

go-reloaded is a simple and basic text completion, editing, and auto-correction tool written in Go. This tool allows you to perform various modifications on a given input text file and saves the modified text to an output file.

## Features

The program performs the following modifications on the input text:

- Replace every instance of `(hex)` with the decimal version of the word before it (assuming the word is a hexadecimal number).
- Replace every instance of `(bin)` with the decimal version of the word before it (assuming the word is a binary number).
- Convert the word before `(up)` to uppercase.
- Convert the word before `(low)` to lowercase.
- Convert the word before `(cap)` to capitalized form.
- If a number appears after `(low)`, `(up)`, or `(cap)` like `(low, <number>)`, it converts the specified number of words to lowercase, uppercase, or capitalized accordingly.
- Move punctuations (`.`, `,`, `!`, `?`, `:`, and `;`) closer to the previous word and add a space before the next word.
- Handle groups of punctuations like `...` or `!?` correctly.
- Place the `'` punctuation mark around the word(s) between them without any spaces.
- Convert `a` to `an` if the next word begins with a vowel (`a`, `e`, `i`, `o`, `u`) or `h`.

## Installation

1. Clone the repository:

    ```bash
    git clone https://learn.zone01kisumu.ke/git/vinomondi/go-reloaded.git
    ```

2. Navigate to the project directory:

    ```bash
    cd go-reloaded
    ```

## Usage

To run the program, use the following command:

```bash
go run . sample.txt result.txt
```

## File Structure

- `main.go`: The main entry point of the program.
- `handlePunctuations/handlepuncts.go`: Contains functions to handle punctuation marks.
- `handleArticles/handleArticles.go`: Contains functions to handle articles (`a/an`).
- `transformWords/transform.go`: Contains functions to transform words (uppercase, lowercase, capitalize, convert hex/bin to decimal).
- `transformWords/toupper.go`: Contains the `ToUpper` function to convert a string to uppercase.
- `transformWords/tolower.go`: Contains the `ToLower` function to convert a string to lowercase.
- `main_test.go`: Contains unit tests for the project.

## Test Cases

```console
$ cat sample.txt
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
$ go run . sample.txt result.txt
$ cat result.txt
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
$ cat sample.txt
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
$ go run . sample.txt result.txt
$ cat result.txt
Simply add 66 and 2 and you will see the result is 68.
$ cat sample.txt
There is no greater agony than bearing a untold story inside you.
$ go run . sample.txt result.txt

$ cat result.txt
There is no greater agony than bearing an untold story inside you.
$ cat sample.txt
Punctuation tests are ... kinda boring ,don't you think !?

$ go run . sample.txt result.txt
$ cat result.txt
Punctuation tests are... kinda boring, don't you think!?
```

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

