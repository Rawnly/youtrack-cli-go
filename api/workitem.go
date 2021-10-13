package api

import (
	"github.com/rawnly/youtrack/util"
	"reflect"
)

type WorkItem struct {
	Id string `json:"id"`
	Created int64 `json:"created"`
	Date int64 `json:"date"`
	Duration PeriodValue `json:"duration"`
	Updated int64 `json:"updated"`
	Author User `json:"author"`
	Creator User `json:"creator"`
	Text string `json:"text"`
	TextPreview string `json:"text_preview"`
	Issue ReducedIssue `json:"issue"`
	UsesMarkdown bool `json:"uses_markdown"`
	Type struct {
		Id string `json:"id"`
		Name string `json:"name"`
	} `json:"type"`
}

func CreateWorkItem(issueId string, text string, duration string) func (storage *util.Storage) (*WorkItem, error) {
	payload := map[string]interface{} {
		"text": text,
		"duration": map[string]string {
			"presentation": duration,
		},
		"type": struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		}{
			Id: "107-0",
			Name: "Development",
		},
	}

	return func(storage *util.Storage) (*WorkItem, error) {
		params := map[string]string{
			"issueId": issueId,
		}

		if request, err := PostRequest( "issues/{{ .issueId }}/timeTracking/workItems", params, payload)(storage); err != nil {
			return nil, err
		} else {
			var response WorkItem
			
			if err := ExecuteRequest(request, &response); err != nil {
				return nil, err
			}

			return &response, nil
		}
	}
}

type JSON = map[string]interface{}
type JSONArray = []JSON

func GetWorkItems (issueId string) func (storage *util.Storage) (*[]WorkItem, error) {
	return func(storage *util.Storage) (*[]WorkItem, error) {
		params := map[string]string {
			"issueId": issueId,
			"fields": util.GetFieldsOf(reflect.TypeOf(WorkItem{})),
		}

		request, err := GetRequest("issues/{{ .issueId }}/timeTracking/workItems?fields={{ .fields }}", params)(storage)

		if err != nil {
			return nil, err
		}

		var response []WorkItem

		if err := ExecuteRequest(request, &response); err != nil {
			return nil, err
		}

		return &response, nil
	}
}