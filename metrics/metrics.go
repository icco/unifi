package metrics

import (
	"context"
	"fmt"

	"github.com/dim13/unifi"
)

// GetClients returns the numbers of clients tracked by this unifi.
func GetClients(ctx context.Context, u *unifi.Unifi) (float64, error) {
	var total float64
	ss, err := u.Sites()
	if err != nil {
		return 0, fmt.Errorf("get sites: %w", err)
	}

	for _, s := range ss {
		clients, err := u.Sta(&s)
		if err != nil {
			return 0, fmt.Errorf("get clients: %w", err)
		}

		total += float64(len(clients))
	}

	return total, nil
}
