package globe

type Skills []Skill

type Skill struct {
	Description string `xml:"Description"`
	Level       int    `xml:"Level,omitempty"`
}
