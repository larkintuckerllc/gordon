package gordon

import (
	"context"
	"fmt"

	"google.golang.org/api/dns/v1"
)

// TODO: DOCUMENT FUNCTIONS
func getRecord(projectId string, instanceName string) error {
	name := fmt.Sprintf("%s.%s", instanceName, dnsName)
	ctx := context.Background()
	dnsService, err := dns.NewService(ctx)
	if err != nil {
		return err
	}
	resourceRecordSetsService := dns.NewResourceRecordSetsService(dnsService)
	resourceRecordSetsGetCall := resourceRecordSetsService.Get(projectId, zoneName, name, "A")
	_, err = resourceRecordSetsGetCall.Do()
	if err != nil {
		return err
	}
	return nil
}
