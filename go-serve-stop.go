package main

import (
	"os"
)

func main() {

	    os.StartProcess("./res/gosuv.exe", []string{"./res/gosuv.exe", "stop"}, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
		

		//select{}


		// reader := bufio.NewReader(os.Stdin)
		// fmt.Println("Simple Shell")
		// fmt.Println("---------------------")

		// for {
		// 	fmt.Print("-> ")
		// 	text, _ := reader.ReadString('\n')
		// 	// convert CRLF to LF
		// 	text = strings.Replace(text, "\n", "", -1)

		// 	if strings.Compare("hi", text) == 0 {
		// 		fmt.Println("hello, Yourself")
		// 	}

		// }
		
}