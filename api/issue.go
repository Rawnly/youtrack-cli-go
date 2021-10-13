package api

import (
	"fmt"
	"github.com/rawnly/youtrack/util"
	"reflect"
	"strings"
)

type User struct {
	Id string `json:"id"`
	Login string `json:"login"`
	FullName string `json:"fullName"`
	Email string `json:"email"`
	Online bool `json:"online"`
}

type Project struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ShortName string `json:"shortName"`
	Description string `json:"description"`
	Archived bool `json:"archived"`
}

type ReducedProject = Project

type Color struct {
	Foreground string `json:"foreground"`
	Background string `json:"background"`
}

type IssueTag struct {
	Name string `json:"name"`
	Color Color `json:"color"`
}
// Missing Project
type Issue struct {
	Id string `json:"id"`
	IdReadable string `json:"idReadable"`
	Created int64 `json:"created"`
	Summary string `json:"summary"`
	Description string `json:"description"`
	Reporter User `json:"reporter"`
	Updater User `json:"updater"`
	Updated int64 `json:"updated"`
	Resolved bool `json:"resolved"`
	Tags []IssueTag `json:"tags"`
	CommentsCount int16 `json:"commentsCount"`
	Fields []CustomField `json:"fields"`
}

type ReducedIssue struct {
	Id string `json:"id"`
	NumberInProject int64 `json:"numberInProject"`
	Created int64 `json:"created"`
	Updated int64 `json:"updated"`
	Summary string `json:"summary"`
	Description string `json:"description"`
	Project ReducedProject `json:"project"`
}

func FetchIssue(issueId string, storage *util.Storage) (*Issue, error) {
	url := fmt.Sprintf("issues/%s?fields=%s$", issueId, util.GetFieldsOf(reflect.TypeOf(Issue{})))
	url = strings.Replace(url, "value)$", fmt.Sprintf("value(%s))", getCustomFieldValueFields()), -1)

	request, err := GetRequest( url, nil)(storage)

	if err != nil {
		return nil, err
	}

	var issue Issue
	if err := ExecuteRequest(request, &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}
