////////////////////////////////////////////////////////////////////////////////
// Copyright © 2024 xx foundation                                             //
//                                                                            //
// Use of this source code is governed by a license that can be found in the  //
// LICENSE file.                                                              //
////////////////////////////////////////////////////////////////////////////////

package format

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"testing"
)

// Happy path.
func TestNewFingerprint(t *testing.T) {
	prng := rand.New(rand.NewSource(42))
	fpBytes := make([]byte, KeyFPLen)
	prng.Read(fpBytes)

	fp := NewFingerprint(fpBytes)
	if !bytes.Equal(fpBytes, fp[:]) {
		t.Errorf("NewFingerprint failed to copy the correct bytes into the "+
			"Fingerprint.\nexpected: %+v\nreceived: %+v", fpBytes, fp)
	}

	// Ensure that the data is copied
	fpBytes[2]++
	if fp[2] == fpBytes[2] {
		t.Errorf("NewFingerprint failed to create a copy of the data.")
	}
}

// Happy path.
func TestFingerprint_Bytes(t *testing.T) {
	prng := rand.New(rand.NewSource(42))
	fpBytes := make([]byte, KeyFPLen)
	prng.Read(fpBytes)

	fp := NewFingerprint(fpBytes)
	testFpBytes := fp.Bytes()
	if !bytes.Equal(fpBytes, testFpBytes) {
		t.Errorf("Bytes failed to return the expected bytes."+
			"\nexpected: %+v\nreceived: %+v", fpBytes, testFpBytes)
	}

	// Ensure that the data is copied
	testFpBytes[2]++
	if fp[2] == testFpBytes[2] {
		t.Errorf("Bytes failed to create a copy of the data.")
	}
}

// Happy path.
func TestFingerprint_String(t *testing.T) {
	prng := rand.New(rand.NewSource(42))
	fpBytes := make([]byte, KeyFPLen)
	prng.Read(fpBytes)
	fp := NewFingerprint(fpBytes)

	expectedString := base64.StdEncoding.EncodeToString(fpBytes)
	if expectedString != fp.String() {
		t.Errorf("String failed to return the expected string."+
			"\nexpected: %s\nreceived: %s", expectedString, fp.String())
	}
}

func Test_Fingerprint_JSON_Marshal_Unmarshal(t *testing.T) {
	prng := rand.New(rand.NewSource(74043))
	fpBytes := make([]byte, KeyFPLen)
	prng.Read(fpBytes)

	expected := NewFingerprint(fpBytes)

	data, err := json.Marshal(expected)
	if err != nil {
		t.Errorf("Failed to marshal %T: %+v", expected, err)
	}

	var fp Fingerprint
	if err = json.Unmarshal(data, &fp); err != nil {
		t.Errorf("Failed to unmarshal %T: %+v", fp, err)
	}

	if expected != fp {
		t.Errorf("Unexpected fingerprint.\nexpected: %s\nreceived: %s",
			expected, fp)
	}
}
