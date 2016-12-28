package tracker

import "net/http"

type TrackingHandler struct {
	serve func(http.ResponseWriter, *http.Request)
	f     func(tracker *Tracker)
}

func (th TrackingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tracker := NewTracker(w)
	th.serve(tracker, r)
	th.f(tracker)
}

func CreateTrackingHandler(ServeHTTP func(http.ResponseWriter, *http.Request), f func(t *Tracker)) TrackingHandler {
	return TrackingHandler{ServeHTTP, f}
}

func TraceOn(ServeHTTP func(http.ResponseWriter, *http.Request), f func(t *Tracker)) TrackingHandler {
	return TrackingHandler{ServeHTTP, f}
}
