package gordon

import (
	"encoding/json"
	"errors"
	"strings"
)

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data"`
	} `json:"message"`
}

type Data struct {
	ProtoPayload struct {
		MethodName   string `json:"methodName"`
		ResourceName string `json:"resourceName"`
	} `json:"protoPayload"`
	Resource struct {
		Labels struct {
			ProjectId string `json:"project_id"`
			Zone      string `json:"zone"`
		} `json:"labels"`
	} `json:"resource"`
}

type Method int

const (
	Allocate Method = iota
	Deallocate
)

func parse(data *[]byte) (*Method, *string, *string, *string, error) {
	var p PubSubMessage
	var d Data
	var m Method
	err := json.Unmarshal(*data, &p)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	err = json.Unmarshal(p.Message.Data, &d)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	switch d.ProtoPayload.MethodName {
	case "beta.compute.instances.insert":
		m = Allocate
	case "v1.compute.instances.stop":
		m = Deallocate
	case "v1.compute.instances.start":
		m = Allocate
	case "v1.compute.instances.delete":
		m = Deallocate
	default:
		return nil, nil, nil, nil, errors.New("invalid MethodName")
	}
	split := strings.Split(d.ProtoPayload.ResourceName, "/")
	instanceName := split[len(split)-1]
	return &m, &d.Resource.Labels.ProjectId, &d.Resource.Labels.Zone, &instanceName, nil
}
