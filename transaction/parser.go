package transaction

import (
	"encoding/json"

	"github.com/overseven/blockchain/interfaces"
)

func FromJSON(js []byte) (*interfaces.Data, error) {
	data := interfaces.Data{}

	err := json.Unmarshal(js, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
