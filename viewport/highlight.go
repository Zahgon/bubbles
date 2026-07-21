package viewport

import (
	"charm.land/lipgloss/v2"
)

func parseMatches(
	content string,
	matches [][]int,
) []highlightInfo {
	_ = "STUB: not implemented"
	return nil
}

type highlightInfo struct {
	lineStart, lineEnd int

	lines map[int][2]int
}

func (hi highlightInfo) coords() (int, int, int) { _ = "STUB: not implemented"; return 0, 0, 0 }

func makeHighlightRanges(
	highlights []highlightInfo,
	line int,
	style lipgloss.Style,
) []lipgloss.Range {
	_ = "STUB: not implemented"
	return nil
}
