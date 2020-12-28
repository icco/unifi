package metrics

import (
	"context"
	"fmt"

	"github.com/unifi-poller/unifi"
)

// GetClients returns the numbers of clients tracked by this unifi.
func GetClients(ctx context.Context, u *unifi.Unifi) (float64, error) {
	ss, err := u.GetSites()
	if err != nil {
		return 0, fmt.Errorf("get sites: %w", err)
	}

	clients, err := u.GetClients(ss)
	if err != nil {
		return 0, fmt.Errorf("get clients: %w", err)
	}

	return float64(len(clients)), nil
}
