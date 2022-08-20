package EvilGo

import (
	"fmt"
	"github.com/MXuDong/test-utils-go/patch"
	"os"
	"time"
)

func init() {
}

const defaultReplaceWeight int = 10

var alreadyInjection bool = false
var printlnValue string = "Evil for go is inject ;-）"

func EvilPrintf(replaceWeight int, replaceValue string, unReplaceDuration time.Duration) {
	if replaceWeight >= 100 || replaceWeight < 0 {
		replaceWeight = defaultReplaceWeight
	}

	printlnValue = replaceValue

	rv := randInstance.Int()

	if rv%100 < replaceWeight && !alreadyInjection {
		r := patch.Cover(fmt.Println, PrintlnReplacer)
		alreadyInjection = true

		// do reset with duration
		go func() {
			time.Sleep(unReplaceDuration)
			r.Restore()
			alreadyInjection = false
		}()
	}
}

func PrintlnReplacer(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stdout, printlnValue)
}
