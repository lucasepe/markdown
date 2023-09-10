package markdown

import (
	"strconv"
	"strings"
)

func parseImageSize(txt string) (w int, h int) {
	if strings.IndexByte(txt, '=') != 0 {
		return
	}

	parts := strings.Split(txt[1:], "x")
	if len(parts) != 2 {
		return
	}

	if x, err := strconv.Atoi(parts[0]); err == nil {
		w = x
	}

	if y, err := strconv.Atoi(parts[1]); err == nil {
		h = y
	}

	return
}
