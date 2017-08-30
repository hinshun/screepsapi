package screepsapi

import (
	"encoding/json"
	"fmt"

	"github.com/hinshun/screepsapi/screepstype"
)

type RawMemoryResponse struct {
	Response
	Data string `json:"data"`
}

// UnmarshalJSON unmarshals the data string, which is of the form `gz:`
// followed by base64-encoded gzipped JSON encoding of the requested
// memory path, into a more functional map[string]interface{}.
func (m *MemoryResponse) UnmarshalJSON(b []byte) error {
	rawMemoryResp := RawMemoryResponse{}
	err := json.Unmarshal(b, &rawMemoryResp)
	if err != nil {
		return fmt.Errorf("failed to unmarshal raw response: %s", err)
	}

	m.Ok = rawMemoryResp.Ok
	if rawMemoryResp.Data != "" {
		unzippedData, err := screepstype.Unzip(rawMemoryResp.Data, screepstype.CompressionTypeGzip)
		if err != nil {
			return fmt.Errorf("failed to unzip gzipped data: %s", err)
		}

		err = json.Unmarshal(unzippedData, &m.Data)
		if err != nil {
			return fmt.Errorf("failed to unmarshal unzipped data: %s", unzippedData)
		}
	}

	return nil
}
