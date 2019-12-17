package p2p

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	var buf [256]byte

	for {
		length, err := conn.Read(buf[0:])
		fmt.Println(string(buf[0:length]))

		if err != nil {
			return
		}

		if string(buf[0:7]) == "connect" {
			port, err := strconv.Atoi(string(buf[8:length]))

			if err == nil {
				Nodes[port] = true
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
		} else if string(buf[0:4]) == "vote" {
			//TODO: Input a vote and send it to peer to verify
			vote, err := strconv.Atoi(string(buf[5:length]))
			if err == nil {
				block := consensus.GenerateBlock(database.Block{}, database.Transaction{
					Sender:   "",
					Vote:     vote,
					Receiver: "",
				})
				fmt.Println("Block created!")
				fmt.Println(block)
			}
		}
	}
}
