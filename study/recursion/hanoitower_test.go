package recursion

import (
	"fmt"
	"testing"
	"time"
)

func TestHanoitower(t *testing.T) {
	h := NewHanoitower()
	fmt.Println("---")
	now := time.Now()
	h.Move(21)
	h.Print()
	fmt.Println(time.Now().Sub(now))
}
