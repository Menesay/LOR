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
	host, port := lor1HostPortServer()

	fmt.Print("[DATA] lor-1:", host, port)

	// Debug
	// host: menesay.duckdns.org
	//  (enter)

	// Debug
	// port: 80
	//  (enter)

	// lor-1 host:port connection is done
	// start proxy server 127.0.0.30:1401
	lor1Proxy(host, port)

}

func lor1HostPortServer() (string, string) {

	// get host:port from lor-1

	listener, err := net.Listen("tcp", "127.0.0.30:1400")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("[INFO] lor-1 host:port server is running at 127.0.0.30:1400")

	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
	}

	//

	defer conn.Close()

	fmt.Println("[INFO] lor-1:", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)

	// host from lor-1
	host, err := reader.ReadString('\n')
	if err != nil {
		log.Println("[ERROR] reading host:", err)
	}
	fmt.Println("[DATA] host:", conn.RemoteAddr().String()+":", host)

	// send "\n" cuz lor-1 will be waiting for "\n" byte.
	_, err = conn.Write([]byte("\n"))

	// port from lor-1
	port, err := reader.ReadString('\n')
	if err != nil {
		log.Println("[ERROR] reading port:", err)
	}
	fmt.Println("[DATA] port:", conn.RemoteAddr().String()+":", port)

	// send "\n" cuz lor-1 will be waiting for "\n" byte.
	_, err = conn.Write([]byte("\n"))

	// after get host:port, disconnect from lor-1
	fmt.Println("[INFO] lor-1 disconnected:", conn.RemoteAddr().String())

	return host, port

}

func lor1Proxy(host string, port string) {

	// Removing "\n"
	host = host[:len(host)-1]
	port = port[:len(port)-1]

	proxyURL, err := url.Parse(fmt.Sprintf("http://%s:%s", host, port))
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(proxyURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	fmt.Println("[INFO] lor-1 proxy server is running at 127.0.0.30:1401")
	log.Fatal(http.ListenAndServe("127.0.0.30:1401", nil))

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
