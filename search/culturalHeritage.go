package search

import (
	"encoding/json"
	"fmt"

	"net/http"

	"strconv"
)

//Search struct is used in order to bind frontend parameters to parameters used in the europeana search api
type Search struct {
	Query       string `form:"query" json:"query" binding:"required"`
	Media       bool   `form:"media" json:"media"`
	Thumbnail   bool   `form:"thumbnail" json:"thumbnail"`
	Landingpage bool   `form:"landingpage" json:"landingpage"`
	Rows        int    `form:"rows" json:"rows"`
}

//CulturalHeritageRecord Simplified Europeana Cultural Heritage Record
type CulturalHeritageRecord struct {
	ID          string   `json:"id"`
	Title       []string `json:"title"`
	Guid        string   `json:"guid"`
	Preview     []string `json:"edmPreview"`
	Image       []string `json:"edmIsShownBy"`
	Country     []string `json:"country"`
	Provider    []string `json:"provider"`
	Description []string `json:"dcDescription"`
	Creator     []string `json:"dcCreator"`
}

//CulturalHeritageItems contains the count of total records and all records retrieved by europeana API
type CulturalHeritageItems struct {
	ItemsCount int                      `json:"itemsCount"`
	Items      []CulturalHeritageRecord `json:"items"`
}

//Service Injectable service
type Service interface {
	SearchCulturalHeritageRecords(searchData Search) CulturalHeritageItems
}

type culturalHeritageRecordService struct {
}

//NewCulturalHeritageRecordService culturalHeritageRecordService constructor
func NewCulturalHeritageRecordService() *culturalHeritageRecordService {
	chs := new(culturalHeritageRecordService)
	return chs
}

//Maps the serach parameters and executes the europeana api call
func (chr *culturalHeritageRecordService) SearchCulturalHeritageRecords(client *http.Client, searchData Search) CulturalHeritageItems {
	req, _ := http.NewRequest("GET", "https://www.europeana.eu/api/v2/search.json", nil)

	// Query params
	q := req.URL.Query()
	q.Add("query", searchData.Query)
	q.Add("rows", strconv.Itoa(searchData.Rows))
	q.Add("media", strconv.FormatBool(searchData.Media))
	q.Add("thumbnail", strconv.FormatBool(searchData.Thumbnail))
	q.Add("landingpage", strconv.FormatBool(searchData.Landingpage))
	q.Add("wskey", "43pWynpwY")

	req.URL.RawQuery = q.Encode()

	items := CulturalHeritageItems{}

	response, err := client.Do(req) //http.Get(req.URL.String())
	if err != nil || response.StatusCode != http.StatusOK {
		return items
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&items)

	if err != nil {
		fmt.Println(err)
		//TODO: log the error here
	}

	return items
}
