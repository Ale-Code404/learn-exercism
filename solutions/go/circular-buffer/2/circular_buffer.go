package circularbuffer

import (
	"errors"
)

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.
type Buffer struct {
	data     []byte
	capacity int
	size     int
	write    int
	read     int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		size:     0,
		capacity: size,
		data:     make([]byte, size),
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.IsEmpty() {
		return 0, errors.New("The buffer is empty")
	}

	value := b.data[b.read]
	b.read = (b.read + 1) % b.capacity
	b.size--

	return value, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.IsFull() {
		return errors.New("The buffer is full")
	}

	b.data[b.write] = c
	b.write = (b.write + 1) % b.capacity
	b.size++

	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if !b.IsFull() {
		b.WriteByte(c)
		return
	}

	b.data[b.read] = c
	b.write = b.read
	b.read = (b.read + 1) % b.capacity
}

func (b *Buffer) IsEmpty() bool {
	return b.size == 0
}

func (b *Buffer) IsFull() bool {
	return b.size == b.capacity
}

func (b *Buffer) Reset() {
	b.size = 0
	b.read = 0
	b.write = 0
}
