package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"
)

type Response struct {
	ServiceId  *string `json:"service_id"`
	IPAddress  string  `json:"ip_address"`
	TimeServer time.Time
}

var serviceId *string
var servicePort *int

func main() {
	serviceId = flag.String("name", fmt.Sprintf("%v", rand.Intn(100)), "Service ID")
	servicePort = flag.Int("port", 8840, "Port Service")
	flag.Parse()

	http.HandleFunc("/", getServer)

	// Start the HTTP server
	fmt.Printf("Server is running on :%d\n", *servicePort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *servicePort), nil))
}

func getServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ip, err := getCurrentIPAddress()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	respJson := Response{
		ServiceId:  serviceId,
		IPAddress:  ip,
		TimeServer: time.Now(),
	}

	json.NewEncoder(w).Encode(respJson)
}

func getCurrentIPAddress() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String(), nil
}
