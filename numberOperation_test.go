package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNumOperation(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(quoteHandler)
	request, err := http.NewRequest(http.MethodGet, "/counter?operation=decrease", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, request)
	status := recorder.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var result Num
	err = json.NewDecoder(recorder.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	expected := Num{
		OldNumber: 0,
		NewNumber: -1,
	}
	if result!= expected {
		t.Errorf("we got %o result should be %o", result, expected)
	}

	request, err = http.NewRequest(http.MethodGet, "/counter?operation=increase", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, request)
	status = recorder.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	err = json.NewDecoder(recorder.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	expected = Num{
		OldNumber: -1,
		NewNumber: 0,
	}
	if result != expected {
		t.Errorf("we got %o result should be %o", result, expected)
	}
}
func Test400(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(quoteHandler)
	request, err := http.NewRequest(http.MethodGet, "/counter", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, request)
	status := recorder.Code
	if status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
func Test404(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(quoteHandler)
	request, err := http.NewRequest(http.MethodGet, "/counterXX?operation=2increase", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, request)
	status := recorder.Code
	if status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}
func Test501(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(quoteHandler)
	request, err := http.NewRequest(http.MethodGet, "/counter?operation=something", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, request)
	status := recorder.Code
	if status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotImplemented)
	}
}
func Test501Second(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(quoteHandler)
	request, err := http.NewRequest(http.MethodPost, "/counter?operation=decrease", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, request)
	status := recorder.Code
	if status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotImplemented)
	}
}