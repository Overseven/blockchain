package transaction

import (
	"encoding/json"
	"github.com/overseven/blockchain/transaction/itransaction"
)

func FromJSON(js []byte) (*itransaction.Data, error) {
	data := itransaction.Data{}

	err := json.Unmarshal(js, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
