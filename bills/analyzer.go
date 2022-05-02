package bills

import (
	"github.com/chamzzzzzz/ocr"
)

type Analyzer interface {
	Matching(result *ocr.Result) bool
	Analyze(result *ocr.Result) (*Bill, error)
}

type AliPayExpenseAnalyzer struct {
}

func (analyzer *AliPayExpenseAnalyzer) Matching(result *ocr.Result) bool {
	return true
}

func (analyzer *AliPayExpenseAnalyzer) Analyze(result *ocr.Result) (*Bill, error) {
	bill := &Bill{
		Platform:       AliPay,
		Classification: Expense,
	}
	return bill, nil
}

type WechatPayExpenseAnalyze struct {
}

func (analyzer *WechatPayExpenseAnalyze) Matching(result *ocr.Result) bool {
	return true
}

func (analyzer *WechatPayExpenseAnalyze) Analyze(result *ocr.Result) (*Bill, error) {
	bill := &Bill{
		Platform:       WechatPay,
		Classification: Expense,
	}
	return bill, nil
}

type UnionPayCreditCardRepaymentAnalyzer struct {
}

func (analyzer *UnionPayCreditCardRepaymentAnalyzer) Matching(result *ocr.Result) bool {
	return true
}

func (analyzer *UnionPayCreditCardRepaymentAnalyzer) Analyze(result *ocr.Result) (*Bill, error) {
	bill := &Bill{
		Platform:       UnionPay,
		Classification: CreditCardRepayment,
	}
	return bill, nil
}
