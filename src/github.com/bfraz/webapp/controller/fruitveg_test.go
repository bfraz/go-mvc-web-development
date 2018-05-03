package controller

import (
  "html/template"
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestFVExecutesCorrectTemplate(t *testing.T){
  fv := new(fruitveg)
  expected := `fruitveg template`
  fv.fruitvegTemplate, _ = template.New("").Parse(expected)
  r := httptest.NewRequest(http.MethodGet, "/fruitveg", nil)
	w := httptest.NewRecorder()

	fv.handleFruitVeg(w, r)

	actual, _ := ioutil.ReadAll(w.Result().Body)

	if string(actual) != expected {
		t.Errorf("Failed  execute correct template")
	}

}

func TestFVExecutesCorrectTemplateWithForwardSlash(t *testing.T){
  fv := new(fruitveg)
  expected := `fruitveg template`
  fv.fruitvegTemplate, _ = template.New("").Parse(expected)
  r := httptest.NewRequest(http.MethodGet, "/fruitveg/", nil)
	w := httptest.NewRecorder()

	fv.handleFruitVeg(w, r)

	actual, _ := ioutil.ReadAll(w.Result().Body)

	if string(actual) != expected {
		t.Errorf("Failed  execute correct template")
	}
}
