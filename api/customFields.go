package api

import (
	"github.com/rawnly/youtrack/util"
	"reflect"
)

type PeriodValue struct {
	Id string `json:"id"`
	Minutes int32 `json:"minutes"`
	Presentation string `json:"presentation"`
}

type SingleUserIssueCustomField struct {
	Name string `json:"name"`
	Value User `json:"value"`
	Type string `json:"$type"`
}

type PeriodIssueCustomField struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Value PeriodValue `json:"value"`
}

type CustomFieldValue struct {
	Name string `json:"name"`

	// User
	Id string `json:"id"`
	Login string `json:"login"`
	FullName string `json:"fullName"`
	Email string `json:"email"`
	Online bool `json:"online"`

	// Time
	Minutes int32 `json:"minutes"`
	Presentation string `json:"presentation"`
}

func getCustomFieldValueFields() string {
	return util.GetFieldsOf(reflect.TypeOf(CustomFieldValue{}))
}

type CustomField struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type string `json:"$type"`
	Value interface{} `json:"value"`
}


func GetPeriodicIssueCustomField(issue Issue, fieldName string, storage util.Storage) (*PeriodIssueCustomField, error) {
	params := map[string]string {
		"issueId": issue.IdReadable,
	}

	for _, field := range issue.Fields {
		if field.Name == fieldName {
			params["fieldId"] = field.Id
		}
	}

	params["fields"] = util.GetFieldsOf(reflect.TypeOf(PeriodIssueCustomField{}))

	req, err := GetRequest( "issues/{{ .issueId }}/customFields/{{ .fieldId }}?fields={{ .fields }}", params)(&storage)

	if err != nil {
		return nil, err
	}

	var field PeriodIssueCustomField
	if err := ExecuteRequest(req, &field); err != nil {
		return nil, err
	}

	return &field, nil
}


func TemplateString(str string, data interface{}) (string, error) {
	return util.Template{
		Data: data,
		Template: str,
	}.String()
}