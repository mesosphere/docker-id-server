package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type server struct {
	id string
}

const defaultPort int = 80
const idSpace = "abcdefghijklmnopqrstuvwxyz0123456789"

func main() {
	port := 0
	s := server{}

	if len(os.Args) >= 2 {
		var err error
		rawPort := os.Args[1]
		port, err = strconv.Atoi(rawPort)
		if err != nil {
			log.Fatalf("could not parse port \"%s\": %s", rawPort, err)
		}
	}
	if len(os.Args) >= 3 {
		s.id = os.Args[2]
	}
	if len(os.Args) > 3 {
		log.Fatal("too many arguments")
	}

	if port == 0 {
		port = defaultPort
	}
	if s.id == "" {
		rand.Seed(time.Now().UTC().UnixNano())
		s.id = randStringBytes(12)
	}

	log.Printf("listening on port %d", port)
	log.Printf("serving id %s", s.id)

	http.HandleFunc("/id", s.defaultHandler)

	addr := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = idSpace[rand.Intn(len(idSpace))]
	}
	return string(b)
}

func (s server) defaultHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(s.id)); err != nil {
		log.Print(err)
	}
}
