package messages

// Response is the basic response on handlers
type Response struct {
	Code int
}

// OKResponse returns a response with code 200
func OKResponse() *Response { _ = "STUB: not implemented"; return nil }
