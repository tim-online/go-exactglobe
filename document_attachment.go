package globe

type DocumentAttachment struct {
	// [ number = integer ]
	// [ code = string ]
	// [ publish = token: { 0 | 1 | 2 } ]
	// [ ID = string ]>
	// sequence
	// (0,1)	 	<DocumentType> (0,1)
	// <DocumentGroup> (0,1)
	// <DocumentCategory> (0,1)
	// <DocumentSubCategory> (0,1)
	// <ParentDocument> (0,1)
	// <Subject>
	// <Summary> (0,1)
	// <SecurityLevel> (0,1)
	// <Division> (0,1)
	// <Manager> (0,1)
	// <Language> (0,1)
	// <Project> (0,1)
	// <Resource> (0,1)
	// <Account> (0,1)
	// <Assortment> (0,1)
	// <Item> (0,1)
	// <SerialNumber> (0,1)
	// <OrderNumber> (0,1)
	// <OurRef> (0,1)
	// <YourRef> (0,1)
	// <Body> (0,1)
	// <Attachment> (0,1)
	// <Images> (0,1)
	// <Version> (0,1)
	// <MaxVersion> (0,1)
	// <StartDate> (0,1)
	// <FinEntry> (0,1)
	// <Description> (0,1)
	// <GenerateVersion> (0,1)
}
