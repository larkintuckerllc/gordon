package gordon

import (
	"context"
	"errors"

	"google.golang.org/api/compute/v1"
)

func getIP(projectId string, zone string, instanceName string) (*string, error) {
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
	if err != nil {
		return nil, err
	}
	instanceService := compute.NewInstancesService(computeService)
	instanceGetCall := instanceService.Get(projectId, zone, instanceName)
	instance, err := instanceGetCall.Do()
	if err != nil {
		return nil, err
	}
	if len(instance.NetworkInterfaces) == 0 {
		return nil, errors.New("no network interfaces")
	}
	ip := instance.NetworkInterfaces[0].NetworkIP
	return &ip, nil
}
