package chess

// State is an internal representation of a chess game.
// - FEN notation: https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation
// - See: http://golang-sizeof.tips/?t=Ly8gTm90ZXM6Ci8vIC0gc2l6ZXM6IGh0dHBzOi8vZ29sYW5nLm9yZy9wa2cvYnVpbHRpbi8KLy8gLSBGRU4gbm90YXRpb246IGh0dHBzOi8vZW4ud2lraXBlZGlhLm9yZy93aWtpL0ZvcnN5dGglRTIlODAlOTNFZHdhcmRzX05vdGF0aW9uCgpzdHJ1Y3QgewoJYm9hcmQgc3RyaW5nIC8vIGFyYml0cmFyeSBzaXplCglpc0JsYWNrIGJvb2wgLy8gZGV0ZXJtaW5lcyBhY3RpdmUgY29sb3IgKGNvdWxkIG1hcCBvbiBjYXN0bGluZykKCWNhc3RsaW5nIHVpbnQ4IC8vIGJpdCBtYXNrZWQgbnVtYmVyIGluIEtRa3Egb3JkZXIKCWVuUGFzc2FudCB1aW50OCAvLyBib2FyZCBwb3NpdGlvbiAwIChuL2EpICsgMSAtIDY0CgloYWxmbW92ZSB1aW50OCAvLyBtYXggNTAgKGxpbWl0ZWQgYnkgcnVsZSkgW3R5cGU6MjU1XQoJY291bnQgdWludDMyIC8vIG1heCBvZiA0Mjk0OTY3Mjk1IChsaW1pdGVkIGJ5IHR5cGUpCn0KCi8vIENhbiBpZ25vcmUgdGhlIGxhc3QgMiBmb3IgcG9zc2libGUgbW92ZXMgbWFw
// - sizes: https://golang.org/pkg/builtin/
type State struct {
	board     string // arbitrary size
	isBlack   bool   // determines active color (could map on castling)
	castling  uint8  // bit masked number in KQkq order
	enPassant uint8  // board position 0 (n/a) + 1 - 64
	halfmove  uint8  // max 50 (limited by rule) [type:255]
	count     uint32 // max of 4294967295 (limited by type)
}

// Can ignore the last 2 for possible moves map
