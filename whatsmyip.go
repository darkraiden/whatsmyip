package whatsmyip

//go:generate mockgen -destination ./mocks/mock_doer.go -package mocks github.com/darkraiden/whatsmyip Doer
//go:generate mockgen -destination ./mocks/mock_getter.go -package mocks github.com/darkraiden/whatsmyip Getter

import (
	"bytes"
	"io"
	"net/http"
)

type Doer interface {
	Do(r *http.Request) (*http.Response, error)
}

type Getter interface {
	Get(Doer) (string, error)
}

var ipGetterURL string = "http://ifconfig.me"

// GetBaseURL gives the user an idea of which site
// is used to fetch the host IP address
func GetBaseURL() string {
	return ipGetterURL
}

// Get fetches the machine's public IP address
func Get(client Doer) (string, error) {
	req, err := http.NewRequest(http.MethodGet, ipGetterURL, nil)
	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	return readCloserToString(res.Body), nil
}

func readCloserToString(rc io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)

	return buf.String()
}