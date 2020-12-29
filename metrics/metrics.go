package metrics

import (
	"context"
	"fmt"

	"github.com/unifi-poller/unifi"
)

// GetClients returns the numbers of clients tracked by this unifi.
func GetClients(ctx context.Context, u *unifi.Unifi) (float64, error) {
	sites, err := u.GetSites()
	if err != nil {
		return 0, fmt.Errorf("get sites: %w", err)
	}

	clients, err := u.GetClients(sites)
	if err != nil {
		return 0, fmt.Errorf("get clients: %w", err)
	}

	return float64(len(clients)), nil
}

type NetworkNumber struct {
	Upload   float64
	Download float64
}

// GetPacketsPerMinute returns the current rate of ppm from the WAN.
func GetPacketsPerMinute(ctx context.Context, u *unifi.Unifi) (*NetworkNumber, error) {
	sites, err := u.GetSites()
	if err != nil {
		return nil, fmt.Errorf("get sites: %w", err)
	}

	devs, err := u.GetDevices(sites)
	if err != nil {
		return nil, fmt.Errorf("get devices: %w", err)
	}

	data := &NetworkNumber{}
	for _, d := range devs.UDMs {
		// TODO: Maybe get this from WAN data instead?
		data.Upload += d.TxBytes.Val
		data.Download += d.RxBytes.Val
	}

	return data, nil
}
