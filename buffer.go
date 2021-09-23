package lex_analyser

import "errors"

const BufferMaxSize int = 100

type Buffer struct {
	buf     []rune
	pBuff   int
	bufSize int
}

var (
	ErrorEmtpyBuffer    error = errors.New("error, buffer is empty")
	ErrorBufferOverflow error = errors.New("error, buffer overflow")
)

func NewBuffer(size int) *Buffer {
	b := &Buffer{bufSize: size, pBuff: -1}
	b.buf = make([]rune, b.bufSize)
	return b
}

func (b *Buffer) getRune() (rune, error) {
	if b.pBuff < 0 {
		return '0', ErrorEmtpyBuffer
	}
	c := b.buf[b.pBuff]
	b.pBuff--
	return c, nil
}

func (b *Buffer) putRune(c rune) error {
	if b.pBuff == b.bufSize {
		return ErrorBufferOverflow
	}
	b.pBuff++
	b.buf[b.pBuff] = c
	return nil
}
