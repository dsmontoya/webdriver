package assert

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func HTTPResponseIsValid(t *testing.T, response *http.Response) {
	if statusCode := response.StatusCode; statusCode > 399 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Fatalf("error reading response body: %v", err)
		}
		t.Fatalf("got http status code %d with body: %s", statusCode, body)
	}
}
