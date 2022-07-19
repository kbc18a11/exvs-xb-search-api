package requestSchema

// GetAirframesResponse - 機体一覧
type GetAirframesResponse struct {

	// すべての機体数
	Total int32 `json:"total"`

	// 機体一覧
	Airframes []GetAirframesResponseAirframesInner `json:"airframes"`
}
