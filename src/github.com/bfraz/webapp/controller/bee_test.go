package controller

import (
  "html/template"
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestBeeExecutesCorrectTemplate(t *testing.T){
  b := new(bee)
  expected := `bee template`
  b.beeTemplate, _ = template.New("").Parse(expected)
  r := httptest.NewRequest(http.MethodGet, "/bee", nil)
	w := httptest.NewRecorder()

	b.handleBee(w, r)

	actual, _ := ioutil.ReadAll(w.Result().Body)

	if string(actual) != expected {
		t.Errorf("Failed  execute correct template")
	}

}

func TestBeeExecutesCorrectTemplateWithForwardSlash(t *testing.T){
  b := new(bee)
  expected := `bee template`
  b.beeTemplate, _ = template.New("").Parse(expected)
  r := httptest.NewRequest(http.MethodGet, "/bee/", nil)
	w := httptest.NewRecorder()

	b.handleBee(w, r)

	actual, _ := ioutil.ReadAll(w.Result().Body)

	if string(actual) != expected {
		t.Errorf("Failed  execute correct template")
	}

}
