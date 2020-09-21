package accounting

//LineItem is a line containing detail on an Invoice
type LineItem struct {
	//The Xero generated identifier for a LineItem. It is recommended that you include LineItemIDs on update requests. If LineItemIDs are not included with line items in an update request then the line items are deleted and recreated.
	LineItemID string `json:"LineItemID,omitempty" xml:"LineItemID,omitempty"`

	// Description needs to be at least 1 char long. A line item with just a description (i.e no unit amount or quantity) can be created by specifying just a <Description> element that contains at least 1 character
	Description string `json:"Description,omitempty" xml:"Description,omitempty"`

	// LineItem Quantity
	Quantity float64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`

	// LineItem Unit Amount
	UnitAmount Float `json:"UnitAmount" xml:"UnitAmount"`

	// See Items
	ItemCode string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`

	// See Accounts
	AccountCode string `json:"AccountCode,omitempty" xml:"AccountCode,omitempty"`

	// Used as an override if the default Tax Code for the selected <AccountCode> is not correct – see TaxTypes.
	TaxType string `json:"TaxType,omitempty" xml:"TaxType,omitempty"`

	// The tax amount is auto calculated as a percentage of the line amount (see below) based on the tax rate. This value can be overriden if the calculated <TaxAmount> is not correct.
	TaxAmount Float `json:"TaxAmount,omitempty" xml:"TaxAmount,omitempty"`

	// If you wish to omit either of the <Quantity> or <UnitAmount> you can provide a LineAmount and Xero will calculate the missing amount for you. The line amount reflects the discounted price if a DiscountRate has been used . i.e LineAmount = Quantity * Unit Amount * ((100 – DiscountRate)/100)
	LineAmount Float `json:"LineAmount,omitempty" xml:"LineAmount,omitempty"`

	// Optional Tracking Category – see Tracking.  Any LineItem can have a maximum of 2 <TrackingCategory> elements.
	Tracking []TrackingCategory `json:"Tracking,omitempty" xml:"Tracking>TrackingCategory,omitempty"`

	// Percentage discount being applied to a line item (only supported on ACCREC invoices – ACC PAY invoices and credit notes in Xero do not support discounts
	DiscountRate float64 `json:"DiscountRate,omitempty" xml:"DiscountRate,omitempty"`

	// The Xero identifier for a Repeating Invoicee.g. 297c2dc5-cc47-4afd-8ec8-74990b8761e9
	RepeatingInvoiceID string `json:"RepeatingInvoiceID,omitempty" xml:"RepeatingInvoiceID,omitempty"`
}
