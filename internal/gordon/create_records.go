package gordon

import (
	"context"
	"fmt"

	"google.golang.org/api/dns/v1"
)

// TODO: PARAM
const Zone string = "my-new-zone"
const Suffix string = ".example.private."

func createRecords(projectId string, instanceName string, ip string) error {
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
