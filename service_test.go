package gweb

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirectToGithub(t *testing.T) {
	s := &GoostService{}
	ts := httptest.NewServer(s)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/encoding/base64")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
