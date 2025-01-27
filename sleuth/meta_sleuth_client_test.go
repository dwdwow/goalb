package sleuth

import (
	"testing"

	"github.com/dwdwow/goalb/addrs"
)

func TestMetaSleuthClient(t *testing.T) {
	c := DefaultClient()
	chains, err := c.GetSupportedChains()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", chains)
}

func TestMetaSleuthClientGetAddressLabels(t *testing.T) {
	c := DefaultClient()
	labels, err := c.GetAddressLabels(ChainIDSolana, "CS94dE2znqTYxcFtfdfpsEjkphBdsVZeVKhLoykMX7Ge")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", labels)
}

func TestMetaSleuthClientGetBatchLabels(t *testing.T) {
	m := addrs.SolBinance
	var addrs []string
	for _, v := range m {
		addrs = append(addrs, v)
	}
	c := DefaultClient()
	labels, err := c.GetBatchLabels(-3, addrs...)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range labels {
		t.Logf("%+v\n", v)
	}
}

func TestMetaSleuthClientGetEntity(t *testing.T) {
	c := DefaultClient()
	entity, err := c.GetEntity("Wintermute")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", entity)
}
