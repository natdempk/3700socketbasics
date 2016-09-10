package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	portPtr := flag.Int("p", 27993, "port")
	sslPtr := flag.Bool("s", false, "SSL enabled")

	flag.Parse()

	port := *portPtr
	ssl := *sslPtr

	hostName := flag.Args()[0]
	neuId := flag.Args()[1]

	var conn net.Conn

	if ssl {
		config := &tls.Config{
			InsecureSkipVerify: true,
		}
		conn, _ = tls.Dial("tcp", hostName+":"+strconv.Itoa(port), config)
	} else {
		conn, _ = net.Dial("tcp", hostName+":"+strconv.Itoa(port))
	}
	fmt.Fprintf(conn, "cs3700fall2016 HELLO %s\n", neuId)

	reader := bufio.NewReader(conn)

	for {
		response, _ := reader.ReadString('\n')
		solution, secret := parse(response)
		if secret != "" {
			fmt.Println(secret)
			return
		} else {
			fmt.Fprintf(conn, "cs3700fall2016 %d\n", solution)
		}
	}

}

func parse(response string) (int, string) {
	splitResponse := strings.Split(strings.TrimSpace(response), " ")
	if len(splitResponse) == 5 {
		x, _ := strconv.Atoi(splitResponse[2])
		y, _ := strconv.Atoi(splitResponse[4])
		return math(x, y, splitResponse[3]), ""
	} else { // found solution
		return 0, splitResponse[2]
	}
}

func math(x int, y int, operator string) int {
	switch operator {
	case "+":
		return x + y
	case "-":
		return x - y
	case "/":
		return x / y
	case "*":
		return x * y
	}

	return 0
}