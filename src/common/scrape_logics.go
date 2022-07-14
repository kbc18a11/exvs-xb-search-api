package common

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type ScrapeLogics interface {
	// 機体情報URL一覧取得
	GetAirframeUrls() []string
}

type ScrapeLogicsImp struct {
}

/*
機体情報URL一覧取得
*/
func (scrapeLogicsImp *ScrapeLogicsImp) GetAirframeUrls() []string {
	res, err := http.Get("https://w.atwiki.jp/exvs2xb/pages/118.html")

	if err != nil {
		// リクエスト先が存在しない場合
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		// 正常にアクセスできなかった場合
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		// htmlドキュメントが取得できなかった場合
		log.Fatal(err)
	}

	// 機体情報URL一覧
	var airframeUrls []string

	// 機体情報URL取得
	doc.Find("tr").Find("a").Each(func(i int, airframe *goquery.Selection) {
		airframeUrl, isAirframeUrl := airframe.Attr("href")

		if !isAirframeUrl {
			// 機体情報URLが存在しない場合
			return
		}

		airframeUrls = append(airframeUrls, "https:"+airframeUrl)
	})

	return airframeUrls
}
