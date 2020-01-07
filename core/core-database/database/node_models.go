package database

import (
	"encoding/json"
	"strconv"
)

var CollectionPrefix string = ""        // Multiple nodes can work on the same host using different collection prefixes
var DatabaseName string = "honestvote"  // Database is the same for all nodes even for a test net
var ElectionHistory string = "election" // Elections
var Connections string = "connections"  // Nodes on network

type PublicKey string

// Email registrants
var EmailRegistrants string = "email_registrants"

type Block struct {
	Index       int               `json:"index"`
	Timestamp   string            `json:"timestamp"`
	Transaction Transaction       `json:"transaction"` // instead of interface, should be transaction
	Hash        string            `json:"hash"`
	PrevHash    string            `json:"prevhash"`
	Signatures  map[string]string `json:"signatures"`
}

type Transaction interface {
	VerifySignature() bool
}

type Vote struct {
	Vote      int       `json:"vote"`
	Election  string    `json:"election"`
	Receiver  []string  `json:"receiver"`
	Sender    PublicKey `json:"sender"`
	Signature string    `json:"signature"`
}

func (vote Vote) VerifySignature() bool {
	vote_weight := strconv.Itoa(vote.Vote)
	plaintext := vote_weight + vote.Election
	for i, _ := range vote.Receiver {
		plaintext += vote.Receiver[i]
	}

	// After decrypt method available: plaintext == crypto.Decrypt(vote.Signature, Sender)
	if plaintext == vote.Signature {
		return true
	}
	return false
}

type Election struct {
	Name           string     `json:"name"`
	Start          string     `json:"start"`
	End            string     `json:"end"`
	EligibleVoters int        `json:"registeredVoters"`
	Positions      []Position `json:"positions"`
	Sender         PublicKey  `json:"sender"`
	Signature      string     `json:"signature"`
}

func (election Election) VerifySignature() bool {
	eligible_voters := strconv.Itoa(election.EligibleVoters)
	plaintext := election.Name + election.Start + election.End + eligible_voters
	for i, _ := range election.Positions {
		json_bytes, _ := json.Marshal(election.Positions[i])
		plaintext += string(json_bytes)
	}

	// After decrypt method available: plaintext == crypto.Decrypt(vote.Signature, Sender)
	if plaintext == election.Signature {
		return true
	}

	return false
}

type Node struct {
	Institution string
	IPAddress   string
	Port        int
	Role        string // peer | full | registry
	Identity    PublicKey
	Signature   string
}

func (node Node) VerifySignature() bool {
	if true {
		return true
	}
	return false
}

type Position struct {
	Name       string      `json:"name"`
	ID         int         `json:"id"`
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Name      string `json:"name"`
	PublicKey string `json:"key"`
	Election  string `json:"election"`
	Votes     int    `json:"votes"`
}

type AwaitingRegistration struct {
	Email     string `json:"email"`
	Election  string `json:"election"`
	Code      string
	PublicKey string `json:"publicKey"`
	Timestamp string
}
