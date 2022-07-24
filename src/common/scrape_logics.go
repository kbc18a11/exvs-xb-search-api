package common

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// アットウィキの機体情報
type AtWikiAirframeInfo struct {
	TitleOfWorkName   string
	PilotName         string
	AirframeCostValue int
	AwakenTypeName    string
	Name              string
	Hp                int
	AirframeInfoUrl   string
	ThumbnailUrl      string
	IsTransformation  bool
	IsDeformation     bool
}

type ScrapeLogics interface {
	// 機体情報URL一覧取得
	GetAirframeUrls() []string

	// URLから機体情報の取得
	GetAirframeInfo(airframeUrl string) (*AtWikiAirframeInfo, error)
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

/*
URLから機体情報の取得
*/
func (scrapeLogicsImp *ScrapeLogicsImp) GetAirframeInfo(airframeUrl string) (*AtWikiAirframeInfo, error) {
	res, err := http.Get(airframeUrl)

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

	atWikiAirframeInfo := &AtWikiAirframeInfo{}
	atWikiAirframeInfo.AirframeInfoUrl = airframeUrl

	// 機体名の取得
	atWikiAirframeInfo.Name = doc.Find("h2").Find("a").Text()

	// 機体サムネイルの取得
	atWikiAirframeInfo.ThumbnailUrl, _ = doc.Find("source").Attr("data-srcset")

	// 以下の情報を文字列にまとめて取得
	// 作品タイトル、パイロット名、機体コスト、耐久値、変形の有無、換装の有無、覚醒タイプ
	airframeInfos := doc.Find(".atwiki_plugin_divclass").Find("table").Find("tbody").Text()

	if len(regexp.MustCompile("\n").Split(airframeInfos, -1)) != 23 {
		// プレイアブルキャラじゃない場合
		return nil, errors.New("")
	}

	// 作品タイトルの取得
	atWikiAirframeInfo.TitleOfWorkName = strings.TrimSpace(regexp.MustCompile("\n").Split(airframeInfos, -1)[1])

	// アットwikiとマスタの作品タイトル表記が違う機体をマスタの表記に適用させる
	switch atWikiAirframeInfo.Name {
	case "キュベレイMk-II(プル)":
		fallthrough
	case "ザクIII改":
		fallthrough
	case "キュベレイMk-II(プルツー)":
		atWikiAirframeInfo.TitleOfWorkName = "機動戦士ガンダムZZ"
	case "νガンダム":
		fallthrough
	case "ヤクト・ドーガ":
		atWikiAirframeInfo.TitleOfWorkName = "機動戦士ガンダム 逆襲のシャア"
	case "トールギスIII":
		atWikiAirframeInfo.TitleOfWorkName = "新機動戦記ガンダムWEndless Waltz"
	case "アルケーガンダム":
		atWikiAirframeInfo.TitleOfWorkName = "機動戦士ガンダム00"
	case "ブレイヴ指揮官用試験機":
		atWikiAirframeInfo.TitleOfWorkName = "劇場版 機動戦士ガンダム00-A wakening of the Trailblazer-"
	case "G-セルフ":
		atWikiAirframeInfo.TitleOfWorkName = "ガンダム Gのレコンギスタ"
	case "ガンダムEz8":
		atWikiAirframeInfo.TitleOfWorkName = "機動戦士ガンダム 第08MS小隊"
	case "アヴァランチエクシア":
		atWikiAirframeInfo.TitleOfWorkName = "機動戦士ガンダム00V"
	}

	// パイロット名の取得
	atWikiAirframeInfo.PilotName = strings.TrimSpace(regexp.MustCompile("\n").Split(airframeInfos, -1)[3])
	// アットwikiとマスタのパイロット名の表記が違う機体をマスタの表記に適用させる
	switch atWikiAirframeInfo.Name {
	case "アルトロンガンダム":
		atWikiAirframeInfo.PilotName = "張五飛"
	case "ザクII改":
		atWikiAirframeInfo.PilotName = "バーナード・ワイズマン"
	case "ビギナ・ギナII(木星決戦仕様)":
		atWikiAirframeInfo.PilotName = "ギリ・ガデューカ・アスピス"
	case "アストレイブルーフレームD":
		atWikiAirframeInfo.PilotName = "叢雲劾"
	case "アストレイブルーフレームセカンドL":
		atWikiAirframeInfo.PilotName = "叢雲劾"
	case "ガンダムダブルオーダイバーエース":
		atWikiAirframeInfo.PilotName = "ミカミ・リク"
	}

	// 機体コストの取得
	atWikiAirframeInfo.AirframeCostValue, _ = strconv.Atoi(strings.TrimSpace(regexp.MustCompile("\n").Split(airframeInfos, -1)[5]))

	// 耐久値の取得
	atWikiAirframeInfo.Hp, _ = strconv.Atoi(strings.TrimSpace(regexp.MustCompile("\n").Split(airframeInfos, -1)[7]))

	// 形態以降の有無の判定
	if strings.TrimSpace(regexp.MustCompile("\n").Split(airframeInfos, -1)[9]) != "なし" {
		atWikiAirframeInfo.IsTransformation = true
	} else {
		atWikiAirframeInfo.IsTransformation = false
	}

	// 変形の有無の判定
	if strings.TrimSpace(regexp.MustCompile("\n").Split(airframeInfos, -1)[17]) == "あり" {
		atWikiAirframeInfo.IsDeformation = true
	} else {
		atWikiAirframeInfo.IsDeformation = false
	}

	// 覚醒タイプの取得
	atWikiAirframeInfo.AwakenTypeName = strings.TrimSpace(regexp.MustCompile("\n").Split(airframeInfos, -1)[21])

	return atWikiAirframeInfo, nil
}
