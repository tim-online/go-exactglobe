package globe

import errors "github.com/tim-online/go-errors"

type PaymentMethod struct {
	// Attributes
	Code PaymentMethodCode `xml:"code,attr"` // { A | B | K | R | C | W | I | O | F | P | V | Y | X | E | H | S }

	Description string `xml:"Description,omitempty"`
}

// K=Cash, B=Bank, G=Giro, I=Collection, A=Giro collection slip, S=Bacs, W=Bill of exchange, D=Bill of exchange, M=Domiciliation, V=ESR payments, Y=ES payments, X=Payments in CHF and FC, Z=Bank cheques
type PaymentMethodCode string

func (p PaymentMethodCode) Validate() error {
	valid := []PaymentMethodCode{"K", "B", "G", "I", "A", "C", "S", "W", "D", "M", "V", "Y", "X", "Z"}

	if p == "" {
		return nil
	}

	for _, v := range valid {
		if p == v {
			return nil
		}
	}

	return errors.Errorf("Invalid PaymentMethodCode '%s'", p)
}
