package main

import (
	"os"
)

func main() {
	    os.StartProcess("./res/browser.exe", []string{"browser.exe", "--url=http://localhost:11313/"}, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
}