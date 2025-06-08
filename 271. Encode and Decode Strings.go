package main

import "strings"

type Codec struct {
	strs []string
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	codec.strs = strs
	return strings.Join(strs, " ")
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	return codec.strs
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
