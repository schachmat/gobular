package gobular

import (
	"testing"
)

func testFit(t *testing.T, s string, must string, len uint32, align Alignment) {
	got := ansiEsc.ReplaceAllLiteralString(fitPad(len, align, s), "")
	if got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}
}

// TODO: test multibyte and/or double-width utf8 characters
func TestFitPad(t *testing.T) {
	testFit(t, "123456", "123456", 6, HAlignLeft)
	testFit(t, "123456", "123456", 6, HAlignCenter)
	testFit(t, "123456", "123456", 6, HAlignRight)
	testFit(t, "123456", "123456    ", 10, HAlignLeft)
	testFit(t, "123456", "  123456  ", 10, HAlignCenter)
	testFit(t, "123456", "    123456", 10, HAlignRight)
	testFit(t, "123456", "123456     ", 11, HAlignLeft)
	testFit(t, "123456", "  123456   ", 11, HAlignCenter)
	testFit(t, "123456", "     123456", 11, HAlignRight)
	testFit(t, "123456", "12", 2, HAlignLeft)
	testFit(t, "123456", "34", 2, HAlignCenter)
	testFit(t, "123456", "56", 2, HAlignRight)
	testFit(t, "123456", "123", 3, HAlignLeft)
	testFit(t, "123456", "345", 3, HAlignCenter)
	testFit(t, "123456", "456", 3, HAlignRight)

	testFit(t, "123\033[38;5;226m456", "123456", 6, HAlignLeft)
	testFit(t, "123\033[38;5;226m456", "123456", 6, HAlignCenter)
	testFit(t, "123\033[38;5;226m456", "123456", 6, HAlignRight)
	testFit(t, "123\033[38;5;226m456", "123456    ", 10, HAlignLeft)
	testFit(t, "123\033[38;5;226m456", "  123456  ", 10, HAlignCenter)
	testFit(t, "123\033[38;5;226m456", "    123456", 10, HAlignRight)
	testFit(t, "123\033[38;5;226m456", "123456     ", 11, HAlignLeft)
	testFit(t, "123\033[38;5;226m456", "  123456   ", 11, HAlignCenter)
	testFit(t, "123\033[38;5;226m456", "     123456", 11, HAlignRight)
	testFit(t, "123\033[38;5;226m456", "12", 2, HAlignLeft)
	testFit(t, "123\033[38;5;226m456", "34", 2, HAlignCenter)
	testFit(t, "123\033[38;5;226m456", "56", 2, HAlignRight)
	testFit(t, "123\033[38;5;226m456", "123", 3, HAlignLeft)
	testFit(t, "123\033[38;5;226m456", "345", 3, HAlignCenter)
	testFit(t, "123\033[38;5;226m456", "456", 3, HAlignRight)
}
