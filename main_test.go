package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type TestCase struct {
	input1 int
	input2 int
	expected int
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []TestCase{
		{1, 2, 3},
		{2, 3, 5},
		{-2, -3, -5},
	}
	for _, c := range cases {
		got := AddTwoNumbers(c.input1, c.input2)
		if got != c.expected {
			t.Errorf("AddTwoNumbers(%d, %d) = %d; want %d", c.input1, c.input2, got, c.expected)
		}
	}
}

var r = gin.Default()

func init() {
	gin.SetMode(gin.TestMode)
	r.GET("/ping", ping)
}
func TestPing(t *testing.T) {
	type testCase struct {
		method string
		path string
		expectedResponseCode int
		expectedBody string
	}
	cases := []testCase{
		{method: "GET", path: "/ping", expectedResponseCode: 200, expectedBody: "{\"message\":\"pong\"}"},
	}

	for _, tc := range cases {
		req := getRequest(tc.method, tc.path)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)
		if tc.expectedResponseCode != w.Code {
			t.Errorf("Expected response code %d, got %d", tc.expectedResponseCode, w.Code)
		}
		if tc.expectedBody != string(body) {
			t.Errorf("Expected body %s, got %s", tc.expectedBody, string(body))
		}
	}
}

func getRequest(method, path string) *http.Request {
	req, _ := http.NewRequest(method, path, nil)
	return req
}

