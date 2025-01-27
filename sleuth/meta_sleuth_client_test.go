package sleuth

import "testing"

func TestMetaSleuthClient(t *testing.T) {
	c := NewDefaultClient()
	chains, err := c.GetSupportedChains()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", chains)
}

func TestMetaSleuthClientGetAddressLabels(t *testing.T) {
	c := NewDefaultClient()
	labels, err := c.GetAddressLabels(ChainIDSolana, "CS94dE2znqTYxcFtfdfpsEjkphBdsVZeVKhLoykMX7Ge")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", labels)
}

func TestMetaSleuthClientGetBatchLabels(t *testing.T) {
	c := NewDefaultClient()
	labels, err := c.GetBatchLabels(-3, "9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM", "5tzFkiKscXHK5ZXCGbXZxdw7gTjjD1mBwuoFbhUvuAi9")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", labels)
}

func TestMetaSleuthClientGetEntity(t *testing.T) {
	c := NewDefaultClient()
	entity, err := c.GetEntity("Binance")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", entity)
}
