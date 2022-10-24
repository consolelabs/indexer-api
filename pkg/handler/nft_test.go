package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/consolelabs/indexer-api/pkg/config"
	"github.com/consolelabs/indexer-api/pkg/entity"
	mock_nft "github.com/consolelabs/indexer-api/pkg/entity/nft_entity/mocks"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetNftDetail(t *testing.T) {
	type fields struct {
		store    *store.Store
		service  *service.Service
		logger   logger.Logger
		entities *entity.Entity
	}

	type args struct {
		data string
	}

	type want struct {
		response   interface{}
		expectCode int
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := config.LoadTestConfig()
	loggerMock := logger.NewLogrusLogger()
	serviceMock := service.New(&cfg)
	storeMock := store.New(&cfg)
	entityMock := entity.New(storeMock, serviceMock)
	nft := mock_nft.NewMockINftEntity(ctrl)
	entityMock.Nft = nft

	tests := []struct {
		name             string
		fields           fields
		givenAddress     string
		givenToken       string
		wantCode         int
		wantErr          error
		wantResponsePath string
	}{
		{
			name: "ok_get_all",
			fields: fields{
				store:    storeMock,
				service:  serviceMock,
				logger:   loggerMock,
				entities: entityMock,
			},
			wantCode:         200,
			givenAddress:     "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73",
			givenToken:       "1",
			wantErr:          nil,
			wantResponsePath: "testdata/get_nft_detail/200.json",
		},
	}

	modelNft := model.NftToken{
		TokenId:           "1",
		CollectionAddress: "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73",
		Name:              "Cyber Neko 1",
		Description:       "Cyber Neko NFT",
		Amount:            "",
		Image:             "https://storage.googleapis.com/cyber-neko/genesis-collection/full/1",
		ImageCDN:          "https://storage.googleapis.com/cyber-neko/genesis-collection/thumb/1",
		ThumbnailCDN:      "",
		ImageContentType:  "",
		Attributes: []model.NftTokenAttribute{
			{
				TraitType: "mask",
				Value:     "None",
				Count:     6556,
				Rarity:    "Common",
				Frequency: "99.36%",
			},
			{
				TraitType: "earring",
				Value:     "None",
				Count:     4965,
				Rarity:    "Common",
				Frequency: "82.36%",
			},
			{
				TraitType: "eyewear",
				Value:     "None",
				Count:     4629,
				Rarity:    "Common",
				Frequency: "81.13%",
			},
			{
				TraitType: "fur",
				Value:     "None",
				Count:     4041,
				Rarity:    "Common",
				Frequency: "71.12%",
			},
			{
				TraitType: "necklace",
				Value:     "None",
				Count:     4024,
				Rarity:    "Common",
				Frequency: "72.29%",
			},
			{
				TraitType: "body",
				Value:     "White",
				Count:     894,
				Rarity:    "Common",
				Frequency: "15.49%",
			},
			{
				TraitType: "eyes",
				Value:     "Tired Eyes",
				Count:     390,
				Rarity:    "Uncommon",
				Frequency: "5.61%",
			},
			{
				TraitType: "mouth",
				Value:     "Bearded",
				Count:     118,
				Rarity:    "Mythic",
				Frequency: "0.59%",
			},
			{
				TraitType: "background",
				Value:     "Paw Pattern",
				Count:     92,
				Rarity:    "Uncommon",
				Frequency: "1.07%",
			},
			{
				TraitType: "hat",
				Value:     "Playboy",
				Count:     68,
				Rarity:    "Rare",
				Frequency: "0.48%",
			},
			{
				TraitType: "clothe",
				Value:     "Japanese Uniform",
				Count:     67,
				Rarity:    "Rare",
				Frequency: "0.47%",
			},
		},
		Rarity: &model.NftTokenRarity{
			Rank:   914,
			Score:  "3.658723",
			Total:  6666,
			Rarity: "Epic",
		},
		Owner: &model.NftOwner{
			OwnerAddress:      "0xD28Cf82b9B8ee25E3C82923aDF6aA6CC2f220932",
			CollectionAddress: "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73",
			TokenId:           "1",
			ChainId:           250,
			Token:             nil,
			CreatedTime:       time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC),
			LastUpdatedTime:   time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC),
		},
		ListingMarketplace: []model.NftListingMarketplace{},
	}
	nft.EXPECT().GetTokenDetail("0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73", "1").Return(&modelNft, nil).AnyTimes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			router := gin.Default()
			h := &Handler{
				store:    tt.fields.store,
				service:  tt.fields.service,
				logger:   tt.fields.logger,
				entities: tt.fields.entities,
			}
			router.GET("/api/v1/nft/:collection_address/:token_id", h.GetNftDetail)
			recorder := httptest.NewRecorder()

			request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/nft/%s/%s", tt.givenAddress, tt.givenToken), nil)
			router.ServeHTTP(recorder, request)

			require.Equal(t, tt.wantCode, recorder.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), recorder.Body.String(), "[Handler.GetNFTDetail] response mismatched")
		})
	}
}
