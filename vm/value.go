package vm

type Value = int64

type ValueArray struct {
	values []Value
}

func (v *ValueArray) init (buf []int64) {
	v.values = buf
}

func (v *ValueArray) write (value Value) error {
	n := len(v.values)
	if n == cap(v.values) {
		return ErrorExaustedBuf
	}
	v.values = v.values[:n+1]
	v.values[n] = value
	return nil;
}

func (v *ValueArray) free () {
	v.values = nil
}