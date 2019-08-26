package main

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"unicode"

	"github.com/MDGSF/utils/log"
)

const (
	StateInit         = 1
	StateSectionEnter = 2
	StateComment      = 3
	StateKeyValueLine = 4
	StateValue        = 5
)

const (
	TokenTypeSection = iota
	TokenTypeComment
	TokenTypeKey
	TokenTypeAssign
	TokenTypeValue
)

type TokenHeader struct {
	TokenType int
	data      []byte
}

func NewToken(tokentype int, data []byte) *TokenHeader {
	result := &TokenHeader{}
	result.TokenType = tokentype
	result.data = data
	return result
}

type TokenSection struct {
	TokenHeader
}

func NewTokenSection(data []byte) *TokenSection {
	result := &TokenSection{}
	result.TokenType = TokenTypeSection
	result.data = data
	return result
}

type TokenComment struct {
	TokenHeader
}

func NewTokenComment(data []byte) *TokenComment {
	result := &TokenComment{}
	result.TokenType = TokenTypeComment
	result.data = data
	return result
}

type TProcessor struct {
	data     []byte
	size     int
	idx      int
	row      int
	column   int
	lines    []int
	state    int
	tokenOut chan interface{}
}

func NewProcessor(data []byte) *TProcessor {
	p := &TProcessor{
		data:     data,
		size:     len(data),
		idx:      0,
		row:      0,
		column:   0,
		lines:    make([]int, 0),
		state:    StateInit,
		tokenOut: make(chan interface{}, 16),
	}
	p.lines = append(p.lines, 0)
	return p
}

func (p *TProcessor) ReadByte() (byte, error) {
	var b byte
	if p.idx >= p.size {
		return b, errors.New("eof")
	}
	b = p.data[p.idx]
	if b == '\n' {
		p.row++
		p.column = 0
	} else {
		p.column++
	}
	p.idx++
	p.lines = append(p.lines, p.idx)
	return b, nil
}

func (p *TProcessor) UnreadByte() error {
	if p.idx <= 0 {
		return errors.New("start of file")
	}
	p.idx--
	return nil
}

func (p *TProcessor) ReadSlice(delim byte) ([]byte, error) {
	for {
		b, err := p.ReadByte()
		if err != nil {
			return err
		}
	}
}

func (p *TProcessor) run() {
	go p.generateTokenStream()
	p.processToken()
}

func (p *TProcessor) processToken() {
	for {
		select {
		case token, ok := <-p.tokenOut:
			if !ok {
				return
			}

			if section, ok := token.(*TokenSection); ok {
				log.Info("token section = %v", string(section.data))
				continue
			}

			if comment, ok := token.(*TokenComment); ok {
				log.Info("token comment = %v", string(comment.data))
				continue
			}

			if tokenmsg, ok := token.(*TokenHeader); ok {
				log.Info("token type = %v, len = %v, data = %v", tokenmsg.TokenType, len(tokenmsg.data), string(tokenmsg.data))
				continue
			}

		}
	}
}

func (p *TProcessor) generateTokenStream() {

	var err error

	rb := bytes.NewReader(data)
	p.r = bufio.NewReader(rb)

	for {
		switch p.state {
		case StateInit:
			err = p.onStateInit()
		case StateSectionEnter:
			err = p.onStateSectionEnter()
		case StateComment:
			err = p.onStateComment()
		case StateKeyValueLine:
			err = p.onStateKeyValueLine()
		case StateValue:
			err = p.onStateValue()
		}

		if err != nil {
			break
		}
	}

	close(p.tokenOut)
}

func (p *TProcessor) onStateInit() error {
	var err error
	var b byte

	for {
		b, err = p.ReadByte()
		if err != nil {
			break
		}

		//log.Info("current byte = [%v, %c]", b, b)

		if unicode.IsSpace(rune(b)) {
			continue
		}

		if b == '[' {
			p.state = StateSectionEnter
			return nil
		}

		if b == '"' || b == '#' || b == ';' {
			p.state = StateComment
			return nil
		}

		if unicode.IsLetter(rune(b)) {
			p.UnreadByte()
			p.state = StateKeyValueLine
			return nil
		}

	}

	return err
}

func (p *TProcessor) onStateSectionEnter() error {
	data, err := p.r.ReadSlice(']')
	if err != nil {
		return err
	}

	p.tokenOut <- NewTokenSection(data[:len(data)-1])

	p.state = StateInit
	return nil
}

func (p *TProcessor) onStateComment() error {
	data, err := p.r.ReadSlice('\n')
	if err != nil {
		return err
	}

	p.tokenOut <- NewTokenComment(data[:len(data)-1])

	p.state = StateInit
	return nil
}

func (p *TProcessor) onStateKeyValueLine() error {
	data, err := p.r.ReadSlice('=')
	if err != nil {
		return err
	}

	p.tokenOut <- NewToken(TokenTypeKey, bytes.TrimSpace(data[:len(data)-1]))
	p.tokenOut <- NewToken(TokenTypeAssign, []byte{'='})

	p.state = StateValue
	return nil
}

func (p *TProcessor) onStateValue() error {
	var err error
	var b byte

	for {
		b, err = p.ReadByte()
		if err != nil {
			break
		}

		if unicode.IsSpace(rune(b)) {
			continue
		}

		if b == '`' {
			data, err := p.r.ReadSlice('`')
			if err != nil {
				return err
			}

			p.tokenOut <- NewToken(TokenTypeValue, bytes.TrimSpace(data[:len(data)-1]))

			p.state = StateInit
			return nil
		}

		if b == '"' {
			data, err := p.r.ReadSlice('"')
			if err != nil {
				return err
			}

			p.tokenOut <- NewToken(TokenTypeValue, bytes.TrimSpace(data[:len(data)-1]))

			p.state = StateInit

			return nil
		}

		if unicode.IsLetter(rune(b)) {
			p.UnreadByte()
			data, err := p.r.ReadSlice('\n')
			if err != nil {
				return err
			}

			p.tokenOut <- NewToken(TokenTypeValue, bytes.TrimSpace(data[:len(data)-1]))

			p.state = StateInit
			return nil
		}
	}

	return err
}

func main() {
	log.Info("main start")
	defer log.Info("main end")

	data, err := ioutil.ReadFile("test.ini")
	if err != nil {
		panic(err)
	}

	p := NewProcessor(data)
	p.run()
}
