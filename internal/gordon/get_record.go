package gordon

import (
	"context"
	"fmt"

	"google.golang.org/api/dns/v1"
)

// getRecord gets an A record.
// It returns any error encountered
func getRecord(instanceName string) error {
	name := fmt.Sprintf("%s.%s", instanceName, dName)
	ctx := context.Background()
	dnsService, err := dns.NewService(ctx)
	if err != nil {
		return err
	}
	resourceRecordSetsService := dns.NewResourceRecordSetsService(dnsService)
	resourceRecordSetsGetCall := resourceRecordSetsService.Get(pId, zName, name, "A")
	_, err = resourceRecordSetsGetCall.Do()
	if err != nil {
		return err
	}
	return nil
}
