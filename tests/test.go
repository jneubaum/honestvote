package main

import (
	"encoding/hex"
	"fmt"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
)

var plaintext string = "this should work, please work.  I wonder why this does not work.  Kernel Sanders is king of all chicken and I wonder how chicken would taste if it were not for the great KFC Lord"

func main() {

	// private_key := "973920ad0b8e597663c791e0332373bd1d04b86bfc8b927419af7e1e044b393a"

	// var vote database.Vote = database.Vote{
	// 	Type:     "Vote",
	// 	Election: "3046022100bbb4e8ed9694d7ea6ebb40fb48b2b3cf8f861a6979ad36c1bf6d40a71585068f022100bb23a10bf55f0bf866baef193a016839d1d23add1a19b7db651c429a02e0baf5",
	// 	Receiver: []database.SelectedCandidate{
	// 		database.SelectedCandidate{
	// 			PositionId: "demfrmeororev",
	// 			Recipient:  "Beverlys Birthdays",
	// 		}, database.SelectedCandidate{
	// 			PositionId: "defmrfmrkmef",
	// 			Recipient:  "Art of Recycle",
	// 		},
	// 	},
	// 	Sender: "02341b3431dd869c584a8ceb0e9e4da5d59e92e444e08bf0d58a3b14296b459b80",
	// }

	// encoded, err := vote.Encode()
	// if err != nil {
	// 	return
	// }
	hash := crypto.CalculateHash([]byte("test"))
	s1 := hex.EncodeToString(hash)
	fmt.Println(s1)
	hash2, _ := hex.DecodeString(s1)
	fmt.Println(hash2)
	// vote.Signature, _ = crypto.Sign([]byte(hash), private_key)

	// crypto.Verify([]byte(hash), "02341b3431dd869c584a8ceb0e9e4da5d59e92e444e08bf0d58a3b14296b459b80", vote.Signature)

}
