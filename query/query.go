package query

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/aseure/jagger/player"
)

type Query struct {
	Command string `json:"command"`
	Params  string `json:"params"`
}

func NewQuery(reader io.Reader) *Query {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("Cannot read from reader")
		return nil
	}

	var query Query
	if err := json.Unmarshal(data, &query); err != nil {
		fmt.Println("Cannot unmarshal data from the reader: %s", err.Error())
		return nil
	}

	return &query
}

func (q *Query) Execute() {
	switch q.Command {
	case "play":
		fmt.Println("[CMD] Play", q.Params)
		player.Play(q.Params)
	case "stop":
		fmt.Println("[CMD] Stop")
		player.Stop()
	}
}
