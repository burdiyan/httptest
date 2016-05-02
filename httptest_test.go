package httptest

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMockServer(t *testing.T) {
	server := MockServer(&Endpoint{"/test", 200, "Hello World"})
	defer server.Close()

	response, err := http.Get(server.URL + "/test")
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	if string(body) != "Hello World" {
		t.Error("Fake HTTP server is not working properly!")
	}

	if response.StatusCode != 200 {
		t.Errorf("Expected 200 response code, %d given", response.StatusCode)
	}

	if response.Header.Get("Content-Type") != "application/json" {
		t.Error("Content-Type should be application/json!")
	}
}
