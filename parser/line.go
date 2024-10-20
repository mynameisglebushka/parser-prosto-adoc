package parser

import (
	"bytes"
	"regexp"
)

type line struct {
	spases  []byte
	content []byte
}

func (l *line) lineKindOf() Kind {

	defaultKind := kindText

	if len(l.content) == 0 {
		return lineEmpty
	}

	if bytes.Equal(l.content, []byte("////")) {
		return lineMultilineComment
	}

	if bytes.HasPrefix(l.content, []byte("//")) {
		return lineComment
	}

	if len(l.spases) > 0 {
		if l.content[0] != '.' && l.content[0] != '*' && l.content[0] != '-' {
			return blockLiteralParagraph
		}
	}

	switch l.content[0] {
	case ':':
		ok, _ := regexp.Match(`^:[a-zA-Z0-9_][-a-zA-Z0-9_]*:`, l.content)
		if ok {
			return lineKindAttribute
		}
	case '=', '#':
		subs := bytes.Fields(l.content)

		switch string(subs[0]) {
		case "=", "#":
			return kindDocumentTitle
		case "==", "##":
			return kindSectionTitleL1
		case "===", "###":
			return kindSectionTitleL2
		case "====", "####":
			return kindSectionTitleL3
		case "=====", "#####":
			return kindSectionTitleL4
		case "======", "######":
			return kindSectionTitleL5
		default:
			return defaultKind
		}
	}

	return defaultKind
}
