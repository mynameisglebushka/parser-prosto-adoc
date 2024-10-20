package ast


type Type string

const (
	BlockType  Type = "block"
	StringType Type = "string"
	InlineType Type = "inline"
)

type Name string

const (
	DocumentName Name = "document" // Document Name

	SectionName Name = "section" // Section Name

	ListName     Name = "list"     // List Name
	ListItemName Name = "listItem" // List Item Name

	DListName     Name = "dlist"     // DList Name
	DListItemName Name = "dlistItem" // DList Item Name

	DiscreteHeadingName Name = "heading" // Discrete Heading Name

    BreakName Name = "break" // Break Name

	ListingName   Name = "listing"   // Leaf Block Name
	LiteralName   Name = "literal"   // Leaf Block Name
	ParagraphName Name = "paragraph" // Leaf Block Name
	PassName      Name = "pass"      // Leaf Block Name
	StemName      Name = "stem"      // Leaf Block Name
	VerseName     Name = "verse"     // Leaf Block Name

	AudioName Name = "audio" // Block Macro Name
	VideoName Name = "video" // Block Macro Name
	ImageName Name = "image" // Block Macro Name
	TocName   Name = "toc"   // Block Macro Name

	AdmonitionName Name = "admonition" // Parent Block Name
	ExampleName    Name = "example"    // Parent Block Name
	SidebarName    Name = "sidebar"    // Parent Block Name
	OpenName       Name = "open"       // Parent Block Name
	QuoteName      Name = "quote"      // Parent Block Name

	RefName     Name = "ref"     // Inline Ref Name

	SpanName    Name = "span"    // Inline Span Name
    
	TextName    Name = "text"    // Inline Literal Name
	CharRefName Name = "charref" // Inline Literal Name
	RawName     Name = "raw"     // Inline Literal Name
)

type Variant string

const (
	CalloutVariant   Variant = "callout"   // List "callout" Variant
	OrderedVariant   Variant = "ordered"   // List "ordered" Variant
	UnorderedVariant Variant = "unordered" // List "unordered" Variant

	PageVariant     Variant = "page"     // Break "page" Variant
	ThematicVariant Variant = "thematic" // Break "thematic" Variant

	CautionVariant   Variant = "caution"   // Parent "caution" Variant
	ImportantVariant Variant = "important" // Parent "important" Variant
	NoteVariant      Variant = "note"      // Parent "note" Variant
	TipVariant       Variant = "tip"       // Parent "tip" Variant
	WarningVariant   Variant = "warning"   // Parent "warning" Variant

	StrongVariant   Variant = "strong"   // Inline Span "strong" Variant
	EmphasisVariant Variant = "emphasis" // Inline Span "emphasis" Variant
	CodeVariant     Variant = "code"     // Inline Span "code" Variant
	MarkVariant     Variant = "mark"     // Inline Span "mark" Variant

	LinkVariant Variant = "link" // Inline Ref "link" Variant
	XRefVariant Variant = "xref" // Inline Ref "ref" Variant
)

type Form string

const (
	DelimitedForm Form = "delimited" // Leaf / Parent Block "delimited" Form
	IndentedForm  Form = "indented"  // Leaf Block "indented" Form
	ParagraphForm Form = "paragraph" // Leaf Block "paragraph" Form

	MacroForm Form = "macro" // Block Macro "macro" Form

	ConstrainedForm   Form = "constrained"   // Inline Span "constrained" Form
	UnConstrainedForm Form = "unconstrained" // Inline Span "unconstrained" Form
)
