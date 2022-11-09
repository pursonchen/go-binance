package binance

import (
	"context"
	"net/http"
)

// GetUserAssetService fetches all asset detail.
//
// See https://binance-docs.github.io/apidocs/spot/en/#asset-detail-user_data
type GetUserAssetService struct {
	c     *Client
	asset *string
}

// Asset sets the asset parameter.
func (s *GetUserAssetService) Asset(asset string) *GetUserAssetService {
	s.asset = &asset
	return s
}

type NewGetUserAssetService struct {
	c     *Client
	asset *string
}

// Do send request
func (s *GetUserAssetService) Do(ctx context.Context) (res []*UserAssetV3, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v3/asset/getUserAsset",
		secType:  secTypeSigned,
	}

	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return []*UserAssetV3{}, err
	}
	res = make([]*UserAssetV3, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*UserAssetV3{}, err
	}
	return res, nil
}

// UserAsset represents the detail of an asset
type UserAssetV3 struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	Ipoable      string `json:"ipoable"`
	BtcValuation string `json:"btcValuation"`
}
