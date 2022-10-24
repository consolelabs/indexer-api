package utils

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestIsEvmAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid lowercase address",
			args: args{
				address: "0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4",
			},
			want: true,
		},
		{
			name: "valid normal hash address",
			args: args{
				address: "0x8E215D06Ea7EC1Fdb4fC5fD21768F4B34eE92EF4",
			},
			want: true,
		},
		{
			name: "invalid address 1",
			args: args{
				address: "8E215D06Ea7EC1Fdb4fC5fD21768F4B34eE92EF4",
			},
			want: false,
		},
		{
			name: "invalid address 2",
			args: args{
				address: "0x00000000000000008E215D06Ea7EC1Fdb4fC5fD21768F4B34eE92EF4",
			},
			want: false,
		},
		{
			name: "invalid address 3",
			args: args{
				address: "123",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEvmAddress(tt.args.address); got != tt.want {
				t.Errorf("IsEvmAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualAddress(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid type string",
			args: args{
				x: "0x00000000000000008E215D06Ea7EC1Fdb4fC5fD21768F4B34eE92EF4",
				y: "0x00000000000000008E215D06Ea7EC1Fdb4fC5fD21768F4B34eE92EF4",
			},
			want: true,
		},
		{
			name: "invalid type string",
			args: args{
				x: "0x10000000000000008E215D06Ea7EC1Fdb4fC5fD21768F4B34eE92EF4",
				y: "0x00000000000000008E215D06Ea7EC1Fdb4fC5fD21768F4B34eE92EF4",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualAddress(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("IsEvmAddress() = %v, want %v", got, tt.want)
			}
		})
	}

	// test case common.Hash
	type args1 struct {
		x common.Hash
		y common.Hash
	}
	tests1 := []struct {
		name string
		args args1
		want bool
	}{
		{
			name: "valid type string",
			args: args1{
				x: common.HexToHash("0x00000000000000008E215D06Ea7EC1Fdb4fC5fD21768F4B34eE92EF4"),
				y: common.HexToHash("0x00000000000000008E215D06Ea7EC1Fdb4fC5fD21768F4B34eE92EF4"),
			},
			want: true,
		},
		{
			name: "invalid type hash",
			args: args1{
				x: common.HexToHash("0x00000000000000008E215D06Ea7EC1Fdb4fC5fD21768F4B34eE92EF4"),
				y: common.HexToHash("0x01000000000000008E215D06Ea7EC1Fdb4fC5fD21768F4B34eE92EF4"),
			},
			want: false,
		},
	}
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualAddress(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("IsEvmAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToHashAddress(t *testing.T) {
	type args struct {
		addrStr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Get checksum address of correct format",
			args: args{
				addrStr: "0x97cEDf6e9116f3422b52Ac1ae339685D527721e7",
			},
			want:    "0x97cEDf6e9116f3422b52Ac1ae339685D527721e7",
			wantErr: false,
		},
		{
			name: "Get check sum address of lowercase format",
			args: args{
				addrStr: "0x97cedf6e9116f3422b52ac1ae339685D527721e7",
			},
			want:    "0x97cEDf6e9116f3422b52Ac1ae339685D527721e7",
			wantErr: false,
		},
		{
			name: "Get checksum error",
			args: args{
				addrStr: "0x12x",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToHashAddress(tt.args.addrStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToChecksumAddr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertToChecksumAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
