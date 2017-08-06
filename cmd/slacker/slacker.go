package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/hwong/slacker"
)

func main() {
	config, err := slacker.GetConfig()
	if err != nil {
		panic(err)
	}

	message := strings.Join(os.Args[1:], " ")
	if len(message) == 0 {
		reader := bufio.NewReader(os.Stdin)
		read, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Printf("Error reading from stdin: %v\n.", err)
			return
		}
		message = string(read)
	}

	slacker.Post(config, message)
	fmt.Println(message)
}
