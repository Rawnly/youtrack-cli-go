package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/purell"
	"github.com/rawnly/youtrack/util"
	"io"
	"io/ioutil"
	"net/http"
  "github.com/sirupsen/logrus"
)

type HTTPError struct {
	Message string `json:"error"`
	Description string `json:"error_description"`
}

func Request(method string, pathname string, templateData map[string]string, payload io.Reader) func (storage *util.Storage) (*http.Request, error) {
	return func(storage *util.Storage) (*http.Request, error) {
		urlPath := pathname

		if templateData != nil {
			tmp, err := TemplateString(pathname, templateData)

			if err != nil {
				return nil, err
			}

			urlPath = tmp
		}

		url := fmt.Sprintf("%s/%s", storage.Url, urlPath)

		normalized, err := purell.NormalizeURLString(url, purell.FlagRemoveTrailingSlash | purell.FlagsSafe)

		if err != nil {
			return nil, err
		}

    logrus.Debug(normalized)


		req, err := http.NewRequest(method, normalized, payload)

		if err != nil {
			return nil, err
		}

		return AddAuthorization(storage)(req), nil
	}
}

func GetRequest(url string, data map[string]string) func (storage *util.Storage) (*http.Request, error) {
	return Request("GET", url, data, nil)
}

func PostRequest(url string, data map[string]string, payload interface{}) func (storage *util.Storage) (*http.Request, error) {
	return func(storage *util.Storage) (*http.Request, error) {
		requestPayload, err := json.Marshal(payload)

		fmt.Println("Post Payload", payload)

		if err != nil {
			return nil, err
		}

		req, err := Request("POST", url, data, bytes.NewBuffer(requestPayload))(storage)

		if err != nil {
			return nil, err
		}

		return SetJsonContent(req), nil
	}
}

func AddAuthorization(storage *util.Storage) func (request *http.Request) *http.Request  {
	return func(request *http.Request) *http.Request {
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", storage.Token))

		return request
	}
}

func SetJsonContent(request *http.Request) *http.Request {
	request.Header.Set("Content-Type", "application/json")

	return request
}

func ExecuteRequest(request *http.Request, responseType interface{}) error {
	client := http.Client{}
	res, err := client.Do(request)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 300 || res.StatusCode < 200 {
		var errorResponse HTTPError

		if err := json.Unmarshal(body, &errorResponse); err != nil {
			return errors.New(res.Status)
		}

		errorMessage, err := TemplateString("Error: {{ .Description }}", errorResponse)

		if err != nil {
			return errors.New(res.Status)
		}

		return errors.New( errorMessage )
	}



	return json.Unmarshal(body, &responseType)
}
