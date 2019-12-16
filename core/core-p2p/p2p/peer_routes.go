package p2p

import (
	"bytes"
	"encoding/json"
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	var buf [256]byte

	for {
		length, err := conn.Read(buf[0:])

		if err != nil {
			return
		}

		if string(buf[0:7]) == "connect" {
			port, err := strconv.Atoi(string(buf[8:length]))

			if err == nil {
				nodes[port] = true
				tmpPeer := database.Peer{
					IPAddress: "127.0.0.1",
					Port:      port,
					Socket:    conn,
				}
				Peers = append(Peers, tmpPeer)
			}
		} else if string(buf[0:12]) == "recieve data" {
			buffer := bytes.NewBuffer(buf[13:length])
			tmpArray := new([]database.Candidate)
			js := json.NewDecoder(buffer)
			err := js.Decode(tmpArray)
			if err == nil {
				database.UpdateMongo(database.MongoDB, *tmpArray)
			}
		} else if string(buf[0:8]) == "get data" {
			database.MoveDocuments(Peers)
		}
	}
}