package models

// AirframeInfo - 機体情報
type AirframeInfo struct {

	// 機体ID
	Id int32 `json:"id"`

	// 機体名
	Name string `json:"name"`

	// 機体コスト
	Cost int32 `json:"cost"`

	// 作品タイトル名
	TitleOfWork string `json:"titleOfWork"`

	// パイロット名
	Pilot string `json:"pilot,omitempty"`

	// 覚醒タイプ名
	AwakenType string `json:"awakenType"`

	// 機体耐久値
	Hp int32 `json:"hp"`

	// 機体情報URL
	AirframeInfoUrl string `json:"airframeInfoUrl"`

	// サムネイルUrl
	ThumbnailImageUrl string `json:"thumbnailImageUrl"`
}
