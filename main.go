package main

<<<<<<< Updated upstream
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		input = cleanInput(input)[0]
		fmt.Printf("Your command was: %v\n", input)
	}
=======
func main() {
	startRepl()
>>>>>>> Stashed changes
}
