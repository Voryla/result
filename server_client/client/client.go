package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

func main() {
	raddr, err := net.ResolveTCPAddr("tcp", ":15551")
	if err != nil {
		return
	}
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		return
	}
	reader := bufio.NewReader(conn)
	readLine, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	isNoDelay, _ := strconv.ParseBool(readLine)
	err = conn.SetNoDelay(isNoDelay)
	if err != nil {
		log.Fatal(err)
	}
	express := ""
	fmt.Scan(&express)
	writer := bufio.NewWriter(conn)
	writer.WriteString(express + "\n")
	writer.Flush()
	answer, err := reader.ReadString('\n')
	fmt.Println(answer)
	defer conn.Close()
}
