package EvilGo

import (
	"math/rand"
	"time"
)

var randInstance *rand.Rand

// init to load rander funcs.
func init() {
	randInstance = rand.New(rand.NewSource(time.Now().Unix()))

}

// EvilConfig enable to control the behavior of the EvilGo. When Evil start, you can't disable in anywhere.
type EvilConfig struct {
	BlockTimeConfig     BlockTimeConfig
	ReplacePrintfConfig ReplacePrintfConfig
}

type ReplacePrintfConfig struct {
	Weight   int
	Value    string
	Duration time.Duration
}

type BlockTimeConfig struct {
	Weight        int
	BlockTime     time.Time
	BlockDuration time.Duration
}

type EvilOptions func(e *EvilConfig)

func SetBlockTimeWeight(weight int) EvilOptions {
	return func(e *EvilConfig) {
		e.BlockTimeConfig.Weight = weight
	}
}
func SetReplacePrintlnWeight(weight int) EvilOptions {
	return func(e *EvilConfig) {
		e.ReplacePrintfConfig.Weight = weight
	}
}
func SetBlockTimeDuration(d time.Duration) EvilOptions {
	return func(e *EvilConfig) {
		e.BlockTimeConfig.BlockDuration = d
	}
}
func SetReplacePrintlnDuration(d time.Duration) EvilOptions {
	return func(e *EvilConfig) {
		e.ReplacePrintfConfig.Duration = d
	}
}
func SetBlockTime(t time.Time) EvilOptions {
	return func(e *EvilConfig) {
		e.BlockTimeConfig.BlockTime = t
	}
}
func SetReplacePrintlnValue(value string) EvilOptions {
	return func(e *EvilConfig) {
		e.ReplacePrintfConfig.Value = value
	}
}

// Evil will start the rand bugs function. Because the golang load model, the developers need invoke this function by
// hand.
func Evil(options ...EvilOptions) {

	config := &EvilConfig{
		BlockTimeConfig: BlockTimeConfig{
			Weight:        1,
			BlockTime:     time.Date(0, 0, 0, 0, 0, 0, 0, time.Local),
			BlockDuration: 5 * time.Second,
		},
		ReplacePrintfConfig: ReplacePrintfConfig{
			Weight:   1,
			Value:    "Evil for go is inject ;-）",
			Duration: 5 * time.Second,
		},
	}

	for _, option := range options {
		option(config)
	}

	go func() {
		for {
			time.Sleep(1 * time.Second)

			// block time
			BlockTimeRandom(
				config.BlockTimeConfig.Weight,
				config.BlockTimeConfig.BlockTime,
				config.BlockTimeConfig.BlockDuration,
			)

			EvilPrintf(
				config.ReplacePrintfConfig.Weight,
				config.ReplacePrintfConfig.Value,
				config.ReplacePrintfConfig.Duration,
			)
		}
	}()
}
