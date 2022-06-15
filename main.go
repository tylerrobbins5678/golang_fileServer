package main

import (
	"bufio"
	filecontroller "fileManager/file/controller"
	"fmt"
	"os"
)

func main() {
	go filecontroller.Start()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	reader.ReadString('\n')
}
