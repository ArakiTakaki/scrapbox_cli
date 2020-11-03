package interfaces

// Result Error is nil
type HttpResult struct {
	OK         []byte
	Error      error
	LogMessage string
}
