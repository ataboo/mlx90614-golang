package main

// #include "bcm2835.h"
// #include <stdlib.h>
import "C"
import (
	"log"
	"os/exec"
	"fmt"
	"bytes"
	"strings"
	"strconv"
)

func main() {
	if err := connect(); err != nil {
		log.Fatal("Failed to connect")
	}


}

func connect() (int16, error) {
	cmd := exec.Command("i2cget", "-y", "1", "0x5a", "0x06", "w")

	out := bytes.Buffer{}
	stdErr := bytes.Buffer{}

	cmd.Stdout = &out
	cmd.Stderr = &stdErr

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stdErr.String())
		return 0, err
	}

	rawTemp := strings.Trim(out.String(), "\n")
	temp64, err := strconv.ParseInt(rawTemp, 0, 16)

	return int16(temp64), nil
}