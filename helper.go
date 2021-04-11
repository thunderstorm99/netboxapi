package netboxapi

import (
	"encoding/json"
)

func resultToOtherStructure(input APIAnswer, output interface{}) error {
	// get JSON representation of results
	c, err := json.Marshal(input.Results)
	if err != nil {
		return err
	}

	// Unmarshal onto specific structure via output
	if err := json.Unmarshal(c, &output); err != nil {
		return err
	}

	return nil
}
