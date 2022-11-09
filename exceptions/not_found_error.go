package exceptions

type NotFoundErrorStruct struct {
	ErrorMsg string
}

func NewNotFoundError(msg string) *NotFoundErrorStruct {
	return &NotFoundErrorStruct{
		ErrorMsg: msg,
	}
}

func (e *NotFoundErrorStruct) Error() string {
	return e.ErrorMsg
}
