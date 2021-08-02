package gordon

import (
	"context"
	"fmt"

	"google.golang.org/api/dns/v1"
)

// TODO: PARAM
// const Zone string = "my-new-zone"
// const Suffix string = ".example.private."

func getRecord(projectId string, instanceName string) (*string, error) {
	name := fmt.Sprintf("%s%s", instanceName, Suffix)
	ctx := context.Background()
	dnsService, err := dns.NewService(ctx)
	if err != nil {
		return nil, err
	}
	resourceRecordSetsService := dns.NewResourceRecordSetsService(dnsService)
	resourceRecordSetsGetCall := resourceRecordSetsService.Get(projectId, Zone, name, "A")
	resourceRecordSet, err := resourceRecordSetsGetCall.Do()
	if err != nil {
		return nil, err
	}
	ip := resourceRecordSet.Rrdatas[0]
	return &ip, nil
}
