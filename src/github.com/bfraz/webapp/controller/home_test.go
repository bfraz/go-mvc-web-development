package controller

import (
    "html/template"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHomeExecutesCorrectTemplate(t *testing.T){
    h := new(home)
    expected := `home template`
    h.homeTemplate, _ = template.New("").Parse(expected)
    r := httptest.NewRequest(http.MethodGet, "/home", nil)
    w := httptest.NewRecorder()

    h.handleHome(w, r)

    actual, _ := ioutil.ReadAll(w.Result().Body)

    if string(actual) != expected {
        t.Errorf("Failed  execute correct template")
    }
}


func TestHomeExecutesCorrectTemplateWithForwardSlash(t *testing.T){
    h := new(home)
    expected := `home template`
    h.homeTemplate, _ = template.New("").Parse(expected)
    r := httptest.NewRequest(http.MethodGet, "/", nil)
    w := httptest.NewRecorder()

    h.handleHome(w, r)

    actual, _ := ioutil.ReadAll(w.Result().Body)

    if string(actual) != expected {
        t.Errorf("Failed  execute correct template")
    }
}
