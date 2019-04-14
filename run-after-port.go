package main

import (
	"fmt"
	"net"
	"strconv"
    "time"
	"bufio"
	"io"
	"os/exec"
)

func main() {

	ip := net.ParseIP("127.0.0.1")
	
	port, err1 := strconv.Atoi("8761")
	
	fmt.Println("Checking 8761 . . . ")
	

	if err1 != nil{}
	
	tcpAddr := net.TCPAddr{
		IP:   ip,
		Port: port,
	}
	

	conn, err := net.DialTCP("tcp", nil, &tcpAddr)

	if err == nil {
	
		printOpeningPort(port)
		
		conn.Close()
		
		fmt.Println("Starting Jar...")
		
		command := "java"
		
		params := []string{"-jar","jwt.jar"}
		
		//执行cmd命令
		execCommand(command, params)

	}else{
		
		fmt.Println("Port Closed ... Waiting 5s ... ")
		
		time.Sleep(5 * time.Second)
		
		main()
	
	}
		
}

// 端口开放打印

func printOpeningPort(port int) {

	fmt.Println("port " + strconv.Itoa(port) + " is opening")

}


func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}
	
	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true
}


