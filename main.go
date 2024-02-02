package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os/exec"
	"strconv"
)

var gr = color.New(color.FgGreen).Add(color.Bold)
var red = color.New(color.FgRed).Add(color.Bold)

func createInterface(a int) {
	for i := 0; i < a; i++ {
		currInterface := `echo ` + strconv.Itoa(i)
		cmd := exec.Command("bash", "-c", currInterface)
		_, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		gr.Println("Created interface no. " + strconv.Itoa(i))
	}
}

func main() {
	fmt.Println("a")
	c := color.New(color.FgGreen).Add(color.Bold)
	c.Println("Creating interface..")

	createInterface(3)
}
