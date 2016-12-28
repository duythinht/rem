package tracker

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func th(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func NewTestTracker(f func(http.ResponseWriter, *http.Request)) *Tracker {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	tracker := NewTracker(w, r)

	f(tracker, r)
	return tracker
}

func TestDummyHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	th(w, r)

	if w.Body.String() != "hello" {
		t.Fail()
	}
}

func TestTrackerStatus(t *testing.T) {
	tracker := NewTestTracker(th)
	if tracker.StatusCode() != 200 {
		t.Error(tracker)
		t.Fail()
	}
}

func TestTrackerCustomStatus(t *testing.T) {
	tracker := NewTestTracker(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
	})

	if tracker.StatusCode() != 201 {
		t.Error(tracker)
		t.Fail()
	}
}

func TestTrackerExecTime(t *testing.T) {
	message := "Hello world!"
	tracker := NewTestTracker(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})

	if tracker.contentLength != len(message) {
		t.Error("Not match content length", tracker.contentLength, len(message))
		t.Fail()
	}
}
