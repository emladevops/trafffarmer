package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func hello() {
	fmt.Println("Hello")
}

func prepare() {
	gr.Println("Preparing system for traffmonetizer...")
	exec.Command("date")
	gr.Println("Cloning xjasonlyu/tun2socks...")
	out, err := exec.Command("bash", "-c", "git clone https://github.com/xjasonlyu/tun2socks.git").Output()
	if err != nil {
		log.Fatal(err)
	}
	gr.Println(string(out))

	gr.Println("Building tun2socks...")
	exec.Command("apt install -y build-essential")

	currDir, _ := os.Getwd()
	os.Chdir(currDir + "/tun2socks")

	out, err = exec.Command("bash", "-c", "make tun2socks").Output()
	if err != nil {
		log.Fatal(err)
	}
	gr.Println(string(out))

	gr.Println("Copying to /usr/local/bin..")
	exec.Command("bash", "-c", "cp /root/trafffarmer/build/tun2socks /usr/local/bin")
}
