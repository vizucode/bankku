package exceptions

type BadRequestErrorStruct struct {
	ErrorMsg string
}

func NewBadRequestError(msg string) *BadRequestErrorStruct {
	return &BadRequestErrorStruct{
		ErrorMsg: msg,
	}
}

func (e *BadRequestErrorStruct) Error() string {
	return e.ErrorMsg
}
