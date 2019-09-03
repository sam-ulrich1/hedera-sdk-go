package hedera

const (
	SELECTOR_LEN = 4
	SELECTOR_LEN_HEX = 8
)

type FunctionSelector struct {
	NeedsComma bool
	Finished []byte
	Complete bool
}

func NewFunctionSelector(function string) *FunctionSelector {
	return &FunctionSelector{
		NeedsComma: false,
		Finished: []byte(function + "("),
		Complete: false,
	}
}

func (fs *FunctionSelector) AddParamType(paramType string) {
	if fs.NeedsComma == true {
		fs.Finished = append(fs.Finished, ","...)
	}
	fs.Finished = append(fs.Finished, paramType...)
	fs.NeedsComma = true
}

func (fs *FunctionSelector) FinishIntermediate() []byte {
	var f []byte
	if fs.Complete == false {
		f = append(fs.Finished, ")"...)
	} else { f = fs.Finished }
	return Keccak256(f)
}

func (fs *FunctionSelector) Finish() []byte {
	if fs.Complete == false {
		fs.Finished = append(fs.Finished, ")"...)
		fs.Complete = true
	}
	return Keccak256(fs.Finished)
}
