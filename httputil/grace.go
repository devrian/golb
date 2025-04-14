package httputil

import (
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	graceful "gopkg.in/tylerb/graceful.v1"
)

var (
	cfgtestFlag bool
	listenPort  string
)

func checkConfigTest() {
	if cfgtestFlag {
		log.Println("config test mode, exiting")
		os.Exit(0)
	}
}

func Listen(hport string) (net.Listener, error) {
	var l net.Listener

	fd := os.Getenv("EINHORN_FDS")
	if fd != "" {
		sock, err := strconv.Atoi(fd)
		if err == nil {
			hport = "socketmaster:" + fd
			log.Println("detected socketmaster, listening on", fd)
			file := os.NewFile(uintptr(sock), "listener")
			fl, err := net.FileListener(file)
			if err == nil {
				l = fl
			}
		}
	}

	if listenPort != "" {
		hport = ":" + listenPort
	}

	checkConfigTest()

	if l == nil {
		var err error
		l, err = net.Listen("tcp4", hport)
		if err != nil {
			return nil, err
		}
	}

	return l, nil
}

func Serve(hport string, handler http.Handler, gracefulTimeout, readTimeout, writeTimeout time.Duration) error {
	checkConfigTest()

	l, err := Listen(hport)
	if err != nil {
		log.Fatalln(err)
	}

	srv := &graceful.Server{
		Timeout: gracefulTimeout,
		Server: &http.Server{
			Handler:      handler,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		},
	}

	log.Println("starting serve on ", hport)

	return srv.Serve(l)
}
