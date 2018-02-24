package main

import (
	"github.com/armon/go-socks5"
	"log"
	"os"
	"flag"
	"fmt"
)

var (
	build string

	l = log.New(os.Stdout, "[socks5-server]", log.Lshortfile)

	port = flag.Int("port", 1080, "socks5 tcp port")
	user = flag.String("u", "guest", "user")
	password = flag.String("p", "guest", "guest")

	version = flag.Bool("version", false, "show version")
)

func main() {
	flag.Parse()
	if !flag.Parsed() {
		flag.Usage()
		return
	}

	if *version {
		fmt.Println("socks5-server build:", build)
		return
	}

	config := &socks5.Config{
		AuthMethods: []socks5.Authenticator{socks5.UserPassAuthenticator{
			Credentials: socks5.StaticCredentials{
				*user: *password,
			},
		}},
	}
	server, err := socks5.New(config)
	if err != nil {
		l.Panic(err)
	}

	err = server.ListenAndServe("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		l.Panic(err)
	}
}


