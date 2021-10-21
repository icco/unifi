package metrics

import (
	"context"
	"fmt"

  "github.com/unpoller/unifi"
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

// GetBytesPerSecond returns the current rate of ppm from the WAN.
func GetBytesPerSecond(ctx context.Context, u *unifi.Unifi) (*NetworkNumber, error) {
	sites, err := u.GetSites()
	if err != nil {
		return nil, fmt.Errorf("get sites: %w", err)
	}

	data := &NetworkNumber{}
	for _, s := range sites {
		for _, h := range s.Health {
			if h.Subsystem == "wan" {
				data.Upload += h.TxBytesR.Val
				data.Download += h.RxBytesR.Val
			}
		}
	}

	return data, nil
}
