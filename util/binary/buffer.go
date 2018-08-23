package binary

import (
	"io"

	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/juju/errors"
)

var errTooSmall = errors.Errorf("binary:Payload is too small!")
var errEOF = errors.Errorf("binary:EOF")
var errNilBuffer = errors.Errorf("binary:not supprot a nil buffer")

const NotLengthEncodeInteger int = -1

//assert is used to protect the program
func assert(b []byte, size int) error {
	if b == nil {
		return errTooSmall
	} else if len(b) < size {
		return errTooSmall
	} else {
		return nil
	}
}

//ReadInt1 reads 1 bytes
func ReadInt1(b []byte) (uint8, error) {
	if err := assert(b, 1); err != nil {
		return 0, errors.Trace(err)
	} else {
		return uint8(b[0]), nil
	}
}

//ReadInt2 reads 2 bytes
func ReadInt2(buffer []byte) (uint16, error) {
	if err := assert(buffer, 2); err != nil {
		return 0, errors.Trace(err)
	} else {
		b := uint16(buffer[0])
		var i uint8 = 1
		for i < 2 {
			b = b | uint16(buffer[i])<<(i*8)
			i++
		}
		return b, nil
	}
}

//ReadInt3 reads 3 bytes
func ReadInt3(buffer []byte) (uint32, error) {
	if err := assert(buffer, 3); err != nil {
		return 0, errors.Trace(err)
	} else {
		b := uint32(buffer[0])
		var i uint8 = 1
		for i < 3 {
			b = b | uint32(buffer[i])<<(i*8)
			i++
		}
		return b, nil
	}
}

//ReadInt4 reads 4 bytes
func ReadInt4(buffer []byte) (uint32, error) {
	if err := assert(buffer, 4); err != nil {
		return 0, errors.Trace(err)
	} else {
		b := uint32(buffer[0])
		var i uint8 = 1
		for i < 4 {
			b = b | uint32(buffer[i])<<(i*8)
			i++
		}
		return b, nil
	}
}

//ReadInt6 reads 6 bytes
func ReadInt6(buffer []byte) (uint64, error) {
	if err := assert(buffer, 6); err != nil {
		return 0, errors.Trace(err)
	} else {
		b := uint64(buffer[0])
		var i uint8 = 1
		for i < 6 {
			b = b | uint64(buffer[i])<<(i*8)
			i++
		}
		return b, nil
	}
}

//ReadInt8 reads 8 bytes
func ReadInt8(buffer []byte) (uint64, error) {
	if err := assert(buffer, 8); err != nil {
		return 0, errors.Trace(err)
	} else {
		b := uint64(buffer[0])
		var i uint8 = 1
		for i < 8 {
			b = b | uint64(buffer[i])<<(i*8)
			i++
		}
		return b, nil
	}
}

//WriteInt1 writes 1 byte into buffer
func WriteInt1(b []byte, val uint8) error {
	if err := assert(b, 1); err != nil {
		return errors.Trace(err)
	} else {
		b[0] = val
		return nil
	}
}

//WriteInt2 writes 2 bytes into buffer
func WriteInt2(b []byte, val uint16) error {
	if err := assert(b, 2); err != nil {
		return errors.Trace(err)
	} else {
		b[0] = uint8(val)
		b[1] = uint8(val >> 8)
		return nil
	}
}

//WriteInt3 writes 3 bytes into buffer
func WriteInt3(b []byte, val uint32) error {
	if err := assert(b, 3); err != nil {
		return errors.Trace(err)
	} else {
		b[0] = uint8(val)
		b[1] = uint8(val >> 8)
		b[2] = uint8(val >> 16)
		return nil
	}
}

//WriteInt4 writes 4 bytes into buffer
func WriteInt4(b []byte, val uint32) error {
	if err := assert(b, 4); err != nil {
		return errors.Trace(err)
	} else {
		b[0] = uint8(val)
		b[1] = uint8(val >> 8)
		b[2] = uint8(val >> 16)
		b[3] = uint8(val >> 24)
		return nil
	}
}

//WriteInt6 writes 6 bytes into buffer
func WriteInt6(b []byte, val uint64) error {
	if err := assert(b, 6); err != nil {
		return errors.Trace(err)
	} else {
		b[0] = uint8(val)
		b[1] = uint8(val >> 8)
		b[2] = uint8(val >> 16)
		b[3] = uint8(val >> 24)
		b[4] = uint8(val >> 32)
		b[5] = uint8(val >> 40)
		return nil
	}
}

//WriteInt8 writes 8 bytes into buffer
func WriteInt8(b []byte, val uint64) error {
	if err := assert(b, 8); err != nil {
		return errors.Trace(err)
	} else {
		b[0] = uint8(val)
		b[1] = uint8(val >> 8)
		b[2] = uint8(val >> 16)
		b[3] = uint8(val >> 24)
		b[4] = uint8(val >> 32)
		b[5] = uint8(val >> 40)
		b[6] = uint8(val >> 48)
		b[7] = uint8(val >> 56)
		return nil
	}
}

//StoreInt1 Store 1 byte into buffer
func StoreInt1(val uint8) []byte {
	b := make([]byte, 1)
	b[0] = val
	return b
}

//StoreInt2 Store 2 bytes into buffer
func StoreInt2(val uint16) []byte {
	b := make([]byte, 2)
	b[0] = uint8(val)
	b[1] = uint8(val >> 8)
	return b
}

//StoreInt3 Store 3 bytes into buffer
func StoreInt3(val uint32) []byte {
	b := make([]byte, 3)
	b[0] = uint8(val)
	b[1] = uint8(val >> 8)
	b[2] = uint8(val >> 16)
	return b
}

//StoreInt4 Store 4 bytes into buffer
func StoreInt4(val uint32) []byte {
	b := make([]byte, 4)
	b[0] = uint8(val)
	b[1] = uint8(val >> 8)
	b[2] = uint8(val >> 16)
	b[3] = uint8(val >> 24)
	return b
}

//StoreInt6 Store 6 bytes into buffer
func StoreInt6(val uint64) []byte {
	b := make([]byte, 6)
	b[0] = uint8(val)
	b[1] = uint8(val >> 8)
	b[2] = uint8(val >> 16)
	b[3] = uint8(val >> 24)
	b[4] = uint8(val >> 32)
	b[5] = uint8(val >> 40)
	return b
}

//StoreInt8 Store 8 bytes into buffer
func StoreInt8(val uint64) []byte {
	b := make([]byte, 8)
	b[0] = uint8(val)
	b[1] = uint8(val >> 8)
	b[2] = uint8(val >> 16)
	b[3] = uint8(val >> 24)
	b[4] = uint8(val >> 32)
	b[5] = uint8(val >> 40)
	b[6] = uint8(val >> 48)
	b[7] = uint8(val >> 56)
	return b
}

//LengthOfInteger return the length of a length encode integer
func LengthOfInteger(val uint64) int {
	switch {
	case val < 0xfb:
		return 1
	case val <= 0xffff:
		return 3
	case val <= 0xffffff:
		return 4
	case val <= 0xffffffffffffffff:
		return 9
	default:
		return NotLengthEncodeInteger
	}
}

func ReadLengthEncodedInteger(buffer []byte) (uint64, int, error) {
	if err := assert(buffer, 1); err != nil {
		return 0, NotLengthEncodeInteger, errors.Trace(err)
	} else {
		i := buffer[0]
		switch {
		case i < 0xfb:
			val, err := ReadInt1(buffer)
			return uint64(val), 1, err
		case i == 0xfc:
			val, err := ReadInt2(buffer[1:])
			return uint64(val), 3, err
		case i == 0xfd:
			val, err := ReadInt3(buffer[1:])
			return uint64(val), 4, err
		case i == 0xfe:
			val, err := ReadInt8(buffer[1:])
			return val, 9, err
		default:
			return 0, NotLengthEncodeInteger, errors.Errorf("binary:not a LengthEncodedInteger")
		}
	}
}

//WriteLengthEncodedInteger write undetermined length integer
func WriteLengthEncodedInteger(b []byte, val uint64) (int, error) {
	if err := assert(b, 1); err != nil {
		return -1, errors.Trace(err)
	} else {
		switch {
		case val < 0xfb:
			return 1, WriteInt1(b, uint8(val))
		case val <= 0xffff:
			_ = WriteInt1(b, 0xfc)
			return 3, WriteInt2(b, uint16(val))
		case val < 0xffffff:
			_ = WriteInt1(b, 0xfd)
			return 4, WriteInt3(b, uint32(val))
		case val <= 0xffffffffffffffff:
			_ = WriteInt1(b, 0xfe)
			return 9, WriteInt8(b, val)
		default:
			return -1, errors.Errorf("binary:not a LengthEncodedInteger")
		}
	}
}

//ReadStringWithFixLen write a fix length string into buffer
func ReadStringWithFixLen(b []byte, l int) ([]byte, error) {
	if err := assert(b, l); err != nil {
		return nil, errors.Trace(err)
	} else {
		return b[:l], nil
	}
}

//WriteStringWithFixLen write a fix length string into buffer
func WriteStringWithFixLen(b []byte, str []byte) error {
	if err := assert(b, len(str)); err != nil {
		return errors.Trace(err)
	}

	if n := copy(b, str); n != len(str) {
		return errors.Errorf("binary:fail to copy data to buffer")
	} else {
		return nil
	}
}

//calculate the encode string length
func LengthOfString(b []byte) int {
	return LengthOfInteger(uint64(len(b))) + len(b)
}

//WriteLengthEncodedString write a string  with fixed length
func ReadLengthEncodedString(b []byte) ([]byte, error) {
	if n, _, err := ReadLengthEncodedInteger(b); err != nil {
		return nil, errors.Trace(err)
	} else {
		return ReadStringWithFixLen(b[1:], int(n))
	}
}

//WriteLengthEncodedString write a string  with fixed length
func WriteLengthEncodedString(b []byte, str []byte) error {
	if _, err := WriteLengthEncodedInteger(b, uint64(len(str))); err != nil {
		return errors.Trace(err)
	}
	return WriteStringWithFixLen(b[1:], str)
}

//Buffer represent a buffer in a packet
type Buffer struct {
	b   []byte
	off int
}

func NewBuffer(b []byte) *Buffer {
	return &Buffer{
		b:   b,
		off: 0,
	}
}

//Set add a new buffer
func (p *Buffer) Set(b []byte) {
	p.b = b
	p.off = 0
}

//Read implement the io.Reader interface
func (p *Buffer) Read(b []byte) (int, error) {
	if len(p.b)-p.off < len(b) {
		return -1, errTooSmall
	}
	n := copy(b, p.b[p.off:])
	p.off += n
	return n, nil
}

//Read implement the io.Reader interface
func (p *Buffer) ToReader() io.Reader {
	p.off = 0
	return p
}

//Write implement the io.Write interface
func (p *Buffer) Write(b []byte) (int, error) {
	if len(p.b)-p.off < len(b) {
		return -1, errTooSmall
	}

	return copy(p.b[p.off:], b), nil
}

//IsEOf return
func (p *Buffer) IsEOF() bool {
	return p.off >= len(p.b)
}

func (p *Buffer) Bytes() []byte {
	return p.b
}

//IsEOf return
func (p *Buffer) Len() int {
	return len(p.b)
}

func (p *Buffer) Off() int {
	return p.off
}

//ReadInt1 reads 1 bytes
func (p *Buffer) ReadInt1() (uint8, error) {
	buffer := p.b[p.off:]
	p.off += 1
	return ReadInt1(buffer)
}

//ReadInt2 reads 2 bytes
func (p *Buffer) ReadInt2() (uint16, error) {
	buffer := p.b[p.off:]
	p.off += 2
	return ReadInt2(buffer)
}

//ReadInt3 reads 3 bytes
func (p *Buffer) ReadInt3() (uint32, error) {
	buffer := p.b[p.off:]
	p.off += 3
	return ReadInt3(buffer)
}

//ReadInt4 reads 4 bytes
func (p *Buffer) ReadInt4() (uint32, error) {
	buffer := p.b[p.off:]
	p.off += 4
	return ReadInt4(buffer)
}

//ReadInt6 reads 6 bytes
func (p *Buffer) ReadInt6() (uint64, error) {
	buffer := p.b[p.off:]
	p.off += 6
	return ReadInt6(buffer)
}

//ReadInt8 reads 8 bytes
func (p *Buffer) ReadInt8() (uint64, error) {
	buffer := p.b[p.off:]
	p.off += 8
	return ReadInt8(buffer)
}

//Skip
func (p *Buffer) Skip(l int) error {
	if p.off+l > len(p.b) {
		return errTooSmall
	}
	p.off += l
	return nil
}

//ReadLengthEncodedInteger read undetermined length integer
func (p *Buffer) ReadLengthEncodedInteger() (uint64, int, error) {
	if p.IsEOF() {
		return 0, NotLengthEncodeInteger, errEOF
	}

	i := p.b[p.off]
	switch {
	case i < 0xfb:
		val, err := p.ReadInt1()
		return uint64(val), 1, err
	case i == 0xfc:
		p.off++
		val, err := p.ReadInt2()
		return uint64(val), 3, err
	case i == 0xfd:
		p.off++
		val, err := p.ReadInt3()
		return uint64(val), 4, err
	case i == 0xfe:
		p.off++
		val, err := p.ReadInt8()
		return val, 9, err
	}

	return 0, NotLengthEncodeInteger, errors.Errorf("binary:not a LengthEncodedInteger")
}

//ReadStringWithNull reads a string terminated with null(0x00)
func (p *Buffer) ReadStringWithNull() ([]byte, error) {
	if p.IsEOF() {
		return nil, errEOF
	}
	for off := p.off; off < len(p.b); off++ {
		if p.b[off] == 0x00 {
			//return not include 0x00
			b := p.b[p.off:off]
			p.off = off + 1
			return b, nil
		}
	}
	//return all bytes
	return p.b[p.off:], nil
}

//ReadStringWithEOF reads a string util EOF
func (p *Buffer) ReadStringWithEOF() ([]byte, error) {
	if p.IsEOF() {
		return nil, errEOF
	}
	return p.b[p.off:], nil
}

//ReadStringWithFixLen reads a string  with fixed length
func (p *Buffer) ReadStringWithFixLen(l int) ([]byte, error) {
	if p.off+l > len(p.b) {
		return nil, errTooSmall
	}
	b := p.b[p.off : p.off+l]
	p.off += l
	return b, nil
}

//ReadLengthEncodedString reads a string  with fixed length
func (p *Buffer) ReadLengthEncodedString() ([]byte, error) {
	if n, _, err := p.ReadLengthEncodedInteger(); err != nil {
		return nil, errors.Trace(err)
	} else {
		return p.ReadStringWithFixLen(int(n))
	}
}

//WriteInt1 write 1 byte into buffer
func (p *Buffer) WriteInt1(val uint8) error {
	buffer := p.b[p.off:]
	p.off += 1
	return WriteInt1(buffer, val)
}

func (p *Buffer) WriteInt2(val uint16) error {
	buffer := p.b[p.off:]
	p.off += 2
	return WriteInt2(buffer, val)
}

func (p *Buffer) WriteInt3(val uint32) error {
	buffer := p.b[p.off:]
	p.off += 3
	return WriteInt3(buffer, val)
}

func (p *Buffer) WriteInt4(val uint32) error {
	buffer := p.b[p.off:]
	p.off += 4
	return WriteInt4(buffer, val)
}

func (p *Buffer) WriteInt6(val uint64) error {
	buffer := p.b[p.off:]
	p.off += 6
	return WriteInt6(buffer, val)
}

func (p *Buffer) WriteInt8(val uint64) error {
	buffer := p.b[p.off:]
	p.off += 8
	return WriteInt8(buffer, val)
}

func (p *Buffer) WriteStringWithNull(str []byte) error {
	if n := copy(p.b[p.off:], str); n != len(str) {
		return errTooSmall
	} else {
		p.off += len(str)
	}
	//write '\0'
	return p.WriteInt1(0x00)
}

func (p *Buffer) WriteStringWithFixLen(str []byte) error {
	if n := copy(p.b[p.off:], str); n != len(str) {
		return errTooSmall
	} else {
		p.off += len(str)
		return nil
	}
}

//WriteLengthEncodedInteger write undetermined length integer
func (p *Buffer) WriteLengthEncodedInteger(val uint64) (int, error) {
	switch {
	case val < 0xfb:
		err := p.WriteInt1(uint8(val))
		return 1, err
	case val <= 0xffff:
		_ = p.WriteInt1(0xfc)
		err := p.WriteInt2(uint16(val))
		return 3, err
	case val < 0xffffff:
		_ = p.WriteInt1(0xfd)
		err := p.WriteInt3(uint32(val))
		return 4, err
	case val <= 0xffffffffffffffff:
		_ = p.WriteInt1(0xfe)
		err := p.WriteInt8(val)
		return 9, err
	default:
		panic("Invaid LengthEncodedInteger flag!")
	}
	return -1, nil
}

//WriteLengthEncodedString write a string  with fixed length
func (p *Buffer) WriteLengthEncodedString(str []byte) error {
	if _, err := p.WriteLengthEncodedInteger(uint64(len(str))); err != nil {
		return errors.Trace(err)
	}
	return p.WriteStringWithFixLen(str)
}

//ReadSlice return the slice when meets delim
func (p *Buffer) ReadSlice(delim byte) ([]byte, error) {
	if p.IsEOF() {
		return nil, io.EOF
	}

	for pos := p.off; pos < len(p.b); pos++ {
		if p.b[pos] == delim {
			b := p.b[p.off:pos]
			p.off = pos + 1
			return b, nil
		}
	}

	b := p.b[p.off:]
	p.off = len(p.b)
	return b, nil
}

//ReadString return the string when meets delim
func (p *Buffer) ReadString(delim byte) (string, error) {
	if b, err := p.ReadSlice(delim); err != nil {
		return "", err
	} else {
		return hack.String(b), nil
	}
}
