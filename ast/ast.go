package ast

type Document struct {
	Type Type // DocumentName
	Name Name // BlockType

	Attributes map[string]string // key pattern ^[a-zA-Z0-9_][-a-zA-Z0-9_]*$

	Header *Header

	Blocks []Block

	Location Location
}

type Header struct {
	Title   Inlines
	Authors []Author

	Location Location
}

type AbstructBlock struct {
	Type     Type // BlockType
	Id       string
	Title    Inlines
	RefText  Inlines
	MetaData BlockMetaData

	Location Location
}

type AbstractHeading struct {
	Level int
	AbstructBlock
}

type AbstractListItem struct {
	Marker    string
	Principal Inlines
	Blocks    []Block

	AbstructBlock
}

type Block interface {
	block()
}

func (b *List) block()            {}
func (b *DescriptionList) block() {}
func (b *DiscreteHeading) block() {}
func (b *Break) block()           {}
func (b *BlockMacro) block()      {}
func (b *LeafBlock) block()       {}
func (b *ParentBlock) block()     {}

type Section struct {
	Name   Name // SectionName
	Blocks []Block

	AbstractHeading
}

type List struct {
	Name    Name // ListName
	Marker  string
	Variant Variant
	Items   []ListItem

	AbstructBlock
}

type DescriptionList struct {
	Name   Name // DListName
	Marker string

	AbstructBlock
}

type ListItem struct {
	Name Name // ListItemName

	AbstractListItem
}

type DescriptionListItem struct {
	Name  Name // DListItemName
	Terms []Inlines

	AbstractListItem
}

type DiscreteHeading struct {
	Name Name

	AbstractHeading
}

type Break struct {
	Name    Name // BreakName
	Variant Variant

	AbstructBlock
}

type BlockMacro struct {
	Name   Name
	Form   Form
	Target string

	AbstructBlock
}

// If "form"="delimiter" then Delimiter required
type LeafBlock struct {
	Name      Name
	Form      Form
	Inlines   []Inlines
	Delimiter string

	AbstructBlock
}

// if "name"="admonition" then Variant required
type ParentBlock struct {
	Name      Name
	Form      Form
	Delimiter string
	Blocks    []Block
	Variant   Variant

	AbstructBlock
}

type BlockMetaData struct {
	Attributes map[string]string // key pattern ^(?:[a-zA-Z_][a-zA-Z0-9_-]*|\\$[1-9][0-9]*)$
	Options    []string
	Roles      []string

	Location Location
}

type Inlines []Inline

type Inline interface {
	inline()
}

func (i *InlineSpan) inline()    {}
func (i *InlineRef) inline()     {}
func (i *InlineLiteral) inline() {}

type AbstractParentInline struct {
	Type    Type
	Inlines Inlines

	Location Location
}

type InlineSpan struct {
	Name    Name
	Variant Variant
	Form    Form

	AbstractParentInline
}

type InlineRef struct {
	Name    Name
	Variant Variant
	Target  string

	AbstractParentInline
}

type InlineLiteral struct {
	Name  Name
	Type  Type
	Value string

	Location Location
}

type Author struct {
	FullName   string
	Initials   string
	FirstName  string
	MiddleName string
	LastName   string
	Address    string
}

type Location []LocationBoundary

type LocationBoundary struct {
	Line    int `json:"line"`
	Collumn int `json:"col"`
}
