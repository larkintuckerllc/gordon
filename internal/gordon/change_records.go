package gordon

import (
	"context"
	"fmt"

	"google.golang.org/api/dns/v1"
)

type Operation int

const (
	Addition Operation = iota
	Deletion
)

// TODO: CHANGE TO RRS
// TODO: PARAM
const Zone string = "my-new-zone"
const Suffix string = ".example.private."

func createRecords(projectId string, instanceName string, ip string) error {
	err := changeRecords(projectId, instanceName, ip, Addition)
	return err
}

func deleteRecords(projectId string, instanceName string, ip string) error {
	err := changeRecords(projectId, instanceName, ip, Deletion)
	return err
}

func changeRecords(projectId string, instanceName string, ip string, operation Operation) error {
	var change dns.Change
	name := fmt.Sprintf("%s%s", instanceName, Suffix)
	ctx := context.Background()
	dnsService, err := dns.NewService(ctx)
	if err != nil {
		return err
	}
	changesService := dns.NewChangesService((dnsService))
	resourceRecordSet := dns.ResourceRecordSet{
		Name:    name,
		Rrdatas: []string{ip},
		Ttl:     300,
		Type:    "A",
	}
	recourceRecordSets := []*dns.ResourceRecordSet{&resourceRecordSet}
	switch operation {
	case Addition:
		change = dns.Change{
			Additions: recourceRecordSets,
		}
	case Deletion:
		change = dns.Change{
			Deletions: recourceRecordSets,
		}
	}
	changesCreateCall := changesService.Create(projectId, Zone, &change)
	_, err = changesCreateCall.Do()
	if err != nil {
		return err
	}
	return nil
}
