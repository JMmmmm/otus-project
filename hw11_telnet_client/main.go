package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "Disk query process time out")
	flag.Parse()

	inputArguments := flag.Args()
	address := net.JoinHostPort(inputArguments[0], inputArguments[1])

	ctxNotify, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	ctx, cancel := context.WithTimeout(ctxNotify, timeout)

	telnetClient := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	err := telnetClient.Connect()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("...Connected to %s", address)
	defer func() {
		err = telnetClient.Close()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("...Connection was closed")
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		makeReceiving(ctx, telnetClient, wg)
		cancel()
	}()

	wg.Add(1)
	go func() {
		makeSending(ctx, telnetClient, wg)
		cancel()
	}()

	wg.Wait()
}

func makeReceiving(ctx context.Context, telnetClient TelnetClient, wg *sync.WaitGroup) {
	defer wg.Done()
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			err := telnetClient.Receive()
			if err != nil {
				log.Println(err)
				break OUTER
			}
		}
	}
}

func makeSending(ctx context.Context, telnetClient TelnetClient, wg *sync.WaitGroup) {
	defer wg.Done()
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			err := telnetClient.Send()
			if err != nil {
				log.Println(err)
				break OUTER
			}
		}
	}
}
