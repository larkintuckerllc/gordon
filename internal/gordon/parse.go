package gordon

import "encoding/json"

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

func parse(data []byte) (*PubSubMessage, error) {
	var m PubSubMessage
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
