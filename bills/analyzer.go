package bills

import (
	"github.com/chamzzzzzz/ocr"
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
	return bill, nil
}
