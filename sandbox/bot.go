package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ChimeraCoder/anaconda"
)

func main() {

	// Json読み込み
	raw, error := ioutil.ReadFile("../TwitterCredentials.json")
	if error != nil {
		fmt.Println(error.Error())
		return
	}

	var twitterCredentials TwitterCredentials
	// 構造体にセット
	json.Unmarshal(raw, &twitterCredentials)

	// 認証
	api := anaconda.NewTwitterApiWithCredentials(twitterCredentials.AccessToken, twitterCredentials.AccessTokenSecret, twitterCredentials.APIKey, twitterCredentials.APISecretKey)
	tweet, error := api.PostTweet("私が知っているなかで、全員同席は思いやりを築くのに最も効果的な方法だ。　The Art Of Agile Development", nil)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	fmt.Println(tweet.Text)
}

type TwitterCredentials struct {
	APIKey            string `json:"apiKey"`
	APISecretKey      string `json:"apiSecretKey"`
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
}
