package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func v1() {
	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out

	err := prc.Run()
	if err != nil {
		fmt.Println(err)
	}

	if prc.ProcessState.Success() {
		fmt.Println("'ls -a' rocess output v1:")
		fmt.Println(out.String())
	}
}

func v2() {
	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out

	err := prc.Start()
	if err != nil {
		fmt.Println(err)
	}

	prc.Wait()

	if prc.ProcessState.Success() {
		fmt.Println("'ls -a' rocess output v2:")
		fmt.Println(out.String())
	}
}

func main() {
	v1()
	v2()
}
