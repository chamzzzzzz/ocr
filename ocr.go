package ocr

import (
	"fmt"
	"github.com/chamzzzzzz/ocr/recognizer"
	"github.com/chamzzzzzz/ocr/recognizer/mac"
	"sort"
	"sync"
)

var (
	recognizersMu sync.RWMutex
	recognizers   = make(map[string]recognizer.Recognizer)
)

// Register makes a ocr recognizer available by the provided name.
// If Register is called twice with the same name or if recognizer is nil,
// it panics.
func Register(name string, recognizer recognizer.Recognizer) {
	recognizersMu.Lock()
	defer recognizersMu.Unlock()
	if recognizer == nil {
		panic("register recognizer is nil")
	}
	if _, dup := recognizers[name]; dup {
		panic("register called twice for recognizer " + name)
	}
	recognizers[name] = recognizer
}

// Recognizers returns a sorted list of the names of the registered recognizers.
func Recognizers() []string {
	recognizersMu.RLock()
	defer recognizersMu.RUnlock()
	list := make([]string, 0, len(recognizers))
	for name := range recognizers {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

func Recognize(name string, file string) (*recognizer.Result, error) {
	recognizersMu.RLock()
	recognizer, ok := recognizers[name]
	recognizersMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("unknown recognizer %q (forgotten import?)", name)
	}
	return recognizer.Recognize(file)
}

func init() {
	Register("mac", &mac.Recognizer{})
}
