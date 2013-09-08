package rdigo

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type PlaybackTokenResponse struct {
	Result string `json:"result"`
	Status string `json:"status"`
}

func (r *Rdio) GetPlaybackToken(domain string) (PlaybackTokenResponse, error) {
	query := make(map[string]string)
	query["method"] = "getPlaybackToken"
	query["domain"] = domain
	resp, err := r.consumer.Post(baseUrl, query, r.AccessToken)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Rdio Call Error:", err.Error())
		return PlaybackTokenResponse{}, err
	}
	rdioResp := PlaybackTokenResponse{}
	err = json.Unmarshal(body, &rdioResp)
	if err != nil {
		log.Println("Rdio Call Error:", err.Error())
		return PlaybackTokenResponse{}, err
	}
	return rdioResp, err
}
