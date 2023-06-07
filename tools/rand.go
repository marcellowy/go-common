// Package tools
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package tools

import (
	"time"
	"unsafe"
)

// randomType 随机类型
type randomType int

const (
	// RandomLowercase lowercase
	RandomLowercase randomType = 2 << iota
	// RandomMajuscule majuscule
	RandomMajuscule
	// RandomDigital digital
	RandomDigital
	// RandomSymbol symbol
	RandomSymbol
	// RandomAll all
	RandomAll = RandomLowercase | RandomMajuscule | RandomDigital | RandomSymbol
)

// randomCharMaxLength all char length
const randomCharMaxLength = 69

// randomLowercase lowercase
var randomLowercase = []byte{
	0x61, 0x62, 0x63, 0x64, 0x65,
	0x66, 0x67, 0x68, 0x69, 0x6A,
	0x6B, 0x6D, 0x6E, 0x70, 0x71,
	0x72, 0x73, 0x74, 0x75, 0x76,
	0x77, 0x78, 0x79, 0x7A,
}

// randomLowercaseLength randomLowercase length
const randomLowercaseLength = 24

// randomMajuscule majuscule
var randomMajuscule = []byte{
	0x41, 0x42, 0x43, 0x44, 0x45,
	0x46, 0x47, 0x48, 0x49, 0x4A,
	0x4B, 0x4D, 0x4E, 0x50, 0x51,
	0x52, 0x53, 0x54, 0x55, 0x56,
	0x57, 0x58, 0x59, 0x5A,
}

// randomMajusculeLength randomMajuscule length
const randomMajusculeLength = 24

// randomDigital digital
var randomDigital = []byte{
	'9', '2', '3', '4', '5',
	'6', '7', '8',
}

// randomDigitalLength randomDigital length
const randomDigitalLength = 8

// randomSymbol symbol
var randomSymbol = []byte{
	'!', '@', '#', '$', '%',
	'^', '&', '*', '(', ')',
	'_', '-', '+',
}

// randomSymbolLength randomSymbol length
const randomSymbolLength = 13

// RandomString random a string
func RandomString(length uint, charType randomType) (s string) {

	var (
		randChar   = make([]byte, randomCharMaxLength)
		strChar    = make([]byte, length)
		randLength uint32
	)

	if charType&RandomLowercase > 0 {
		copy(randChar[randLength:], randomLowercase)
		randLength += randomLowercaseLength
	}
	if charType&RandomMajuscule > 0 {
		copy(randChar[randLength:], randomMajuscule)
		randLength += randomMajusculeLength
	}
	if charType&RandomSymbol > 0 {
		copy(randChar[randLength:], randomSymbol)
		randLength += randomSymbolLength
	}
	if charType&RandomDigital > 0 {
		copy(randChar[randLength:], randomDigital)
		randLength += randomDigitalLength
	}

	var i uint

	rdm := &random{}
	nn := randLength - 1
	for i = 0; i < length; i++ {
		strChar[i] = randChar[rdm.Uint32n(nn)]
	}
	return *(*string)(unsafe.Pointer(&strChar))
}

type random struct {
	seed uint32
}

func (r *random) Uint32() uint32 {
	for r.seed == 0 {
		x := time.Now().UnixNano()
		r.seed = uint32((x >> 32) ^ x)
	}

	// See https://en.wikipedia.org/wiki/Xorshift
	randomNum := r.seed
	randomNum ^= randomNum << 13
	randomNum ^= randomNum >> 17
	randomNum ^= randomNum << 5
	r.seed = randomNum
	return randomNum
}

func (r *random) Uint32n(maxN uint32) uint32 {
	x := r.Uint32()
	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return uint32((uint64(x) * uint64(maxN)) >> 32)
}
