package controller

import (
    "html/template"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestLivestockExecutesCorrectTemplate(t *testing.T){
    l := new(livestock)
    expected := `livestock template`
    l.livestockTemplate, _ = template.New("").Parse(expected)
    r := httptest.NewRequest(http.MethodGet, "/livestock", nil)
    w := httptest.NewRecorder()

    l.handleLivestock(w, r)

    actual, _ := ioutil.ReadAll(w.Result().Body)

    if string(actual) != expected {
        t.Errorf("Failed  execute correct template")
    }
}

func TestLivestockExecutesCorrectTemplateWithForwardSlash(t *testing.T){
    l := new(livestock)
    expected := `livestock template`
    l.livestockTemplate, _ = template.New("").Parse(expected)
    r := httptest.NewRequest(http.MethodGet, "/livestock/", nil)
    w := httptest.NewRecorder()

    l.handleLivestock(w, r)

    actual, _ := ioutil.ReadAll(w.Result().Body)

    if string(actual) != expected {
        t.Errorf("Failed  execute correct template")
    }
}
