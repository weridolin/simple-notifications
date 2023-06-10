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

	ModelRecordCreatedError = CustomError{
		Msg:  "record created error",
		Code: 10004,
	}

	ModelRecordUpdatedError = CustomError{
		Msg:  "record updated error",
		Code: 10005,
	}

	ModelQueryError = CustomError{
		Msg:  "record query error",
		Code: 10006,
	}

	InternalError = CustomError{
		Msg:  "internal error",
		Code: 10007,
	}
)
