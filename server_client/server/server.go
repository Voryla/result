package main

import (
	"bufio"
	"net"
	"regexp"
	"strconv"
	"strings"
)

type Config struct {
	TcpNoDelay bool
	Addr       string
	Protocol   string
}

func main() {
	config := Config{
		TcpNoDelay: false,
		Addr:       ":15551",
		Protocol:   "tcp",
	}
	doServer(config)
}

func doServer(config Config) {
	regexpObj := regexp.MustCompile(`^\d+\+\d+$`)
	listener, err := net.Listen(config.Protocol, config.Addr)
	if err != nil {
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		go func(conn net.Conn) {
			defer conn.Close()
			writer := bufio.NewWriter(conn)
			reader := bufio.NewReader(conn)
			// control noDelay
			writer.WriteString(strconv.FormatBool(config.TcpNoDelay) + "\n")
			writer.Flush()
			readLine, err := reader.ReadString('\n')
			if err != nil {
				return
			}
			readLine = strings.Replace(readLine, "\n", "", -1)
			if regexpObj.MatchString(readLine) {
				data := strings.Split(readLine, "+")
				a, _ := strconv.Atoi(data[0])
				b, _ := strconv.Atoi(data[0])
				writer.WriteString(strconv.Itoa(a+b) + "\n")
				writer.Flush()
				if err != nil {
					return
				}
			} else {
				writer.WriteString("0\n")
				writer.Flush()
			}
		}(conn)
	}
}
