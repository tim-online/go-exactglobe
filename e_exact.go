package globe

type EExact struct {
	// Attributes
	SchemaVersion string `xml:"schemaVersion,attr"`

	// Nodes
	Settings              Settings              `xml:"Settings,omitempty"`
	ExchangeRates         ExchangeRates         `xml:"ExchangeRates,omitempty"`
	Accounts              Accounts              `xml:"Accounts,omitempty"`
	Assortments           Assortments           `xml:"Assortments,omitempty"`
	Items                 Items                 `xml:"Items,omitempty"`
	ItemPrices            ItemPrices            `xml:"ItemPrices,omitempty"`
	ItemNumbers           ItemNumbers           `xml:"ItemNumbers,omitempty"`
	Resources             Resources             `xml:"Resources,omitempty"`
	GLAccounts            GLAccounts            `xml:"GLAccounts,omitempty"`
	Journals              Journals              `xml:"Journals,omitempty"`
	VATs                  VATs                  `xml:"VATs,omitempty"`
	FinEntries            FinEntries            `xml:"FinEntries,omitempty"`
	GLEntries             GLEntries             `xml:"GLEntries,omitempty"`
	Budgets               Budgets               `xml:"Budgets,omitempty"`
	InternalUses          InternalUses          `xml:"InternalUses,omitempty"`
	DocumentTypes         DocumentTypes         `xml:"DocumentTypes,omitempty"`
	DocumentGroups        DocumentGroups        `xml:"DocumentGroups,omitempty"`
	Documents             Documents             `xml:"Documents,omitempty"`
	Projects              Projects              `xml:"Projects,omitempty"`
	Assets                Assets                `xml:"Assets,omitempty"`
	Invoices              Invoices              `xml:"Invoices,omitempty"`
	Orders                Orders                `xml:"Orders,omitempty"`
	ProductionOrders      ProductionOrders      `xml:"ProductionOrders,omitempty"`
	Contracts             Contracts             `xml:"Contracts,omitempty"`
	Requests              Requests              `xml:"Requests,omitempty"`
	ResourceRoles         ResourceRoles         `xml:"ResourceRoles,omitempty"`
	PayrollComponentTypes PayrollComponentTypes `xml:"PayrollComponentTypes,omitempty"`
	PayrollResources      PayrollResources      `xml:"PayrollResources,omitempty"`
	PayrollTransactions   PayrollTransactions   `xml:"PayrollTransactions,omitempty"`
	Costcenters           Costcenters           `xml:"Costcenters,omitempty"`
	Costunits             Costunits             `xml:"Costunits,omitempty"`
	Currencies            Currencies            `xml:"Currencies,omitempty"`
	Topics                Topics                `xml:"Topics,omitempty"`
	Messages              Messages              `xml:"Messages,omitempty"`
}
