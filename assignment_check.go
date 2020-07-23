package bs_jvm

// This file contains primarily a function for checking whether assigning one
// object to another is valid, or whether the types mismatch.

import (
	"fmt"
)

// Returns nil if it's okay to overwrite object dst with src. This means the
// types must be compatible.  Largely intended to be used when storing
// variables in fields.
func AssignmentOK(src, dst Object) error {
	if src.IsPrimitive() != dst.IsPrimitive() {
		return TypeError(fmt.Sprintf("Can't overwrite a %s with %s",
			dst.TypeName(), src.TypeName()))
	}
	// TODO: More extensive type checking! At the moment we don't bother
	// checking types except to ensure that both objects are either primitives
	// or non-primitives.
	return nil
}
