package gordon

import (
	"context"
	"fmt"

	"google.golang.org/api/dns/v1"
)

const Zone string = "my-new-zone"
const Suffix string = ".example.private."

func createRecord(projectId string, instanceName string, ip string) error {
	name := fmt.Sprintf("%s%s", instanceName, Suffix)
	ctx := context.Background()
	dnsService, err := dns.NewService(ctx)
	if err != nil {
		return err
	}
	resourceRecordSetsService := dns.NewResourceRecordSetsService(dnsService)
	resourceRecordSet := dns.ResourceRecordSet{
		Name:    name,
		Rrdatas: []string{ip},
		Ttl:     300,
		Type:    "A",
	}
	resourceRecordSetsCreateCall := resourceRecordSetsService.Create(projectId, Zone, &resourceRecordSet)
	_, err = resourceRecordSetsCreateCall.Do()
	if err != nil {
		return err
	}
	return nil
}
