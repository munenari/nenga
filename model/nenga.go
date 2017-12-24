package model

// Nenga seems many nenga post cards
type Nenga struct {
	Destinations []Destination
	Sender       Sender
}

type atenaAbstract struct {
	Names    []string
	Postcode string
	Address  string
}
