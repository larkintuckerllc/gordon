package gordon

import (
	"context"
	"fmt"

	"google.golang.org/api/dns/v1"
)

// const Zone string = "my-new-zone"
// const Suffix string = ".example.private."

func deleteRecord(projectId string, instanceName string) error {
	name := fmt.Sprintf("%s%s", instanceName, Suffix)
	ctx := context.Background()
	dnsService, err := dns.NewService(ctx)
	if err != nil {
		return err
	}
	resourceRecordSetsService := dns.NewResourceRecordSetsService(dnsService)
	resourceRecordSetsDeleteCall := resourceRecordSetsService.Delete(projectId, Zone, name, "A")
	_, err = resourceRecordSetsDeleteCall.Do()
	if err != nil {
		return err
	}
	return nil
}
