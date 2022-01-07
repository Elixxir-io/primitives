////////////////////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                                       //
//                                                                                        //
// Use of this source code is governed by a license that can be found in the LICENSE file //
////////////////////////////////////////////////////////////////////////////////////////////

package format

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strconv"

	jww "github.com/spf13/jwalterweatherman"
)

const (
	KeyFPLen        = 32
	MacLen          = 32
	EphemeralRIDLen = 8
	IdentityFPLen   = 25
	RecipientIDLen  = EphemeralRIDLen + IdentityFPLen

	MinimumPrimeSize = 2*MacLen + RecipientIDLen // NOTE(david): why is this derived from the maclen and recipeintidlen?

	AssociatedDataSize = KeyFPLen + MacLen + RecipientIDLen

	messagePayloadVersion = 0
)

/*
                            Message Structure (not to scale)
+--------------------------------------------------------------------------------------------------------+
|                                               Message                                                  |
|                                          2*primeSize bits                                              |
+------------------------------------------+-------------------------------------------------------------+
|                 payloadA                 |                         payloadB                            |
|              primeSize bits              |                     primeSize bits                          |
+---------+----------+---------------------+---------+----------+------------+--------------+------------+
| grpBitA |  keyFP   |version| Contents1   | grpBitB | reserved | Contents2  | ephemeralRID | identityFP |
|  1 bit  | 255 bits |1 byte |  *below*    |  1 bit  | 7 bits   |  *below*   |   64 bits    |  200 bits  |
+ --------+----------+---------------------+---------+----------+------------+--------------+------------+
|                              Raw Contents                              |
|                    2*primeSize - recipientID bits                      |
+------------------------------------------------------------------------+

* size: size in bits of the data which is stored
* Contents1 size = primeSize - grpBitASize - KeyFPLen - sizeSize - 1
* AEADOverheadLen = 16 bytes
* Contents2 size = primeSize - grpBitBSize - AEADOverheadLen - RecipientIDLen - timestampSize
* the size of the data in the two contents fields is stored within the "size" field

/////Adherence to the group/////////////////////////////////////////////////////
The first bits of keyFingerprint and MAC are enforced to be 0, thus ensuring
PayloadA and PayloadB are within the group
*/

// Message structure stores all the data serially. Subsequent fields point to
// subsections of the serialised data.
type Message struct {
	data []byte

	// Note: These are mapped to locations in the data object
	payloadA []byte
	payloadB []byte

	keyFP        []byte
	version      []byte
	contents1    []byte
	reserved     []byte
	contents2    []byte
	ephemeralRID []byte // Ephemeral reception ID
	identityFP   []byte // Identity fingerprint

	rawContents []byte
}

// NewMessage creates a new empty message based upon the size of the encryption
// primes. All subcomponents point to locations in the internal data buffer.
// Panics if the prime size to too small.
func NewMessage(numPrimeBytes int) Message {
	if numPrimeBytes < MinimumPrimeSize {
		jww.FATAL.Panicf("Failed to create new Message: minimum prime length "+
			"is %d, received prime size is %d.", MinimumPrimeSize, numPrimeBytes)
	}

	data := make([]byte, 2*numPrimeBytes)

	return Message{
		data: data,

		payloadA: data[:numPrimeBytes],
		payloadB: data[numPrimeBytes:],

		keyFP:     data[:KeyFPLen],
		version:   data[KeyFPLen : KeyFPLen+1],
		contents1: data[1+KeyFPLen : numPrimeBytes],

		reserved:     data[numPrimeBytes : numPrimeBytes+ReservedLen],
		contents2:    data[numPrimeBytes+MacLen : 2*numPrimeBytes-RecipientIDLen],
		ephemeralRID: data[2*numPrimeBytes-RecipientIDLen : 2*numPrimeBytes-IdentityFPLen],
		identityFP:   data[2*numPrimeBytes-IdentityFPLen:],

		rawContents: data[:2*numPrimeBytes-RecipientIDLen],
	}
}

// Marshal marshals the message into a byte slice.
func (m *Message) Marshal() []byte {
	return copyByteSlice(m.data)
}

// Unmarshal unmarshalls a byte slice into a new Message.
func Unmarshal(b []byte) (Message, error) {
	m := NewMessage(len(b) / 2)
	copy(m.data, b)

	// if m.Version() != messagePayloadVersion {
	// 	return Message{}, fmt.Errorf(
	// 		"message encoding version mismatch, got %d expected %d",
	// 		m.Version(), messagePayloadVersion)
	// }

	return m, nil
}

// Version returns the encoding version.
func (m *Message) Version() uint8 {
	return m.version[0]
}

// Copy returns a copy of the message.
func (m Message) Copy() Message {
	m2 := NewMessage(len(m.data) / 2)
	copy(m2.data, m.data)
	return m2
}

// GetPrimeByteLen returns the size of the prime used.
func (m Message) GetPrimeByteLen() int {
	return len(m.data) / 2
}

// GetPayloadA returns payload A, which is the first half of the message.
func (m Message) GetPayloadA() []byte {
	return copyByteSlice(m.payloadA)
}

// SetPayloadA copies the passed byte slice into payload A. If the specified
// byte slice is not exactly the same size as payload A, then it panics.
func (m Message) SetPayloadA(payload []byte) {
	if len(payload) != len(m.payloadB) {
		jww.ERROR.Panicf("Failed to set Message payload A: length must be %d, "+
			"length of received data is %d.", len(m.payloadA), len(payload))
	}

	copy(m.payloadA, payload)
}

// GetPayloadB returns payload B, which is the last half of the message.
func (m Message) GetPayloadB() []byte {
	return copyByteSlice(m.payloadB)
}

// SetPayloadB copies the passed byte slice into payload B. If the specified
// byte slice is not exactly the same size as payload B, then it panics.
func (m Message) SetPayloadB(payload []byte) {
	if len(payload) != len(m.payloadB) {
		jww.ERROR.Panicf("Failed to set Message payload B: length must be %d, "+
			"length of received data is %d.", len(m.payloadB), len(payload))
	}

	copy(m.payloadB, payload)
}

// ContentsSize returns the maximum size of the contents.
func (m Message) ContentsSize() int {
	return len(m.data) - AssociatedDataSize - 1
}

// GetContents returns the exact contents of the message. This size of the
// return is based on the size of the contents actually stored.
func (m Message) GetContents() []byte {
	c := make([]byte, len(m.contents1)+len(m.contents2))

	copy(c[:len(m.contents1)], m.contents1)
	copy(c[len(m.contents1):], m.contents2)

	return c
}

// SetContents sets the contents of the message. This overwrites any storage
// already in the message but will not clear bits beyond the size of the passed
// contents. Panics if the passed contents is larger than the maximum contents
// size.
func (m Message) SetContents(c []byte) {
	if len(c) > len(m.contents1)+len(m.contents2) {
		jww.ERROR.Panicf("Failed to set Message contents: length must be "+
			"equal to or less than %d, length of received data is %d.",
			len(m.contents1)+len(m.contents2), len(c))
	}

	if len(c) <= len(m.contents1) {
		copy(m.contents1, c)
	} else {
		copy(m.contents1, c[:len(m.contents1)])
		copy(m.contents2, c[len(m.contents1):])
	}
}

// GetRawContentsSize returns the exact contents of the message.
func (m Message) GetRawContentsSize() int {
	return len(m.rawContents)
}

// GetRawContents returns the exact contents of the message. This field crosses
// over the group barrier and the setter of this is responsible for ensuring the
// underlying payloads are within the group.
// flips the first bit to 0 on return
func (m Message) GetRawContents() []byte {
	newRaw := copyByteSlice(m.rawContents)
	clearFirstBit(newRaw)
	newRaw[m.GetPrimeByteLen()] &= 0b01111111
	return newRaw
}

// SetRawContents sets the raw contents of the message. This field crosses over
// the group barrier and the setter of this is responsible for ensuring the
// underlying payloads are within the group. This will panic if the payload is
// greater than the maximum size. This overwrites any storage already in the
// message. If the passed contents is larger than the maximum contents size this
// will panic.
func (m Message) SetRawContents(c []byte) {
	if len(c) != len(m.rawContents) {
		jww.ERROR.Panicf("Failed to set Message raw contents: length must be %d, "+
			"length of received data is %d.", len(m.rawContents), len(c))
	}

	copy(m.rawContents, c)
}

// GetKeyFP gets the key Fingerprint
// flips the first bit to 0 on return
func (m Message) GetKeyFP() Fingerprint {
	newFP := NewFingerprint(m.keyFP)
	clearFirstBit(newFP[:])
	return newFP
}

// SetKeyFP sets the key Fingerprint. Checks that the first bit of the Key
// Fingerprint is 0, otherwise it panics.
func (m Message) SetKeyFP(fp Fingerprint) {
	if fp[0]>>7 != 0 {
		jww.ERROR.Panicf("Failed to set Message key fingerprint: first bit " +
			"of provided data must be zero.")
	}

	copy(m.keyFP, fp.Bytes())
}

// GetEphemeralRID returns the ephemeral recipient ID.
func (m Message) GetEphemeralRID() []byte {
	return copyByteSlice(m.ephemeralRID)
}

// SetEphemeralRID copies the ephemeral recipient ID bytes into the message.
func (m Message) SetEphemeralRID(ephemeralRID []byte) {
	if len(ephemeralRID) != EphemeralRIDLen {
		jww.ERROR.Panicf("Failed to set Message ephemeral recipient ID: "+
			"length must be %d, length of received data is %d.",
			EphemeralRIDLen, len(ephemeralRID))
	}
	copy(m.ephemeralRID, ephemeralRID)
}

// GetIdentityFP return the identity fingerprint.
func (m Message) GetIdentityFP() []byte {
	return copyByteSlice(m.identityFP)
}

// SetIdentityFP sets the identity fingerprint, which should be generated via
// fingerprint.IdentityFP.
func (m Message) SetIdentityFP(identityFP []byte) {
	if len(identityFP) != IdentityFPLen {
		jww.ERROR.Panicf("Failed to set Message identity fingerprint: length "+
			"must be %d, length of received data is %d.",
			IdentityFPLen, len(identityFP))
	}
	copy(m.identityFP, identityFP)
}

// gets a digest of the message contents, primarily used for debugging
func (m Message) Digest() string {
	h := md5.New()
	h.Write(m.GetContents())
	d := h.Sum(nil)
	digest := base64.StdEncoding.EncodeToString(d[:15])
	return digest[:20]
}

// copyByteSlice is a helper function to make a copy of a byte slice.
func copyByteSlice(s []byte) []byte {
	c := make([]byte, len(s))
	copy(c, s)
	return c
}

// GoString returns the Message key fingerprint, ephemeral recipient ID,
// identity fingerprint, and contents as a string. This functions satisfies the
// fmt.GoStringer interface.
func (m Message) GoString() string {
	keyFP := "<nil>"
	if len(m.keyFP) > 0 {
		keyFP = m.GetKeyFP().String()
	}
	ephID := "<nil>"
	if len(m.ephemeralRID) > 0 {
		ephID = strconv.FormatUint(binary.BigEndian.Uint64(m.GetEphemeralRID()), 10)
	}
	identityFP := "<nil>"
	if len(m.identityFP) > 0 {
		identityFP = base64.StdEncoding.EncodeToString(m.GetIdentityFP())
	}

	return "format.Message{" +
		"keyFP:" + keyFP +
		", ephemeralRID:" + ephID +
		", identityFP:" + identityFP +
		", contents:" + fmt.Sprintf("%q", m.GetContents()) + "}"
}

// SetGroupBits allows the first and second bits to be set in the payload.
// This should be used with code which determines if the bit can be set
// to 1 before proceeding.
func (m Message) SetGroupBits(bitA, bitB bool) {
	setFirstBit(m.payloadA, bitA)
	setFirstBit(m.payloadB, bitB)
}

func setFirstBit(b []byte, bit bool) {
	if bit {
		b[0] |= 0b10000000
	} else {
		b[0] &= 0b01111111
	}
}

func clearFirstBit(b []byte) {
	b[0] = 0b01111111 & b[0]
}
