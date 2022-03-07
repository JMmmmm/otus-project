package main

import (
	"errors"
	"io"
	"net"
	"time"
)

var ErrTelnetAbsentConnection = errors.New("error of telnet absent connection")

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &telnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

type telnetClient struct {
	conn    net.Conn
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
}

func (client *telnetClient) Connect() error {
	if client.conn != nil {
		return nil
	}
	var err error
	client.conn, err = net.DialTimeout("tcp", client.address, client.timeout)
	if err != nil {
		return err
	}
	return nil
}

func (client *telnetClient) Send() error {
	if client.conn == nil {
		return ErrTelnetAbsentConnection
	}
	request := make([]byte, 1024)
	n, err := client.in.Read(request)
	if err != nil {
		return err
	}
	_, err = client.conn.Write(request[:n])
	return err
}

func (client *telnetClient) Receive() error {
	if client.conn == nil {
		return ErrTelnetAbsentConnection
	}
	response := make([]byte, 1024)
	n, err := client.conn.Read(response)
	if err != nil {
		return err
	}
	_, err = client.out.Write(response[:n])
	return err
}

func (client *telnetClient) Close() error {
	if client.conn == nil {
		return ErrTelnetAbsentConnection
	}
	return client.conn.Close()
}
