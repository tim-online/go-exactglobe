package globe

type eExact struct {
	// Attributes
	SchemaVersion string `xml:"schemaVersion"`

	// Nodes
	Settings              Settings
	ExchangeRates         ExchangeRates         `xml:"ExchangeRates"`
	Accounts              Accounts              `xml:"Accounts"`
	Assortments           Assortments           `xml:"Assortments"`
	Items                 Items                 `xml:"Items"`
	ItemPrices            ItemPrices            `xml:"ItemPrices"`
	ItemNumbers           ItemNumbers           `xml:"ItemNumbers"`
	Resources             Resources             `xml:"Resources"`
	GLAccounts            GLAccounts            `xml:"GLAccounts"`
	Journals              Journals              `xml:"Journals"`
	VATs                  VATs                  `xml:"VATs"`
	FinEntries            FinEntries            `xml:"FinEntries"`
	GLEntries             GLEntries             `xml:"GLEntries"`
	Budgets               Budgets               `xml:"Budgets"`
	InternalUses          InternalUses          `xml:"InternalUses"`
	DocumentTypes         DocumentTypes         `xml:"DocumentTypes"`
	DocumentGroups        DocumentGroups        `xml:"DocumentGroups"`
	Documents             Documents             `xml:"Documents"`
	Projects              Projects              `xml:"Projects"`
	Assets                Assets                `xml:"Assets"`
	Invoices              Invoices              `xml:"Invoices"`
	Orders                Orders                `xml:"Orders"`
	ProductionOrders      ProductionOrders      `xml:"ProductionOrders"`
	Contracts             Contracts             `xml:"Contracts"`
	Requests              Requests              `xml:"Requests"`
	ResourceRoles         ResourceRoles         `xml:"ResourceRoles"`
	PayrollComponentTypes PayrollComponentTypes `xml:"PayrollComponentTypes"`
	PayrollResources      PayrollResources      `xml:"PayrollResources"`
	PayrollTransactions   PayrollTransactions   `xml:"PayrollTransactions"`
	Costcenters           Costcenters           `xml:"Costcenters"`
	Costunits             Costunits             `xml:"Costunits"`
	Currencies            Currencies            `xml:"Currencies"`
	Topics                Topics                `xml:"Topics"`
	Messages              Messages              `xml:"Messages"`
}
