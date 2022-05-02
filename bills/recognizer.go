package bills

import (
	"errors"
	"github.com/chamzzzzzz/ocr"
)

var (
	ErrMatchingAnalyzerNotFound = errors.New("matching analyzer not found")
)

var analyzers = []Analyzer{
	&AliPayExpenseAnalyzer{},
	&WechatPayExpenseAnalyze{},
	&UnionPayCreditCardRepaymentAnalyzer{},
}

type Recognizer struct {
}

func (recognizer *Recognizer) Recognize(file string) (*Bill, error) {
	recognizer_ := &ocr.MacRecognizer{}
	result, err := recognizer_.Recognize(file)
	if err != nil {
		return nil, err
	}
	return recognizer.Analyze(result)
}

func (recognizer *Recognizer) Analyze(result *ocr.Result) (*Bill, error) {
	analyzer, err := recognizer.MatchingAnalyzer(result)
	if err != nil {
		return nil, err
	}
	return analyzer.Analyze(result)
}

func (recognizer *Recognizer) MatchingAnalyzer(result *ocr.Result) (Analyzer, error) {
	for _, analyzer := range analyzers {
		if analyzer.Matching(result) {
			return analyzer, nil
		}
	}
	return nil, ErrMatchingAnalyzerNotFound
}
