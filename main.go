package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"

	"github.com/grongor/panicwatch"
)

func main() {
	err := panicwatch.Start(panicwatch.Config{
		OnPanic: func(p panicwatch.Panic) {
			cmd := exec.Command("ls")
			cmd.Dir = "."

			var stdo, stde bytes.Buffer
			cmd.Stdout = &stdo
			cmd.Stderr = &stde
			err := cmd.Run()

			writeFile("stdout.txt", stdo.String())
			writeFile("stderr.txt", stde.String())
			writeFile("err.txt", err.Error())
		},
	})
	if err != nil {
		log.Fatalln("failed to start panicwatch: " + err.Error())
	}

	panic("test panic")
}

// write the file
func writeFile(filename string, file string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(file)

	if err != nil {
		return err
	}

	return nil
}
