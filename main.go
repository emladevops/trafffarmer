package main

import (
	"errors"
	"github.com/fatih/color"
	"log"
	"os"
	"os/exec"
	user2 "os/user"
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
		gr.Println("Created interface tun" + strconv.Itoa(i))
	}
}

// 3 stages

func main() {

	user, err := user2.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	username := user.Username

	if username != "root" {
		red.Println("Run me as root!")
		panic(0)
	}

	if len(os.Args) > 1 {
		red.Println("Warming up...")

	} else {
		if _, err := os.Stat("/usr/traffarmer"); err == nil {
			gr.Println("Starting the container, you've built the package before")
		} else if errors.Is(err, os.ErrNotExist) {
			red.Println("Oops, initializing...")

		}
	}

}
