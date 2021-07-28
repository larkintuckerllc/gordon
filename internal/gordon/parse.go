package gordon

import (
	"encoding/json"
	"errors"
)

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data"`
	} `json:"message"`
}

type Data struct {
	ProtoPayload struct {
		MethodName string `json:"methodName"`
	} `json:"protoPayload"`
	Resource struct {
		Labels struct {
			InstanceId string `json:"instance_id"`
			Zone       string `json:"zone"`
		} `json:"labels"`
	} `json:"resource"`
}

type Method int

const (
	Allocate Method = iota
	Deallocate
)

type Instance struct {
	InstanceId string
	Zone       string
}

func parse(data *[]byte) (*Method, *Instance, error) {
	var p PubSubMessage
	var d Data
	var m Method
	err := json.Unmarshal(*data, &p)
	if err != nil {
		return nil, nil, err
	}
	err = json.Unmarshal(p.Message.Data, &d)
	if err != nil {
		return nil, nil, err
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
		return nil, nil, errors.New("invalid MethodName")
	}
	i := Instance{d.Resource.Labels.InstanceId, d.Resource.Labels.Zone}
	return &m, &i, nil
}
