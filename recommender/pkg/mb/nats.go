package mb

import (
	"context"
	"github.com/nats-io/nats.go"
	"fmt"
	"os"
	"log"
	"sync"
	"time"
)

type NatsConnection interface {
	Close()
	Subscribe(subject string, callback func(msg *nats.Msg))
	Listen()
}

type natsConn struct {
	conn *nats.Conn
	ctx  context.Context
	wg   sync.WaitGroup
}

func NewMessageConnection() NatsConnection {
	c := natsConn{
		conn: nil,
		wg:   sync.WaitGroup{},
	}
	var err error

	c.conn, err = nats.Connect(getURL())
	if err != nil {
		log.Panicln(err.Error())
	}

	getStatusTxt := func(m *nats.Conn) string {
		switch c.conn.Status() {
		case nats.CONNECTED:
			return "CONNECTED"
		case nats.RECONNECTING:
			return "RECONNECTING"
		case nats.CONNECTING:
			return "CONNECTING"
		case nats.CLOSED:
			return "CLOSED"
		case nats.DISCONNECTED:
			return "DISCONNECTED"
		default:
			return "UNKNOWN"
		}
	}
	start := time.Now()
	for getStatusTxt(c.conn) != "CONNECTED" {
		log.Println("Waiting for connection...")
		time.Sleep(time.Second)
		if time.Since(start) > time.Second*10 {
			log.Panicln("Could not connect to nats server")
		}
	}
	log.Println("Connected to nats server")

	return &c
}

func (c *natsConn) Close() {
	c.wg.Done()
	c.conn.Close()
}

func (c *natsConn) Subscribe(subject string, callback func(msg *nats.Msg)) {
	_, err := c.conn.QueueSubscribe(subject, "workers", callback)
	if err != nil {
		log.Panicln(err.Error())
	}
}

func (c *natsConn) Listen() {
	c.wg.Add(10)
	c.wg.Wait()
}

func getURL() string {
	return fmt.Sprintf("nats://%s:%s@%s:%s",
		os.Getenv("NATS_USER"),
		os.Getenv("NATS_PASS"),
		os.Getenv("NATS_HOST"),
		os.Getenv("NATS_PORT"),
	)
	// return nats.DefaultURL
}
