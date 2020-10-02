package appStatus

const (
	Success                 = 0
	StatusNotYetImplemented = 500
	UnknownError            = 6
	ErrorNotMatchValidation = 1
)

var statusMessage = map[int]string{
	Success:                 "Success",
	StatusNotYetImplemented: "Not yet implemented",
	UnknownError:            "Unknown error, please contact IT",
	ErrorNotMatchValidation: "Unsatisfied validation",
}

func StatusText(code int) string {
	return statusMessage[code]
}
