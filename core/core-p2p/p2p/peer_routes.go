package p2p

import (
	"bytes"
	"encoding/json"
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-registration/registration"
	"github.com/jneubaum/honestvote/tests/logger"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	//decode json data
	d := json.NewDecoder(conn)

	for {

		var write Message
		err := d.Decode(&write)

		if err != nil {
			logger.Println("peer_routes.go", "HandleConn()", err.Error())
			return
		}

		switch write.Message {
		case "connect":
			logger.Println("peer_routes.go", "HandleConn()", "Recieved Connect Message")

			var node database.Node
			json.Unmarshal(write.Data, &node)

			AcceptConnectMessage(node, conn)
		case "get id":
			var node database.Node
			json.Unmarshal(write.Data, &node)
			node.IPAddress = conn.RemoteAddr().String()[0:9]
			if !database.DoesNodeExist(node) {
				database.AddNode(node)
			}
		case "recieve data":
			buffer := bytes.NewBuffer(write.Data)
			DecodeData(buffer)
		case "get data":
			database.MoveDocuments(Nodes, database.DatabaseName, database.CollectionPrefix+database.ElectionHistory)
		case "vote":
			ReceiveVote(write.Vote)
		case "register":
			tcp_port := strconv.Itoa(TCP_PORT)
			registration.EmailRegistration("jacob@neubaum.com (senders_email)", "election_name", "senders_public_key", PublicIP, tcp_port)
		case "become peer":
			//vote on
		case "new election":
			//Create a new election
		case "verify":
			block := new(database.Block)
			json.Unmarshal(write.Data, block)
			logger.Println("peer_routes.go", "HandleConn()", "Verifying")
			VerifyBlock(*block, conn)
		case "sign":
			answer, err := strconv.ParseBool(string(write.Data))

			if err == nil {
				ReceiveResponses(answer, write.Signature)
			}
		case "update":
			block := new(database.Block)
			json.Unmarshal(write.Data, block)
			if database.UpdateBlockchain(database.MongoDB, *block) {
				PrevHash = block.Hash
				PrevIndex = block.Index
				logger.Println("peer_routes.go", "HandleConn()", string(PrevIndex)+" "+PrevHash)
			}
		}
	}
}
