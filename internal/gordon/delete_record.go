package gordon

import (
	"context"
	"fmt"

	"google.golang.org/api/dns/v1"
)

// deleteRecord deletes an A record.
// It returns any error encountered.
func deleteRecord(projectId string, instanceName string) error {
	name := fmt.Sprintf("%s.%s", instanceName, dName)
	ctx := context.Background()
	dnsService, err := dns.NewService(ctx)
	if err != nil {
		return err
	}
	resourceRecordSetsService := dns.NewResourceRecordSetsService(dnsService)
	resourceRecordSetsDeleteCall := resourceRecordSetsService.Delete(projectId, zName, name, "A")
	_, err = resourceRecordSetsDeleteCall.Do()
	if err != nil {
		return err
	}
	return nil
}
