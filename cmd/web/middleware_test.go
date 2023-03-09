package main

import (
	"net/http"
	"testing"
)

func TestNosurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Error("type is http.Handler ", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Error("type is http.Handler ", v)
	}
}
