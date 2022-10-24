package chainstore

import (
	"reflect"
	"testing"

	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/request"
	"github.com/consolelabs/indexer-api/pkg/utils"
	"github.com/consolelabs/indexer-api/pkg/utils/testhelper"
)

func Test_store_Get(t *testing.T) {

	type args struct {
		query request.GetChainsQuery
	}
	tests := []struct {
		name      string
		args      args
		want      []model.Chain
		wantTotal int64
		wantErr   bool
	}{
		{
			name: "ok_all_chains",
			args: args{query: request.GetChainsQuery{
				Pagination: &model.Pagination{Page: 0, Size: 10},
			}},
			wantTotal: 3,
			want: []model.Chain{
				{
					Id:         1,
					Name:       "Ethereum",
					Symbol:     "ETH",
					ChainId:    1,
					ScanUrl:    "",
					ScanApiUrl: "",
					ScanApiKey: "",
					IsEvm:      true,
				},
				{
					Id:         2,
					Name:       "Fantom Opera",
					Symbol:     "FTM",
					ChainId:    250,
					ScanUrl:    "",
					ScanApiUrl: "",
					ScanApiKey: "",
					IsEvm:      true,
				},
				{
					Id:         3,
					Name:       "Optimism",
					Symbol:     "OP",
					ChainId:    10,
					ScanUrl:    "",
					ScanApiUrl: "",
					ScanApiKey: "",
					IsEvm:      true,
				},
			},
		},
		{
			name: "ok_all_chains_page_1",
			args: args{query: request.GetChainsQuery{
				Pagination: &model.Pagination{Page: 0, Size: 1},
			}},
			wantTotal: 3,
			want: []model.Chain{
				{
					Id:         1,
					Name:       "Ethereum",
					Symbol:     "ETH",
					ChainId:    1,
					ScanUrl:    "",
					ScanApiUrl: "",
					ScanApiKey: "",
					IsEvm:      true,
				},
			},
		},
		{
			name: "ok_all_chains_filter",
			args: args{query: request.GetChainsQuery{
				Pagination: &model.Pagination{Page: 0, Size: 1},
				ChainId:    utils.InlinePointer[uint64](1),
				Name:       utils.InlinePointer("Ether"),
				Symbol:     utils.InlinePointer("eth"),
				IsEvm:      utils.InlinePointer(true),
			}},
			wantTotal: 1,
			want: []model.Chain{
				{
					Id:         1,
					Name:       "Ethereum",
					Symbol:     "ETH",
					ChainId:    1,
					ScanUrl:    "",
					ScanApiUrl: "",
					ScanApiKey: "",
					IsEvm:      true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &store{
				db: testhelper.LoadTestDB(),
			}
			got, gotTotal, err := s.Get(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("store.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("store.Get() got = %v, want %v", got, tt.want)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("store.Get() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
