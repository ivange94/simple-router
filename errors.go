package simplerouter

import "net/http"

type HTTPError struct {
	Status int
	Code   string
	Msg    string
	Err    error
}

func (e *HTTPError) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return http.StatusText(e.Status)
}

func (e *HTTPError) Unwrap() error { return e.Err }

func BadRequest(msg string) error {
	return &HTTPError{Status: http.StatusBadRequest, Code: "bad_request", Msg: msg}
}

func NotFound(msg string) error {
	return &HTTPError{Status: http.StatusNotFound, Code: "not_found", Msg: msg}
}
