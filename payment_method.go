package globe

type PaymentMethod struct {
	// Attributes
	Code string `xml:"code,attr"` // { A | B | K | R | C | W | I | O | F | P | V | Y | X | E | H | S }

	Description string `xml:"Description,omitempty"`
}
