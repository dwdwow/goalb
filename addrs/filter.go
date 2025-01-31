package addrs

import (
	"slices"

	"github.com/dwdwow/goalb/labels"
)

func AddrsFilterWithAllLabels(addrsMap map[string]labels.AddressLabels, labels []labels.Label) []string {
	filtered := []string{}
	for addr, addrLabels := range addrsMap {
		hasAllLabels := true
		for _, label := range labels {
			if !slices.Contains(addrLabels.Labels, label) {
				hasAllLabels = false
				break
			}
		}
		if hasAllLabels {
			filtered = append(filtered, addr)
		}
	}
	return filtered
}
