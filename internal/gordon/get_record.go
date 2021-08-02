package gordon

import (
	"context"
	"fmt"

	"google.golang.org/api/dns/v1"
)

// const Zone string = "my-new-zone"
// const Suffix string = ".example.private."
// TODO: DOCUMENT FUNCTIONS
func getRecord(projectId string, instanceName string) error {
	name := fmt.Sprintf("%s%s", instanceName, Suffix)
	ctx := context.Background()
	dnsService, err := dns.NewService(ctx)
	if err != nil {
		return err
	}
	resourceRecordSetsService := dns.NewResourceRecordSetsService(dnsService)
	resourceRecordSetsGetCall := resourceRecordSetsService.Get(projectId, Zone, name, "A")
	_, err = resourceRecordSetsGetCall.Do()
	if err != nil {
		return err
	}
	return nil
}
