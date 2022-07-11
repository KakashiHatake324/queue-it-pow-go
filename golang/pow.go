package golang

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"strconv"
	"strings"
)

// Main function
func main() {
	PostFix, err := SolvePoW("af085a5f-ae56-4450-8bf8-11cabf2b140a", 3, 25)
	if err != nil {
		log.Fatal(err)
	}

	Solution := QueueItSolution{
		Hash: PostFix,
		Type: "t.Footsites.QueueIT.PowChallenge.Parameters.Type",
	}

	log.Println(Solution)
}

// Use the information from https://footlocker.queue-it.net/challengeapi/pow/challenge/{userID} to solve POW
func SolvePoW(inputString string, complexity, runs int) ([]QueueItPowPostFix, error) {
	Solutions := []QueueItPowPostFix{}
	CurrentRuns := 0
	loop := true
	for postfix := 0; loop; postfix++ {
		shaObj := sha256.New()
		shaObj.Write([]byte(inputString))
		shaObj.Write([]byte(strconv.Itoa(postfix)))
		hash := hex.EncodeToString(shaObj.Sum(nil))

		if strings.HasPrefix(string(hash), strings.Repeat("0", complexity)) {
			Solutions = append(Solutions, QueueItPowPostFix{
				Postfix: postfix,
				Hash:    hash,
			})
			CurrentRuns++
		}
		if CurrentRuns == runs {
			return Solutions, nil
		}
	}
	return []QueueItPowPostFix{}, errors.New("error solving pow")
}

// Postfix struct
type QueueItPowPostFix struct {
	Postfix int    `json:"postfix"`
	Hash    string `json:"hash"`
}

// Solution to send in verify post
type QueueItSolution struct {
	Hash []QueueItPowPostFix `json:"hash"`
	Type string              `json:"type"`
}
