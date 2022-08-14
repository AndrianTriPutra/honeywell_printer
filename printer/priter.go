package printer

import (
	"errors"
	"log"
	"net"
	"time"
)

var conn *net.TCPConn

func Connect(host string) error {

	var ch = make(chan bool)
	go func() {
		tcpAddr, err := net.ResolveTCPAddr("tcp", host)
		if err != nil {
			ch <- false
		}
		log.Printf("succ ResolveTCPAddr %s", host)

		conn, err = net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			ch <- false
		}
		log.Printf("succ dial %s", host)
		ch <- true
	}()

	select {
	case res := <-ch:
		if !res {
			err := errors.New("process Connect error")
			return err
		}

	case <-time.After(5 * time.Second):
		err := errors.New("process Connect timeout")
		return err
	}

	return nil
}

func Print(data string) (err error) {
	var ch = make(chan bool)
	go func() {
		resp := make([]byte, 1024)
		conn.Write([]byte(string(data)))
		_, err = conn.Read(resp)
		if err != nil {
			ch <- false
		}
		ch <- true
	}()

	select {
	case res := <-ch:
		if !res {
			err := errors.New("process print error")
			return err
		}

	case <-time.After(5 * time.Second):
		err := errors.New("process print timeout")
		return err
	}

	return nil
}
