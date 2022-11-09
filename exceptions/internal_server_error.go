package exceptions

type InternalServerErrorStruct struct {
	ErrorMsg string
}

func NewInternalServerError(msg string) *InternalServerErrorStruct {
	return &InternalServerErrorStruct{
		ErrorMsg: msg,
	}
}

func (e *InternalServerErrorStruct) Error() string {
	return e.ErrorMsg
}