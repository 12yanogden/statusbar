package statusbar

import (
	"testing"
	"time"
)

func TestStatusBar(t *testing.T) {
	var bar StatusBar

	bar.Start("Gotta get stuff done")

	sleepSeconds(3)

	bar.Pass()
}

func sleepSeconds(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}
