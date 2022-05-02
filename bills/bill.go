package bills

type Platform int

const (
	AliPay Platform = iota
	WechatPay
	UnionPay
)

type Classification int

const (
	Income Classification = iota
	Expense
	Transfer
	Refund
	CreditCardRepayment
)

type DetailClassification int

type Bill struct {
	Platform             Platform
	Classification       Classification
	DetailClassification DetailClassification
	PaymentAccount       string
	ReceiptAccount       string
	Amount               string
	Merchant             string
	Description          string
	Date                 string
}
