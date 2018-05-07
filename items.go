package globe

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/tim-online/go-exactglobe/omitempty"
)

type Items []Item

type Item struct {
	// Attributs
	// code = string
	// [ type = token: { S | B | C | L | M | P | R } ]
	// [ searchcode = string ]>

	// <Description> (0,1)
	// <LongDescription> (0,1)
	// <MultiDescriptions> (0,1)
	// <Assortment> (0,1)
	// <Availability> (0,1)
	// <Condition> (0,1)
	// <IsSalesItem> (0,1)
	// <IsPurchaseItem> (0,1)
	// <IsSerialNumberItem> (0,1)
	// <IsBatchItem> (0,1)
	// <IsSubAssemblyItem> (0,1)
	// <IsAssembledItem> (0,1)
	// <IsStockItem> (0,1)
	// <IsBackOrderItem> (0,1)
	// <IsFractionAllowedItem> (0,1)
	// <IsPriceRegulationItem> (0,1)
	// <IsTextItem> (0,1)
	// <IsDiscountItem> (0,1)
	// <IsExplodeItem> (0,1)
	// <IsPrintItem> (0,1)
	// <IsOutsourcedItem> (0,1)
	// <RequiresApprovedSupplier> (0,1)
	// <GLRevenue> (0,1)
	// <GLCosts> (0,1)
	// <GLPurchase> (0,1)
	// <GLAsset> (0,1)
	// <GLItemsToBeReceived> (0,1)
	// <GLAccountDiscount> (0,1)
	// <Sales> (0,1)
	// <Purchase> (0,1)
	// <Costs> (0,1)
	// <Dimension> (0,1)
	// <Statistical> (0,1)
	// <ValuationMethod> (0,1)
	// <BOMs> (0,1)
	// <Resource> (0,1)
	// <Image> (0,1)
	// <Note> (0,1)
	// <FreeFields> (0,1)
	// <ItemTexts> (0,1)
	// <ItemCategory> (0,+)
	// <ItemAccounts> (0,1)
	// <ItemWarehouses> (0,1)
	// <ItemRelations> (0,1)
	// <ShelfLife> (0,1)
	// <Warranty> (0,1)
	// <IsServiceItem> (0,1)
	// <IntrastatEnabled> (0,1)
	// <AddExtraReceiptToOrder> (0,1)
	// <SerialBatchMask> (0,1)
	// <IncrementFactor> (0,1)
	// <IsCommissionable> (0,1)
	// <CommissionMethod> (0,1)
	// <CommissionValue> (0,1)
	// <TaxItemClassification> (0,1)
	// <Charges> (0,1)
	// <HTS> (0,1)
	// <CSCustomField> (0,1)
	// <GTIN> (0,1)
}

func (i Item) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(i, e, start)
}

func (i Item) IsEmpty() bool {
	return zero.IsZero(i)
}
