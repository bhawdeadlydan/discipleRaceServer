package helper


import (
	"github.com/tomdionysus/binarytree"
)
// Int64Key is a type of base type int64 that implements the Comparable interface.
type Int64Key int64

// Return true if this key is less than the supplied Int64Key.
func (me Int64Key) LessThan(other binarytree.Comparable) bool {
	return me < other.(Int64Key)
}

// Return true if this key is equal to the supplied Int64Key.
func (me Int64Key) EqualTo(other binarytree.Comparable) bool {
	return me == other.(Int64Key)
}

// Return true if this key is greater than the supplied Int64Key.
func (me Int64Key) GreaterThan(other binarytree.Comparable) bool {
	return me > other.(Int64Key)
}

// Return the int value as an interface
func (me Int64Key) ValueOf() interface{} {
	return int(me)
}
