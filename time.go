package EvilGo

import (
	"github.com/MXuDong/test-utils-go/base"
	"time"
)



const defaultBlockTimeWeight int = 4

func init() {
}

var frozen bool = false

func BlockTimeRandom(blockWeight int, blockTime time.Time, unblockDuration time.Duration) {
	if blockWeight >= 100 || blockWeight < 0 {
		blockWeight = defaultBlockTimeWeight
	}

	rv := randInstance.Int()
	if rv%100 < blockWeight && !frozen {
		_ = base.FreezeWithTimeStruct(blockTime)
		frozen = true

		// do reset with duration
		go func() {
			time.Sleep(unblockDuration)
			base.UnFreezeTime()
			frozen = false
		}()
	}
}
