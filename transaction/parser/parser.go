package parser

import (
	"encoding/json"

	tr "github.com/Overseven/blockchain/transaction"
)

func FromJSON(js []byte) (*tr.Data, error) {
	data := tr.Data{}

	err := json.Unmarshal(js, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
