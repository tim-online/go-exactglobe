package globe

type Marital struct {
	// Attributes
	Type string `xml:"type,attr"` // { D | L | M | N | R | S | W }

	Date Date `xml:"Date"`
}
