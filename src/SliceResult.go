package src

type SliceResult struct {
	Result []interface{}
	Err    error
}

func NewSliceResult(result []interface{}, err error) *SliceResult {
	return &SliceResult{Result: result, Err: err}
}

func (this *SliceResult) Iterator() *Iterator {
	return NewIterator(this.Result)
}
