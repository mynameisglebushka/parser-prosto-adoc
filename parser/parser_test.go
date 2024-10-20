package parser

import (
	"reflect"
	"strings"
	"testing"

	"github.com/mynameisglebushka/parser-prosto-adoc/ast"
)

func TestParseHeader(t *testing.T) {

	const (
		TitleOnly             = "Title only"
		TitleComment          = "Title and Comment"
		TitleMultilineComment = "Title and Multiline Comment"
		TitleAuthor           = "Title and Author"
		TitleAuthors          = "Title and Authors"
		TitleAttribute        = "Title and Attribute"
		TitleAttributes       = "Title and Attributes"
	)

	cases := map[string][]string{
		TitleOnly: {
			"= Document Title",
		},
		TitleComment: {
			"// Coomment",
			"\n",
			"= Document Title",
		},
		TitleMultilineComment: {
			"//// ",
			"\n",
			"Multi Line Coomment",
			"\n",
			"////",
			"\n",
			"= Document Title",
		},
		TitleAuthor: {
			"= Document Title",
			"\n",
			"Gleb S. Glazkov <https://github.com/mynameisglebushka>",
		},
		TitleAuthors: {
			"= Document Title",
			"\n",
			"Gleb S. Glazkov <https://github.com/mynameisglebushka>; Glebushka Sergeevich Glazkov <mail@google.com>",
		},
		TitleAttribute: {
			"= Document Title",
			"\n",
			":nickname: mynameisglebushka",
		},
		TitleAttributes: {
			"= Document Title",
			"\n",
			":nickname: mynameisglebushka",
			"\n",
			":bool-attr:",
		},
	}

	tests := []struct {
		name  string
		input []byte
		want  *ast.Document
	}{
		{
			name: TitleOnly,
			input: []byte(
				strings.Join(cases[TitleOnly], ""),
			),
			want: &ast.Document{
				Type: ast.BlockType,
				Name: ast.DocumentName,
				Header: &ast.Header{
					Title: []ast.Inline{
						&ast.InlineLiteral{
							Name:  ast.TextName,
							Type:  ast.StringType,
							Value: "Document Title",
							Location: []ast.LocationBoundary{
								{
									Line:    1,
									Collumn: 3,
								},
								{
									Line:    1,
									Collumn: 16,
								},
							},
						},
					},
				},
				Attributes: map[string]string{},
				Location: []ast.LocationBoundary{
					{
						Line:    1,
						Collumn: 1,
					},
					{
						Line:    1,
						Collumn: 16,
					},
				},
			},
		},
		{
			name: TitleComment,
			input: []byte(
				strings.Join(cases[TitleComment], ""),
			),
			want: &ast.Document{
				Type: ast.BlockType,
				Name: ast.DocumentName,
				Header: &ast.Header{
					Title: []ast.Inline{
						&ast.InlineLiteral{
							Name:  ast.TextName,
							Type:  ast.StringType,
							Value: "Document Title",
							Location: []ast.LocationBoundary{
								{
									Line:    2,
									Collumn: 3,
								},
								{
									Line:    2,
									Collumn: 16,
								},
							},
						},
					},
				},
				Attributes: map[string]string{},
				Location: []ast.LocationBoundary{
					{
						Line:    1,
						Collumn: 1,
					},
					{
						Line:    2,
						Collumn: 16,
					},
				},
			},
		},
		{
			name: TitleMultilineComment,
			input: []byte(
				strings.Join(cases[TitleMultilineComment], ""),
			),
			want: &ast.Document{
				Type: ast.BlockType,
				Name: ast.DocumentName,
				Header: &ast.Header{
					Title: []ast.Inline{
						&ast.InlineLiteral{
							Name:  ast.TextName,
							Type:  ast.StringType,
							Value: "Document Title",
							Location: []ast.LocationBoundary{
								{
									Line:    4,
									Collumn: 3,
								},
								{
									Line:    4,
									Collumn: 16,
								},
							},
						},
					},
				},
				Attributes: map[string]string{},
				Location: []ast.LocationBoundary{
					{
						Line:    1,
						Collumn: 1,
					},
					{
						Line:    4,
						Collumn: 16,
					},
				},
			},
		},
		{
			name: TitleAuthor,
			input: []byte(
				strings.Join(cases[TitleAuthor], ""),
			),
			want: &ast.Document{
				Type: ast.BlockType,
				Name: ast.DocumentName,
				Header: &ast.Header{
					Title: []ast.Inline{
						&ast.InlineLiteral{
							Name:  ast.TextName,
							Type:  ast.StringType,
							Value: "Document Title",
							Location: []ast.LocationBoundary{
								{
									Line:    1,
									Collumn: 3,
								},
								{
									Line:    1,
									Collumn: 16,
								},
							},
						},
					},
					Authors: []ast.Author{
						{
							FullName:   "Gleb S. Glazkov",
							Initials:   "GSG",
							FirstName:  "Gleb",
							MiddleName: "S.",
							LastName:   "Glazkov",
							Address:    "https://github.com/mynameisglebushka",
						},
					},
				},
				Attributes: map[string]string{},
				Location: []ast.LocationBoundary{
					{
						Line:    1,
						Collumn: 1,
					},
					{
						Line:    2,
						Collumn: 54,
					},
				},
			},
		},
		{
			name: TitleAuthors,
			input: []byte(
				strings.Join(cases[TitleAuthors], ""),
			),
			want: &ast.Document{
				Type: ast.BlockType,
				Name: ast.DocumentName,
				Header: &ast.Header{
					Title: []ast.Inline{
						&ast.InlineLiteral{
							Name:  ast.TextName,
							Type:  ast.StringType,
							Value: "Document Title",
							Location: []ast.LocationBoundary{
								{
									Line:    1,
									Collumn: 3,
								},
								{
									Line:    1,
									Collumn: 16,
								},
							},
						},
					},
					Authors: []ast.Author{
						{
							FullName:   "Gleb S. Glazkov",
							Initials:   "GSG",
							FirstName:  "Gleb",
							MiddleName: "S.",
							LastName:   "Glazkov",
							Address:    "https://github.com/mynameisglebushka",
						},
						{
							FullName:   "Glebushka Sergeevich Glazkov",
							Initials:   "GSG",
							FirstName:  "Glebushka",
							MiddleName: "Sergeevich",
							LastName:   "Glazkov",
							Address:    "mail@google.com",
						},
					},
				},
				Attributes: map[string]string{},
				Location: []ast.LocationBoundary{
					{
						Line:    1,
						Collumn: 1,
					},
					{
						Line:    2,
						Collumn: 102,
					},
				},
			},
		},
		{
			name: TitleAttribute,
			input: []byte(
				strings.Join(cases[TitleAttribute], ""),
			),
			want: &ast.Document{
				Type: ast.BlockType,
				Name: ast.DocumentName,
				Header: &ast.Header{
					Title: []ast.Inline{
						&ast.InlineLiteral{
							Name:  ast.TextName,
							Type:  ast.StringType,
							Value: "Document Title",
							Location: []ast.LocationBoundary{
								{
									Line:    1,
									Collumn: 3,
								},
								{
									Line:    1,
									Collumn: 16,
								},
							},
						},
					},
				},
				Attributes: map[string]string{
					"nickname": "mynameisglebushka",
				},
				Location: []ast.LocationBoundary{
					{
						Line:    1,
						Collumn: 1,
					},
					{
						Line:    2,
						Collumn: 28,
					},
				},
			},
		},
		{
			name: TitleAttributes,
			input: []byte(
				strings.Join(cases[TitleAttributes], ""),
			),
			want: &ast.Document{
				Type: ast.BlockType,
				Name: ast.DocumentName,
				Header: &ast.Header{
					Title: []ast.Inline{
						&ast.InlineLiteral{
							Name:  ast.TextName,
							Type:  ast.StringType,
							Value: "Document Title",
							Location: []ast.LocationBoundary{
								{
									Line:    1,
									Collumn: 3,
								},
								{
									Line:    1,
									Collumn: 16,
								},
							},
						},
					},
				},
				Attributes: map[string]string{
					"nickname":  "mynameisglebushka",
					"bool-attr": "",
				},

				Location: []ast.LocationBoundary{
					{
						Line:    1,
						Collumn: 1,
					},
					{
						Line:    3,
						Collumn: 11,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newParser(tt.input)

			doc := p.parseDocument()

			if !reflect.DeepEqual(doc, tt.want) {
				t.Errorf("parseHeader() = %v, want %v", doc, tt.want)
			}

		})
	}
}

func TestSome(t *testing.T) {
	a := "Doc Writer <doc.writer@example.org>; J. R. Wordsmith <https://example.org/jr>"
	authors := strings.Split(a, ";")
	for _, author := range authors {
		author = strings.TrimLeft(author, `\t\n\v\f\r \x85\xA0`)

		params := strings.SplitN(author, " ", 4)

		for i, p := range params {
			t.Logf("%v %v\n", i, p)
		}

	}
	b := []string{"a"}
	t.Log(len(b[:len(b)-1]))
}
