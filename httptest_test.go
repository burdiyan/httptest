package httptest

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMockServer(t *testing.T) {
	server := MockServer(
		JSONEndpoint("/test", 200, "Hello World"),
		JSONEndpoint("/foo", 200, `{"foo":true}`),
	)
	defer server.Close()

	cases := []struct {
		url      string
		expected string
	}{
		{"/test", "Hello World"},
		{"/foo", `{"foo":true}`},
	}

	for _, c := range cases {
		resp, err := http.Get(server.URL + c.url)
		if err != nil {
			t.Fatal(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			t.Fatal(err)
		}

		if string(body) != c.expected {
			t.Fatalf("expected: %q, got: %q", c.expected, body)
		}
	}
}
