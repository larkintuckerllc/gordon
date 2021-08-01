package gordon

import (
	"context"

	"google.golang.org/api/dns/v1"
)

// TODO: PARAM
const Zone string = "my-new-zone"

func createRecords(projectId string, instanceName string, ip string) error {
	ctx := context.Background()
	dnsService, err := dns.NewService(ctx)
	if err != nil {
		return err
	}
	changesService := dns.NewChangesService((dnsService))
	resourceRecordSet := dns.ResourceRecordSet{
		Kind:    "A",
		Name:    instanceName,
		Ttl:     300,
		Rrdatas: []string{ip},
	}
	recourceRecordSets := []*dns.ResourceRecordSet{&resourceRecordSet}
	change := dns.Change{
		Additions: recourceRecordSets,
	}
	changesCreateCall := changesService.Create(projectId, Zone, &change)
	_, err = changesCreateCall.Do()
	if err != nil {
		return err
	}
	return nil
}
