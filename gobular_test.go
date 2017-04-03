package gobular

import (
	"testing"
)

func testFit(mustLen uint32, align Alignment, s string) string {
	return ansiEsc.ReplaceAllLiteralString(fitPad(mustLen, align, s), "")
}

// TODO: simplify by extracting single test function
// TODO: test multibyte and/or double-width utf8 characters
func TestFitPad(t *testing.T) {
	s := "123456"
	must := "123456"
	if got := testFit(6, HAlignLeft, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "123456"
	if got := testFit(6, HAlignCenter, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "123456"
	if got := testFit(6, HAlignRight, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "123456    "
	if got := testFit(10, HAlignLeft, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "  123456  "
	if got := testFit(10, HAlignCenter, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "    123456"
	if got := testFit(10, HAlignRight, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "123456     "
	if got := testFit(11, HAlignLeft, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "  123456   "
	if got := testFit(11, HAlignCenter, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "     123456"
	if got := testFit(11, HAlignRight, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "12"
	if got := testFit(2, HAlignLeft, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "34"
	if got := testFit(2, HAlignCenter, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "56"
	if got := testFit(2, HAlignRight, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "123"
	if got := testFit(3, HAlignLeft, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "345"
	if got := testFit(3, HAlignCenter, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "456"
	if got := testFit(3, HAlignRight, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	s = "123\033[38;5;226m456"

	must = "123456"
	if got := testFit(6, HAlignLeft, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "123456"
	if got := testFit(6, HAlignCenter, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "123456"
	if got := testFit(6, HAlignRight, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "123456    "
	if got := testFit(10, HAlignLeft, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "  123456  "
	if got := testFit(10, HAlignCenter, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "    123456"
	if got := testFit(10, HAlignRight, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "123456     "
	if got := testFit(11, HAlignLeft, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "  123456   "
	if got := testFit(11, HAlignCenter, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "     123456"
	if got := testFit(11, HAlignRight, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "12"
	if got := testFit(2, HAlignLeft, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "34"
	if got := testFit(2, HAlignCenter, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "56"
	if got := testFit(2, HAlignRight, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "123"
	if got := testFit(3, HAlignLeft, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "345"
	if got := testFit(3, HAlignCenter, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}

	must = "456"
	if got := testFit(3, HAlignRight, s); got != must {
		t.Errorf("Must:|%s| Got:|%s|\n", must, got)
	}
}
