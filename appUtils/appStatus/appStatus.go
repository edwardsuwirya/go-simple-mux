package appStatus

const (
	Success                 = 0
	StatusNotYetImplemented = 500
	UnknownError            = 6
	ErrorNotMatchValidation = 1
	ErrorLackInfo           = 2
)

var statusMessage = map[int]string{
	Success:                 "Success",
	StatusNotYetImplemented: "Not yet implemented",
	UnknownError:            "Unknown error, please contact IT",
	ErrorNotMatchValidation: "Unsatisfied validation",
	ErrorLackInfo:           "Please fill required %s",
}

func StatusText(code int) string {
	return statusMessage[code]
}
