package addrs_test

import (
	"slices"
	"testing"

	"github.com/dwdwow/goalb/addrs"
	"github.com/dwdwow/goalb/labels"
)

func TestAddrsFilterWithAllLabels(t *testing.T) {
	tests := []struct {
		name     string
		addrsMap map[string]labels.AddressLabels
		labels   []labels.Label
		want     []string
	}{
		{
			name: "filter CEX Binance hot wallets",
			addrsMap: map[string]labels.AddressLabels{
				"2ojv9BAiHUrvsm9gxDe7fJSzbNZSJcxZvf8dqmWGHG8S": {
					Address: "2ojv9BAiHUrvsm9gxDe7fJSzbNZSJcxZvf8dqmWGHG8S",
					Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_ON_CURVE, labels.CEX, labels.CEX_BINANCE, labels.HOT_WALLET},
				},
				"4f77K3QgVREaoAZ7U6EG1BsQMdPKRjzPbNznzzKT2gEJ": {
					Address: "4f77K3QgVREaoAZ7U6EG1BsQMdPKRjzPbNznzzKT2gEJ",
					Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.CEX, labels.CEX_BINANCE, labels.SOL_STAKE},
				},
			},
			labels: []labels.Label{labels.CEX_BINANCE, labels.HOT_WALLET},
			want:   []string{"2ojv9BAiHUrvsm9gxDe7fJSzbNZSJcxZvf8dqmWGHG8S"},
		},
		{
			name: "filter DEX programs",
			addrsMap: map[string]labels.AddressLabels{
				"srmqPvymJeFKQ4zGQed1GFppgkRHL9kaELCbyksJtPX": {
					Address: "srmqPvymJeFKQ4zGQed1GFppgkRHL9kaELCbyksJtPX",
					Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_OPENBOOK},
				},
				"opnb2LAfJYbRMAHHvqjCwQxanZn7ReEHp1k81EohpZb": {
					Address: "opnb2LAfJYbRMAHHvqjCwQxanZn7ReEHp1k81EohpZb",
					Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_PROGRAM, labels.SOL_OFF_CURVE, labels.DEX, labels.DEX_OPENBOOK},
				},
				"GKnHiWh3RRrE1zsNzWxRkomymHc374TvJPSTv2wPeYdB": {
					Address: "GKnHiWh3RRrE1zsNzWxRkomymHc374TvJPSTv2wPeYdB",
					Labels:  []labels.Label{labels.CHAIN_SOLANA, labels.SOL_OFF_CURVE, labels.DEFI_KAMINO},
				},
			},
			labels: []labels.Label{labels.SOL_PROGRAM, labels.DEX},
			want:   []string{"srmqPvymJeFKQ4zGQed1GFppgkRHL9kaELCbyksJtPX", "opnb2LAfJYbRMAHHvqjCwQxanZn7ReEHp1k81EohpZb"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addrs.AddrsFilterWithAllLabels(tt.addrsMap, tt.labels)
			if len(got) != len(tt.want) {
				t.Errorf("AddrsFilterWithAllLabels() got %v, want %v", got, tt.want)
				return
			}
			for _, addr := range tt.want {
				if !slices.Contains(got, addr) {
					t.Errorf("AddrsFilterWithAllLabels() got %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
