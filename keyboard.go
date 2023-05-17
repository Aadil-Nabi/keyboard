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

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	trimmedNumber := strings.TrimSpace(input)

	num, err := strconv.Atoi(trimmedNumber)
	if err != nil {
		log.Fatal(err)
	}

	return num
}
