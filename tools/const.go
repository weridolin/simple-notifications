package tools

type CustomError struct {
	Msg  string
	Code int
}

var (
	ModelRecordNotFound = CustomError{
		Msg:  "record not found",
		Code: 10001,
	}

	ModelRecordDeletedError = CustomError{
		Msg:  "record deleted error",
		Code: 10002,
	}

	ModelAlreadyExist = CustomError{
		Msg:  "record already exist",
		Code: 10003,
	}
)
