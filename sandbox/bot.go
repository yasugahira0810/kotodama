package main

import (
	"github.com/ChimeraCoder/anaconda"
)

func main() {
	anaconda.SetConsumerKey("CONSUMER_KEY")
	anaconda.SetConsumerSecret("CONSUMER_SECRET")
	api := anaconda.NewTwitterApi("ACCESS_TOKEN", "ACCESS_TOKEN_SECRET")
	api.PostTweet("規模を見積り、期間は導出する(アジャイルな見積りと計画づくり)", nil)
}
