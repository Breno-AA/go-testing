package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	testcases := []struct {
		Name     string
		Expected string
	}{
		{Name: "Tiago", Expected: "Hello Tiago"},
		{Name: "Maria", Expected: "Hello Maria"},
		{Name: "", Expected: "Hello World"},
		// {Name: "", Expected: ""},
		// {Name: "", Expected: ""},
	}

	srv := httptest.NewServer(http.HandlerFunc(HandleHello))
	defer srv.Close()

	for _, tc := range testcases {
		path := fmt.Sprintf("%s/?name=%s", srv.URL, tc.Name)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if response.StatusCode != 200 {
			t.Fatalf("expected: HTTP 200, got: HTTP %d", response.StatusCode)
		}
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Fatal(err)
		}

		var respBody map[string]string
		err = json.Unmarshal(body, &respBody)
		if err != nil {
			t.Fatal(err)
		}

		if respBody["message"] != tc.Expected {
			t.Fatalf("expected: %s, got: %s", tc.Expected, respBody["message"])
		}
	}
}
