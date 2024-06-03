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
	host, port := ncHostPortServer()
	fmt.Println("[DATA] nc:", host, port)

	// nc host:port connection is done
	// send host:port data to lor-1
	lor1HostPortClient(host, port)

	// start proxy server 127.0.0.30:1401
	ncProxy()

}

func ncHostPortServer() (string, string) {

	// get host:port from nc

	listener, err := net.Listen("tcp", "127.0.0.10:1400")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("[INFO] nc host:port server is running at 127.0.0.10:1400")

	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
	}

	//

	defer conn.Close()

	fmt.Println("[INFO] nc:", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)

	// host from nc
	host, err := reader.ReadString('\n')
	if err != nil {
		log.Println("[ERROR] reading host:", err)
	}
	fmt.Println("[DATA] nc host:", conn.RemoteAddr().String()+":", host)
	_, err = conn.Write([]byte("ok\n"))

	// port from nc
	port, err := reader.ReadString('\n')
	if err != nil {
		log.Println("[ERROR] reading port:", err)
	}
	fmt.Println("[DATA] nc port:", conn.RemoteAddr().String()+":", port)
	_, err = conn.Write([]byte("ok\n"))

	// after get host:port, disconnect from nc
	fmt.Println("[INFO] nc disconnected:", conn.RemoteAddr().String())

	return host, port

}

func lor1HostPortClient(host string, port string) {

	// Removing "\n"
	//host = host[:len(host)-1]
	//port = port[:len(port)-1]

	// connect to lor-1 host:port
	conn, err := net.Dial("tcp", "127.0.0.20:1400")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// send host to lor-1
	_, err = conn.Write([]byte(host))
	if err != nil {
		log.Fatal("[ERROR] send host:", err)
	}

	// get \n from lor-1
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal("[ERROR] reading \\n", err)
	}

	// Debug
	fmt.Println("[DATA] lor-1:", response)

	//

	// send port to lor-1
	_, err = conn.Write([]byte(port))
	if err != nil {
		log.Fatal("[ERROR] send port:", err)
	}

	// get \n from lor-1
	response, err = bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal("[ERROR] reading \\n", err)
	}

	// Debug
	fmt.Println("[DATA] lor-1:", response)

}

func ncProxy() {

	// config

	// lor-1 ip
	PROXYHOST := "127.0.0.20"

	// lor-1 port
	PROXYPORT := "1401"

	proxyURL, err := url.Parse(fmt.Sprintf("http://%s:%s", PROXYHOST, PROXYPORT))
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(proxyURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	fmt.Println("[INFO] lor-0 proxy server is running at 127.0.0.10:1401")
	log.Fatal(http.ListenAndServe("127.0.0.10:1401", nil))

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
