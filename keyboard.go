package keyboard

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInput() int {

	fmt.Println("Enter the Number")

	// Create a new reader to read data from the terminal
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	// Trim the white space that comes when user hits enter key
	trimmedNumber := strings.TrimSpace(input)

	// Atoi to parse the string into an int
	num, err := strconv.Atoi(trimmedNumber)
	if err != nil {
		log.Fatal(err)
	}

	return num
}
