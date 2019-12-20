package p2p

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte

	for {
		length, err := conn.Read(buf[0:])

		if err != nil {
			return
		}

		if string(buf[0:7]) == "connect" {

			//ADD TO DATABASE AS WELL
			port, err := strconv.Atoi(string(buf[8:length]))

			log.Println("Recieved Connect Message")
			if err == nil {
				// Nodes[port] = true
				tmpNode := database.TempNode{
					IPAddress: "127.0.0.1",
					Port:      port,
					Socket:    conn,
				}
				Nodes = append(Nodes, tmpNode)
				fmt.Println(Nodes)
				// permNode := database.Node{}
				// database.AddToTable(permNode.IPAddress, permNode.Port)
			}
		} else if string(buf[0:12]) == "recieve data" {
			buffer := bytes.NewBuffer(buf[13:length])
			tmpArray := new([]database.Candidate)
			js := json.NewDecoder(buffer)
			err := js.Decode(tmpArray)
			if err == nil {
				database.UpdateMongo(database.MongoDB, *tmpArray, database.DatabaseName, database.CollectionPrefix+database.ElectionHistory)
			}
		} else if string(buf[0:8]) == "get data" {
			database.MoveDocuments(Nodes, database.DatabaseName, database.CollectionPrefix+database.ElectionHistory)
		} else if string(buf[0:4]) == "vote" { //Get a vote and make a block out of it
			sVote := string(buf[5:length])
			sVote = strings.TrimSuffix(sVote, "\n")
			vote, err := strconv.Atoi(sVote)
			if err == nil {
				block := consensus.GenerateBlock(database.Block{}, database.Transaction{
					Sender:   "",
					Vote:     vote,
					Receiver: "",
				}, Port)

				//Check if there is a proposed block currently, if so, add to the queue
				if ProposedBlock == (database.Block{}) {
					fmt.Println("Empty, proposing this block.")
					ProposedBlock = block
					ProposeBlock(ProposedBlock, Nodes)
				} else {
					fmt.Println("Not Empty, sending to queue.")
					BlockQueue = append(BlockQueue, block)
					fmt.Println(BlockQueue)
				}
			}
		} else if string(buf[0:6]) == "verify" { //Verifying that the sent block is correct(sign/reject)
			block := new(database.Block)
			json.Unmarshal(buf[7:length], block)
			VerifyBlock(*block)
		} else if string(buf[0:4]) == "sign" { //Response from all Nodes verifying block
			fmt.Println("Block Signed!")
		}
	}
}
