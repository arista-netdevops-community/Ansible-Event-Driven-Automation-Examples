package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
)

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'", i, m.Subject, string(m.Data))
	fmt.Println(m.Subject)

}

func main() {
	url := flag.String("url", "127.0.0.1:4222", "Connection for the NATS bus default is 127.0.0.1:4222")
	subject := flag.String("subject", "test", "Subject for nats.  Default is test")
	flag.Parse()
	opts := []nats.Option{nats.Name(*subject)}
	opts = setupConnOptions(opts)

	nc, err := nats.Connect(*url, opts...)
	if err != nil {
		log.Fatal(err)
	}

	i := 0

	nc.Subscribe(*subject, func(msg *nats.Msg) {
		i += 1
		printMsg(msg, i)
	})
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]", *subject)

	runtime.Goexit()
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectHandler(func(nc *nats.Conn) {
		log.Printf("Disconnected: will attempt reconnects for %.0fm", totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}
