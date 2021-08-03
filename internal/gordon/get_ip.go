package gordon

import (
	"context"
	"errors"

	"google.golang.org/api/compute/v1"
)

// getIP gets the IP address of an instance.
// It returns the IP address and any error encountered.
func getIP(instanceProjectId string, instanceZone string, instanceName string) (string, error) {
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
	if err != nil {
		return "", err
	}
	instanceService := compute.NewInstancesService(computeService)
	instanceGetCall := instanceService.Get(instanceProjectId, instanceZone, instanceName)
	instance, err := instanceGetCall.Do()
	if err != nil {
		return "", err
	}
	if len(instance.NetworkInterfaces) == 0 {
		return "", errors.New("no network interfaces")
	}
	ip := instance.NetworkInterfaces[0].NetworkIP
	return ip, nil
}
