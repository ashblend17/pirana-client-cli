package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/The-Alfred-Project/alfred/ahhhhh"
	"github.com/The-Alfred-Project/alfred/ter"
	"golang.org/x/term"
)

/* payload := []byte(`
				{
					"Roll": ["RNO123456"],    	-> array of strings
					"AlfredPassword": "password", 	-> string
					"Flags" : "scp"					-> string
				}
			`)

---------------------------------------------------------------------
	Response:

*/

/*
payload := []byte(`
	{
		"alfredPassword": password,
		"name": "john"
	}
`)
---------------------------------------------------------------------
Response:
[
    {
        "name": "John Doe",
        "roll": "RNO123456",
    }
]
*/

func main() {
	if len(os.Args) < 1 {
		ter.Help()
		os.Exit(1)
	} else if len(os.Args) < 3 {
		fmt.Println("Incomplete Command")
		ter.Help()
		os.Exit(1)
	}
	fmt.Print("Enter pass:")
	password, err := term.ReadPassword(int(syscall.Stdin))
	if !ahhhhh.CheckPass(string(password)) {
		fmt.Println("Incorrect password")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
	}
	flag.Parse()

	mainCommandFlag := flag.Arg(0)
	fmt.Println("Executing " + mainCommandFlag + "...")
	rolls := flag.Args()[1:]

	if strings.Contains(mainCommandFlag, "help") {
		ter.Help()
	} else if strings.Contains(mainCommandFlag, "find") {
		ahhhhh.GetNames(string(password), rolls[0])
	} else if strings.Contains(mainCommandFlag, "s") || strings.Contains(mainCommandFlag, "c") || strings.Contains(mainCommandFlag, "p") {
		ahhhhh.GetData(rolls, mainCommandFlag, string(password))
	} else {
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}
