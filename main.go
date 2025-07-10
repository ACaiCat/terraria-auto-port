package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strings"
)

var config Config

func main() {

	configPoint := ReadConfig()
	if configPoint == nil {
		return
	}
	config = *configPoint

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.ListenPort))

	if err != nil {
		log.Println("Failed listen port: ", err)
		return
	}

	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		log.Println("Accept connection from " + conn.RemoteAddr().String())
		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	readLength, _ := conn.Read(buffer)
	reader := bytes.NewReader(buffer[2:readLength])

	var packageNum byte
	binary.Read(reader, binary.BigEndian, &packageNum)

	var versionLength uint8
	binary.Read(reader, binary.BigEndian, &versionLength)

	versionBytes := make([]byte, versionLength)
	binary.Read(reader, binary.BigEndian, &versionBytes)

	version := string(versionBytes)
	var p2sConn net.Conn
	var err error

	switch {
	case strings.HasPrefix(version, "Terraria"):
		p2sConn, err = net.Dial("tcp", config.VanillaAddress)
	case strings.HasPrefix(version, "tModLoader"):
		p2sConn, err = net.Dial("tcp", config.TModLoaderAddress)
	default:
		log.Println("Unknown client: " + conn.RemoteAddr().String())
		conn.Close()
		return

	}

	if err != nil {
		log.Println("Cannot connect server: ", err)
		return
	}

	p2sConn.Write(buffer[:readLength])
	go buildSocketBridge(p2sConn, conn)
	go buildSocketBridge(conn, p2sConn)

}

func buildSocketBridge(readConn net.Conn, sendConn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		readLength, err := readConn.Read(buffer)

		if err != nil {
			log.Println("Connection lost: ", err)
			readConn.Close()
			sendConn.Close()
			break
		}

		if _, err = sendConn.Write(buffer[:readLength]); err != nil {
			log.Println("Connection lost: ", err)
			readConn.Close()
			sendConn.Close()
			break
		}
	}

}
