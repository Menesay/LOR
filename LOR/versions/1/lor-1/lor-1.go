package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {

	go terminal()

	// not concurrent
	host, port := lor0HostPortServer()
	fmt.Print("[DATA] lor-0:", host, port)

	// Debug
	// host: menesay.duckdns.org
	//  (enter)

	// Debug
	// port: 80
	//  (enter)

	// lor-0 host:port connection is done
	// send host:port data to lor-2
	lor2HostPortClient(host, port)

	// start proxy server 127.0.0.20:1401
	lor0Proxy()

}

func lor0HostPortServer() (string, string) {

	// get host:port from lor-0

	listener, err := net.Listen("tcp", "127.0.0.20:1400")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("[INFO] lor-0 host:port server is running at 127.0.0.20:1400")

	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
	}

	//

	defer conn.Close()

	fmt.Println("[INFO] lor-0:", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)

	// host from lor-0
	host, err := reader.ReadString('\n')
	if err != nil {
		log.Println("[ERROR] reading host:", err)
	}
	fmt.Println("[DATA] host:", conn.RemoteAddr().String()+":", host)

	// send "\n" cuz lor-0 will be waiting for "\n" byte.
	_, err = conn.Write([]byte("\n"))

	// port from lor-0
	port, err := reader.ReadString('\n')
	if err != nil {
		log.Println("[ERROR] reading port:", err)
	}
	fmt.Println("[DATA] port:", conn.RemoteAddr().String()+":", port)

	// send "\n" cuz lor-0 will be waiting for "\n" byte.
	_, err = conn.Write([]byte("\n"))

	// after get host:port, disconnect from lor-0
	fmt.Println("[INFO] lor-0 disconnected:", conn.RemoteAddr().String())

	return host, port

}

func lor2HostPortClient(host string, port string) {

	// Removing "\n"
	//host = host[:len(host)-1]
	//port = port[:len(port)-1]

	// connect to lor-2 host:port
	conn, err := net.Dial("tcp", "127.0.0.30:1400")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// send host to lor-2
	_, err = conn.Write([]byte(host))
	if err != nil {
		log.Fatal("[ERROR] send host:", err)
	}

	// get \n from lor-2
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal("[ERROR] reading \\n", err)
	}

	// Debug
	fmt.Println("[DATA] lor-2:", response)

	//

	// send port to lor-2
	_, err = conn.Write([]byte(port))
	if err != nil {
		log.Fatal("[ERROR] send port:", err)
	}

	// get \n from lor-2
	response, err = bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal("[ERROR] reading \\n", err)
	}

	// Debug
	fmt.Println("[DATA] lor-2:", response)

}

func lor0Proxy() {

	// config

	// lor-2 ip
	PROXYHOST := "127.0.0.30"

	// lor-2 port
	PROXYPORT := "1401"

	proxyURL, err := url.Parse(fmt.Sprintf("http://%s:%s", PROXYHOST, PROXYPORT))
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(proxyURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	fmt.Println("[INFO] lor-0 proxy server is running at 127.0.0.20:1401")
	log.Fatal(http.ListenAndServe("127.0.0.20:1401", nil))

}

func terminal() {

	for {

		cmd := ""
		fmt.Scan(&cmd)

		if cmd == "exit" {
			os.Exit(0)
		}

	}

}
