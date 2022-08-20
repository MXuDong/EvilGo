package EvilGo

import (
	"fmt"
	"testing"
	"time"
)

func TestEvil(t *testing.T) {
	Evil()

	for ;; {
		time.Sleep(1 * time.Second)
		fmt.Println(time.Now())
	}


}
