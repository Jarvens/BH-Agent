// date: 2019-02-27
package sys

import (
	"github.com/gorilla/websocket"
	"golang.org/x/sys/unix"
	"log"
	"reflect"
	"sync"
)

type Epoll struct {
	fd       int
	connMap  map[int]*websocket.Conn
	lock     *sync.RWMutex
	inChan   chan []byte
	outChan  chan []byte
	isClosed bool
	event    map[string]interface{}
}

type Connection struct {
	wsCon *websocket.Conn
	event map[string]interface{}
}

func NewEpoll() (*Epoll, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}
	return &Epoll{fd: fd, lock: &sync.RWMutex{}, connMap: make(map[int]*websocket.Conn)}, nil
}

func (e *Epoll) Register(conn *websocket.Conn) error {
	fd := makeFd(conn)
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, fd, &unix.EpollEvent{Events: unix.POLLIN | unix.POLLHUP, Fd: int32(fd)})

	if err != nil {
		log.Printf("EpollCtl register error: %v", err)
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	e.connMap[fd] = conn
	if len(e.connMap)%100 == 0 {
		log.Printf("Total number of connections: %d", len(e.connMap))
	}
	return nil
}

func (e *Epoll) UnRegister(conn *websocket.Conn) error {
	fd := makeFd(conn)
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		log.Printf("EpollCtl un_register error: %v", err)
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	delete(e.connMap, fd)
	return nil
}

func makeFd(conn *websocket.Conn) int {
	connVal := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn").Elem()
	tcpConn := reflect.Indirect(connVal).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")
	return int(pfdVal.FieldByName("Sysfd").Int())
}

func (e *Epoll) Close() {

}
