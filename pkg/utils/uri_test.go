package utils

import "testing"

func TestStandardizeUri(t *testing.T) {
	type args struct {
		uri        string
		ipfsServer string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid uri",
			args: args{
				uri:        "https://backend.pod.so/api/v1/nft/0x97cEDf6e9116f3422b52Ac1ae339685D527721e7/metadata/49859891009376400918199007836401937385395276069393698364865814159156331861131",
				ipfsServer: "ipfs.io",
			},
			want: "https://backend.pod.so/api/v1/nft/0x97cEDf6e9116f3422b52Ac1ae339685D527721e7/metadata/49859891009376400918199007836401937385395276069393698364865814159156331861131",
		},
		{
			name: "trailing uri",
			args: args{
				uri:        "https://backend.pod.so/api/v1/nft/0x97cEDf6e9116f3422b52Ac1ae339685D527721e7/metadata/49859891009376400918199007836401937385395276069393698364865814159156331861131  ",
				ipfsServer: "ipfs.io",
			},
			want: "https://backend.pod.so/api/v1/nft/0x97cEDf6e9116f3422b52Ac1ae339685D527721e7/metadata/49859891009376400918199007836401937385395276069393698364865814159156331861131",
		},
		{
			name: "ipfs uri",
			args: args{
				uri:        "ipfs://0x97cEDf6e9116f3422b52Ac1ae339685D527721e7/metadata/49859891009376400918199007836401937385395276069393698364865814159156331861131",
				ipfsServer: "ipfs.io",
			},
			want: "https://ipfs.io/ipfs/0x97cEDf6e9116f3422b52Ac1ae339685D527721e7/metadata/49859891009376400918199007836401937385395276069393698364865814159156331861131",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardizeUri(tt.args.uri, tt.args.ipfsServer); got != tt.want {
				t.Errorf("StandardizeUri() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StandardizeSortQuery(t *testing.T) {
	type args struct {
		sortQ string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				sortQ: "",
			},
			want: "",
		},
		{
			name: "no sort operator",
			args: args{
				sortQ: "chain_id",
			},
			want: "chain_id ASC",
		},
		{
			name: "multiple empty sort",
			args: args{
				sortQ: ",,,",
			},
			want: "",
		},
		{
			name: "multiple sort",
			args: args{
				sortQ: "-token_id,name",
			},
			want: "token_id DESC,name ASC",
		},
		{
			name: "invalid sort_operator/field",
			args: args{
				sortQ: "token_id?",
			},
			want: "token_id ASC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardizeSortQuery(tt.args.sortQ); got != tt.want {
				t.Errorf("StandardizeUri() = %v, want %v", got, tt.want)
			}
		})
	}
}
