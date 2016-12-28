package tracker

import "net/http"

type Tracker struct {
	http.ResponseWriter
	status        int
	contentLength int
}

// NewTracker create new tracker for http.ResponseWriter
func NewTracker(w http.ResponseWriter) *Tracker {
	return &Tracker{w, 200, 0}
}

func (t *Tracker) WriteHeader(status int) {
	t.status = status
	t.ResponseWriter.WriteHeader(status)
}

func (t *Tracker) Write(buff []byte) (int, error) {
	var err error
	t.contentLength, err = t.ResponseWriter.Write(buff)
	return t.contentLength, err
}

func (t *Tracker) StatusCode() int {
	return t.status
}
