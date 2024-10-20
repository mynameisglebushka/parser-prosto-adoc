package ast

func NewDocument() *Document {
	return &Document{
		Type:   BlockType,
		Name:   DocumentName,
		Header: &Header{},
		Attributes: map[string]string{},
	}
}
