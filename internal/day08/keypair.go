package day08

// This type is meant to be a composite key for a
// hashmap where the sub-keys are meant to be sorted
// such that the order of their combination does not
// matter. Because of this, it is recommended to NEVER
// create a custom instance of this, but to use the
// provided `NewKeyPair` function instead.
type KeyPair struct {
	Max int
	Min int
}

// Takes any integers (a, b) and creates a new
// KeyPair instance, regardless of the order of
// either of the two integer values.
func NewKeyPair(a, b int) KeyPair {
	// Assume the keys are in order
	if a > b {
		return KeyPair{Max: a, Min: b}
	}

	// And if not, switch them
	return KeyPair{Max: b, Min: a}
}
