package parser

type Kind string

const (
	lineEmpty             Kind = "empty line"
	kindText              Kind = "text line"
	kindDocumentTitle     Kind = "document title"          // =
	kindSectionTitleL1    Kind = "section title level 1"   // ==
	kindSectionTitleL2    Kind = "section title level 2"   // ===
	kindSectionTitleL3    Kind = "section title level 3"   // ====
	kindSectionTitleL4    Kind = "section title level 4"   // =====
	kindSectionTitleL5    Kind = "section title level 5"   // ======
	blockLiteralParagraph Kind = "paragraph block"         // Line start with space
	lineKindAttribute     Kind = "document attribute line" // line match "^:[a-zA-Z0-9_][-a-zA-Z0-9_]*:"
	lineComment           Kind = "inline comment"          // line like "// .*"
	lineMultilineComment  Kind = "block comment"           // line like "////"
)
