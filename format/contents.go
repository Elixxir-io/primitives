////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

package format

const (
	// Length, start index, and end index of the Contents serial
	contentsLen   = 399 // 3192 bits
	contentsStart = 0
	contentsEnd   = contentsStart + contentsLen

	// The smallest length that the padding can be
	padMinLen = 11 // 88 bits

	// The initial value of position; indicates content start is unknown
	invalidPosition = -1
)

// Structure for the content section of the message points to a subsection of
// the serialised Message structure. For the purpose of E2E, padding is added to
// the front of serial, with a minimum length of 11 bytes. The first byte of
// padding is always all zeros so that a byte does not need to be placed at the
// beginning of serial to ensure it is in the group.
type Contents struct {
	// Stores the data of Contents and points to region in master
	serial []byte

	// Starting index of data (excluding padding) in serial. If it is -1, then
	// it means no padding is specified and functions relying on the padding
	// will panic.
	position int
}

// NewContents creates a new Contents for a message and sets serial. If the new
// serial is not exactly the same length as serial, then it panics.
func NewContents(newSerial []byte) *Contents {
	if len(newSerial) != contentsLen {
		panic("new serial not the same size as Contents serial")
	}

	newContents := &Contents{
		serial:   make([]byte, contentsLen),
		position: invalidPosition,
	}

	newContents.serial = newSerial

	return newContents
}

// Get returns the complete serialised data of Content. The caller can read or
// write the data within this slice, but cannot change the slice header in the
// actual structure.
func (c *Contents) Get() []byte {
	return c.serial
}

// Set sets the entire serial content. The number of bytes copied is returned.
// If the specified byte array is not exactly the same size as serial, then it
// panics.
func (c *Contents) Set(newSerial []byte) int {
	if len(newSerial) != contentsLen {
		panic("new serial not the same size as Contents serial")
	}

	return copy(c.serial, newSerial)
}

// GetRightAligned returns the entire serial content, excluding the padding. If
// the position of the data is not specified (position < 0), then it panics. The
// caller can read or write the data within this slice, but cannot change the
// slice header in the actual structure.
func (c *Contents) GetRightAligned() []byte {
	if c.position == invalidPosition {
		panic("invalid padding when getting right-aligned data")
	}

	return c.serial[c.position:]
}

// SetRightAligned sets the entire serial content right-aligned. The number of
// bytes copied is returned. If the specified byte array is larger than serial,
// then it panics.
func (c *Contents) SetRightAligned(newSerial []byte) int {
	if len(newSerial) > contentsLen-padMinLen {
		panic("new serial is larger than Contents serial")
	}

	c.position = contentsLen - len(newSerial)
	return copy(c.serial[c.position:], newSerial)
}

// GetPosition returns the index of the start of actual data (not padding) in
// serial.
func (c *Contents) GetPosition() int {
	return c.position
}

// DeepCopy creates a copy of Contents.
func (c *Contents) DeepCopy() *Contents {
	newCopy := NewContents(make([]byte, contentsLen))
	copy(newCopy.serial[:], c.serial)

	return newCopy
}
