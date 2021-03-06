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
	Insert Method = iota
	Stop
	Start
	Delete
)

// parse parses Cloud Pub/Sub messages.
// It returns the method, Instance Project ID, Instance zone, Instanace name, and any error encountered.
func parse(data *[]byte) (Method, string, string, string, error) {
	var p PubSubMessage
	var d Data
	var m Method
	err := json.Unmarshal(*data, &p)
	if err != nil {
		return 0, "", "", "", err
	}
	err = json.Unmarshal(p.Message.Data, &d)
	if err != nil {
		return 0, "", "", "", err
	}
	switch d.ProtoPayload.MethodName {
	case "beta.compute.instances.insert":
		m = Insert
	case "v1.compute.instances.stop":
		m = Stop
	case "v1.compute.instances.start":
		m = Start
	case "v1.compute.instances.delete":
		m = Delete
	default:
		return 0, "", "", "", errors.New("invalid MethodName")
	}
	split := strings.Split(d.ProtoPayload.ResourceName, "/")
	instanceName := split[len(split)-1]
	return m, d.Resource.Labels.ProjectId, d.Resource.Labels.Zone, instanceName, nil
}
