package rdigo

import (
	"encoding/json"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"log"
)

const baseUrl = "http://api.rdio.com/1/"

type Rdio struct {
	Consumer    *oauth.Consumer
	AccessToken *oauth.AccessToken
}

func NewClient(consumerKey, consumerSecret string) Rdio {
	c := oauth.NewConsumer(consumerKey, consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "http://api.rdio.com/oauth/request_token",
			AuthorizeTokenUrl: "https://www.rdio.com/oauth/authorize",
			AccessTokenUrl:    "http://api.rdio.com/oauth/access_token",
		})
	return Rdio{Consumer: c}
}

func (r *Rdio) BeginAuthentication(callbackUrl string) (*oauth.RequestToken, string, error) {
	return r.Consumer.GetRequestTokenAndUrl(callbackUrl)
}

func (r *Rdio) CompleteAuthentication(requestToken, requestTokenSecret, verifier string) error {
	rToken := oauth.RequestToken{Token: requestToken, Secret: requestTokenSecret}
	token, err := r.Consumer.AuthorizeToken(&rToken, verifier)
	r.AccessToken = token
	return err
}

// Authenticated only for now
func (r *Rdio) Call(method string, query map[string]string) (interface{}, error) {
	query["method"] = method
	resp, err := r.Consumer.Post(baseUrl, query, r.AccessToken)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Rdio Call Error:", err.Error())
		return nil, err
	}
	var rdioResp interface{}
	err = json.Unmarshal(body, &rdioResp)
	if err != nil {
		log.Println("Rdio Call Error:", err.Error())
		return nil, err
	}
	return rdioResp, err
}
