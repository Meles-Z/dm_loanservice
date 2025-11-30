package grpc

import "time"

func parseDateFlexible(s string) (time.Time, error) {
	// Try RFC3339 first
	t, err := time.Parse(time.RFC3339, s)
	if err == nil {
		return t, nil
	}

	// Try Go's default string layout produced by %v / Time.String()
	layout := "2006-01-02 15:04:05 -0700 MST"
	return time.Parse(layout, s)
}
