package globe

type SalesEmployees []SalesEmployee

type SalesEmployee struct {
	Resource             Resource `xml:"Resource"`
	TerritoryCode        string   `xml:"TerritoryCode,omitempty"`
	CommissionPercentage float64  `xml:"CommissionPercentage"`
}
