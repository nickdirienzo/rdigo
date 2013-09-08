package rdigo

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type SearchResult struct {
	Result SearchResults `json:"result"`
}

type SearchResults struct {
	AlbumCount    int32         `json:"album_count"`
	PersonCount   int32         `json:"person_count"`
	TrackCount    int32         `json:"track_count"`
	PlaylistCount int32         `json:"playlist_count"`
	ArtistCount   int32         `json:"artist_count"`
	NumberResults int32         `json:"number_results"`
	Results       []interface{} `json:"results"`
}

func (r *Rdio) Search(query, types string, options map[string]string) (SearchResult, error) {
	if options == nil {
		options = make(map[string]string)
	}
	options["method"] = "search"
	options["query"] = query
	options["types"] = types
	resp, err := r.consumer.Post(baseUrl, options, r.AccessToken)
	if err != nil {
		log.Println("Rdio Search Error:", err.Error())
		return SearchResult{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Rdio Search Error:", err.Error())
		return SearchResult{}, err
	}
	rdioResp := SearchResult{}
	err = json.Unmarshal(body, &rdioResp)
	if err != nil {
		log.Println("Rdio Search Error:", err.Error())
		return SearchResult{}, err
	}
	return rdioResp, err
}
