package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewHTTPClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestSearchCulturalHeritageRecords(t *testing.T) {
	chrs := NewCulturalHeritageRecordService()
	client := NewHTTPClient(GetResponse)

	searchData := Search{
		Query:       "query_to_search",
		Media:       true,
		Thumbnail:   true,
		Landingpage: true,
		Rows:        1,
	}

	respItems := chrs.SearchCulturalHeritageRecords(client, searchData)
	builtItems := buildChItems()
	fmt.Println("quero é cona")
	assert.DeepEqual(t, respItems, builtItems)
}

func GetResponse(req *http.Request) *http.Response {
	var chItems = buildChItems()

	if req == nil {
		return &http.Response{StatusCode: 500, Body: ioutil.NopCloser(bytes.NewBufferString(``))}
	}
	chItemsMarshal, _ := json.Marshal(chItems)
	chItemsJSON := string(chItemsMarshal)

	return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString(chItemsJSON))}
}

func buildChItems() CulturalHeritageItems {
	return CulturalHeritageItems{
		ItemsCount: 1,
		Items: []CulturalHeritageRecord{
			{
				ID:          "1",
				Title:       []string{"teste_item_1"},
				Guid:        "guid",
				Preview:     []string{"preview"},
				Image:       []string{"image"},
				Country:     []string{"portugal"},
				Provider:    []string{"europeana"},
				Description: []string{"description"},
				Creator:     []string{"aruna"},
			},
		},
	}
}
