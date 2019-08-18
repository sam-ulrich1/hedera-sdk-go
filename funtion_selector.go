package hedera

const (
	SELECTOR_LEN = 4
	SELECTOR_LEN_HEX = 8
)

type FunctionSelector struct {
	NeedsComma bool
	Finished []byte
	Debug string
	Complete bool
}

func NewFunctionSelector(function string) *FunctionSelector {
	return &FunctionSelector{
		NeedsComma: false,
		Finished: []byte(function + "("),
		Debug: function + "(",
		Complete: false,
	}
}

func (fs *FunctionSelector) AddParamType(paramType string) {
	if fs.NeedsComma == true {
		fs.Finished = append(fs.Finished, ","...)
		fs.Debug += ","
	}
	fs.Finished = append(fs.Finished, paramType...)
	fs.Debug += paramType
	fs.NeedsComma = true
}

func (fs *FunctionSelector) FinishIntermediate() []byte {
	var f []byte
	if fs.Complete == false {
		f = append(fs.Finished, ")"...)
	}
	return Keccak256(f)
}

func (fs *FunctionSelector) Finish() []byte {
	if fs.Complete == false {
		fs.Finished = append(fs.Finished, ")"...)
		fs.Debug += ")"
		fs.Complete = true
	}
	return Keccak256(fs.Finished)
}
