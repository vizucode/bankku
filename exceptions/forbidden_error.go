package exceptions

type ForbiddenErrorStruct struct {
	ErrorMsg string
}

func NewForbiddenError(msg string) *ForbiddenErrorStruct {
	return &ForbiddenErrorStruct{
		ErrorMsg: msg,
	}
}

func (e *ForbiddenErrorStruct) Error() string {
	return e.ErrorMsg
}
