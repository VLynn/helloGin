package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
)

func TestRootRoute(t *testing.T) {
    router := setupRouter()
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/", nil)
    router.ServeHTTP(w, req)

    if w.Code != 200 {
        t.Log("http code not 200")
        t.Fail()
    }

    if w.Body.String() != "hello world" {
        t.Log("response is not right")
        t.Fail()
    }
}