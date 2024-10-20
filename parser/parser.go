package parser

import (
	"bytes"
	"os"
	"strings"

	"github.com/mynameisglebushka/parser-prosto-adoc/ast"
)

func Parse(path string) (*ast.Document, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	p := newParser(content)

	document := p.parseDocument()

	return document, nil
}

type parser struct {
	lines    [][]byte
	lineNum  int
	prevKind Kind
	kind     Kind
}

func newParser(content []byte) *parser {
	p := &parser{}

	content = bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))
	p.lines = bytes.Split(content, []byte("\n"))

	var (
		wspaces = "\t\n\v\f\r \x85\xA0"

		line []byte
		x    int
	)
	for x, line = range p.lines {
		p.lines[x] = bytes.TrimRight(line, wspaces)
	}

	return p
}

func (p *parser) parseDocument() *ast.Document {

	doc := ast.NewDocument()

	doc.Location = append(doc.Location, ast.LocationBoundary{
		Line:    1,
		Collumn: 1,
	})

	p.parseHeader(doc)

	doc.Location = append(doc.Location, ast.LocationBoundary{
		Line:    len(p.lines),
		Collumn: len(p.lines[len(p.lines)-1]),
	})

	return doc
}

// Header contains:
//
// # Title
//
// Pattern: "= Document title"
//
// # Text line contains Authors
//
// Pattern: "firstname middlename lastname <email>; firstname middlename lastname <email>"
//
// # Document attributes
//
// Pattern: ":^[a-zA-Z0-9_][-a-zA-Z0-9_]*$: null | string"
func (p *parser) parseHeader(doc *ast.Document) {
	p.skipEmptyOrCommentLines()

	for {
		line := p.nextLine()

		if line == nil {
			break
		}

		switch p.kind {
		case kindDocumentTitle:
			clearTitle := bytes.TrimLeft(line.content[1:], " ")
			start := bytes.Index(line.content, clearTitle)
			title := &ast.InlineLiteral{
				Name:  ast.TextName,
				Type:  ast.StringType,
				Value: string(clearTitle),
				Location: []ast.LocationBoundary{
					{
						Line:    p.lineNum,
						Collumn: start + 1,
					},
					{
						Line:    p.lineNum,
						Collumn: start + len(clearTitle),
					},
				},
			}
			doc.Header.Title = append(doc.Header.Title, title)
		case kindText:
			authors := bytes.Split(line.content, []byte(";"))
			for _, author := range authors {
				var full, in, fn, mn, ln, addr string

				author = bytes.TrimLeft(author, "\t\n\v\f\r \x85\xA0")

				params := bytes.SplitN(author, []byte(" "), 4)

				if bytes.HasPrefix(params[len(params)-1], []byte("<")) && bytes.HasSuffix(params[len(params)-1], []byte(">")) {
					addr = string(bytes.Trim(params[len(params)-1], "<>"))
					params = params[:len(params)-1]
				}

				switch len(params) {
				case 1:
					fn = string(params[0])
					full = fn
					in = string(params[0][0])
				case 2:
					fn = string(params[0])
					ln = string(params[1])
					full = strings.Join([]string{fn, ln}, " ")
					in = string(params[0][0]) + string(params[1][0])
				case 3:
					fn = string(params[0])
					mn = string(params[1])
					ln = string(params[2])
					full = strings.Join([]string{fn, mn, ln}, " ")
					in = string(params[0][0]) + string(params[1][0]) + string(params[2][0])
				}

				doc.Header.Authors = append(doc.Header.Authors, ast.Author{
					FullName:   full,
					Initials:   in,
					FirstName:  fn,
					MiddleName: mn,
					LastName:   ln,
					Address:    addr,
				})
			}
		case lineKindAttribute:
			k, v := parseDocumentAttribute(line.content)
			doc.Attributes[k] = v
		}
	}
}

func (p *parser) skipEmptyOrCommentLines() {
	var (
		isMLC bool
	)

	for {
		line := p.nextLine()

		if line == nil {
			break
		}

		switch p.kind {
		case lineComment:
			continue
		case lineMultilineComment:
			isMLC = true
			continue
		case lineEmpty:
			continue
		}

		if isMLC {
			continue
		}

		break
	}

	p.lineNum--
}

// Not safe document attribute parser
//
// Put here only ":key: value" lines
func parseDocumentAttribute(line []byte) (key string, value string) {
	attr := bytes.SplitN(line, []byte(" "), 2)

	key = string(bytes.Trim(attr[0], ":"))

	if len(attr) == 2 {
		value = string(attr[1])
	}

	return key, value
}

func (p *parser) nextLine() *line {
	p.prevKind = p.kind

	if p.lineNum >= len(p.lines) {
		return nil
	}

	content := p.lines[p.lineNum]
	p.lineNum++

	var (
		x int
		r byte
	)
	for x, r = range content {
		if r == ' ' || r == '\t' {
			continue
		}
		break
	}

	line := &line{
		spases:  content[:x],
		content: content[x:],
	}

	p.kind = line.lineKindOf()

	return line
}
