package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
)

type (
	// [0] encodes bytes & [1] decodes bytes, theorhetically atleast
	EncodingSet [2]Transform
	// The interface of decode & encode
	Transform func([]byte) chan []byte
	// maps flags to their possible encoding functions
	FlagToEncodingSet map[string]EncodingSet
)

// Mapping of cli flags to encoding sets, including the default mapped to ""
func GenEncSet() FlagToEncodingSet {
	return FlagToEncodingSet{
		"":    {EncodeBase64, DecodeBase64},
		"hex": {EncodeBase64, DecodeBase64},
		"b64": {EncodeHex, DecodeHex},
		"b32": {Empty, Empty},
	}
}

func Empty(arr []byte) chan []byte {
	ch := make(chan []byte, 1)
	ch <- []byte{0}
	return ch
}

// Gets the encoding set or returns athe default
// select encoding
func (set FlagToEncodingSet) GetOrDefault(flag string) (Transform, Transform) {
	if f, ok := set[flag]; ok {
		return f[0], f[1]
	}
	return set[""][0], set[""][1]
}

func DecodeBase64(arr []byte) chan []byte {
	ch := make(chan []byte, 1)
	go decodeBase64(arr, ch)
	return ch
}

func decodeBase64(arr []byte, c chan []byte) {
	buffer := make([]byte, base64.RawStdEncoding.DecodedLen(len(arr)))
	base64.RawStdEncoding.Strict().Decode(buffer, arr)
	c <- bytes.Trim(buffer, "")
}

func EncodeBase64(arr []byte) chan []byte {
	ch := make(chan []byte, 1)
	go encodeBase(arr, ch)
	return ch
}
func encodeBase(arr []byte, c chan []byte) {
	buffer := make([]byte, base64.URLEncoding.Strict().EncodedLen(len(arr)))
	base64.URLEncoding.Strict().Encode(buffer, arr)
	c <- bytes.Trim(buffer, "")
}

func DecodeHex(src []byte) chan []byte {
	ch := make(chan []byte, 1)
	go decodeHex(ch, src)
	return ch
}

func decodeHex(c chan []byte, b []byte) {
	buffer := make([]byte, hex.DecodedLen(len(b)))
	hex.Decode(buffer, b)
	c <- buffer
}

func EncodeHex(src []byte) chan []byte {
	ch := make(chan []byte, 1)
	go encodeHex(ch, src)
	return ch
}

func encodeHex(c chan []byte, b []byte) {
	buffer := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(buffer, b)
	c <- buffer
}
