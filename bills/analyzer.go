package bills

import (
	"github.com/chamzzzzzz/ocr"
	"strings"
)

var (
	AliPayExpenseCharacteristics               = []string{"账单详情", "付款方式", "商品说明", "创建时间"}
	WechatPayExpenseCharacteristics            = []string{"全部账单", "当前状态", "商户全称", "支付时间", "支付方式", "交易单号", "商户单号"}
	UnionPayCreditCardRepaymentCharacteristics = []string{"还款详情", "订单类型", "还款卡号", "付款卡号", "创建时间"}
)

func CharacteristicsMatching(characteristics []string, result *ocr.Result) bool {
	for _, characteristic := range characteristics {
		matching := false
		for _, observation := range result.Observations {
			if observation.Text == characteristic {
				matching = true
				break
			}
		}
		if !matching {
			return false
		}
	}
	return true
}

type Analyzer interface {
	Matching(result *ocr.Result) bool
	Analyze(result *ocr.Result) (*Bill, error)
}

type AliPayExpenseAnalyzer struct {
}

func (analyzer *AliPayExpenseAnalyzer) Matching(result *ocr.Result) bool {
	return CharacteristicsMatching(AliPayExpenseCharacteristics, result)
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
	return CharacteristicsMatching(WechatPayExpenseCharacteristics, result)
}

func (analyzer *WechatPayExpenseAnalyze) Analyze(result *ocr.Result) (*Bill, error) {
	bill := &Bill{
		Platform:       WechatPay,
		Classification: Expense,
	}
	bill.PaymentAccount, _ = ExtractHorizontalNextItem(result, "支付方式")
	bill.Date, _ = ExtractHorizontalNextItem(result, "支付时间")
	bill.Merchant, _ = ExtractHorizontalNextItem(result, "商户全称")
	if HasItem(result, "原价") {
		bill.Amount, _ = ExtractHorizontalPreviousItem(result, "原价")
	} else {
		bill.Amount, _ = ExtractHorizontalPreviousItem(result, "当前状态")
	}
	return bill, nil
}

type UnionPayCreditCardRepaymentAnalyzer struct {
}

func (analyzer *UnionPayCreditCardRepaymentAnalyzer) Matching(result *ocr.Result) bool {
	return CharacteristicsMatching(UnionPayCreditCardRepaymentCharacteristics, result)
}

func (analyzer *UnionPayCreditCardRepaymentAnalyzer) Analyze(result *ocr.Result) (*Bill, error) {
	bill := &Bill{
		Platform:       UnionPay,
		Classification: CreditCardRepayment,
	}

	bill.ReceiptAccount, _ = ExtractHorizontalNextItem(result, "还款卡号")
	bill.PaymentAccount, _ = ExtractHorizontalNextItem(result, "付款卡号")
	bill.Date, _ = ExtractHorizontalNextItem(result, "创建时间")
	bill.Amount, _ = ExtractColonJoinedItem(result, "还款金额")
	return bill, nil
}

func HasItem(result *ocr.Result, itemName string) bool {
	for _, observation := range result.Observations {
		if observation.Text == itemName {
			return true
		}
	}
	return false
}

func ExtractHorizontalOffsetItem(result *ocr.Result, itemName string, offset int) (string, bool) {
	ok := false
	itemValue := ""
	for i, observation := range result.Observations {
		if observation.Text == itemName {
			j := i + offset
			if j >= 0 && j < len(result.Observations) {
				ok = true
				itemValue = result.Observations[j].Text
			}
		}
	}
	return itemValue, ok
}

func ExtractHorizontalNextItem(result *ocr.Result, itemName string) (string, bool) {
	return ExtractHorizontalOffsetItem(result, itemName, 1)
}

func ExtractHorizontalPreviousItem(result *ocr.Result, itemName string) (string, bool) {
	return ExtractHorizontalOffsetItem(result, itemName, -1)
}

func ExtractSeparatorJoinedItem(result *ocr.Result, itemName string, separator string) (string, bool) {
	ok := false
	itemValue := ""
	for _, observation := range result.Observations {
		itemNameValue := strings.SplitN(observation.Text, separator, 2)
		if len(itemNameValue) == 2 {
			if itemNameValue[0] == itemName {
				ok = true
				itemValue = itemNameValue[1]
				break
			}
		}
	}
	return itemValue, ok
}

func ExtractColonJoinedItem(result *ocr.Result, itemName string) (string, bool) {
	return ExtractSeparatorJoinedItem(result, itemName, "：")
}
