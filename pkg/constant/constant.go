package constant

import "github.com/ethereum/go-ethereum/common"

type Topic struct {
	TopicHash   string
	InterfaceId [4]byte
}

type Constant struct {
	ERC721Topic              map[string]Topic
	PaintswapTopic           map[string]Topic
	OpenseaTopic             map[string]Topic
	X2Y2Topic                map[string]Topic
	NFTKeyTopic              map[string]Topic
	QuixoticTopic            map[string]Topic
	ZeroAddressTopic         string
	ZeroAddressHash          common.Hash
	ChainId                  ChainId
	ERC721ContractAddress    string
	PaintswapContractAddress string
	X2Y2ContractAddress      string
	NFTKeyContractAddress    string
	OpenseaContractAddress   string
	QuixoticContractAddress  string
}

// TODO: beautify
var C = Constant{
	ZeroAddressTopic: "0x0000000000000000000000000000000000000000000000000000000000000000",
	ZeroAddressHash:  common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
	ChainId: ChainId{
		Eth: 1,
		Ftm: 250,
		Opt: 10,
		Bsc: 56,
		Apt: 99999999,
	},
	ERC721ContractAddress:    "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
	PaintswapContractAddress: "0x6125fD14b6790d5F66509B7aa53274c93dAE70B9",
	X2Y2ContractAddress:      "0x74312363e45DCaBA76c59ec49a7Aa8A65a67EeD3",
	NFTKeyContractAddress:    "0x1A7d6ed890b6C284271AD27E7AbE8Fb5211D0739",
	OpenseaContractAddress:   "0x00000000006c3852cbEf3e08E8dF289169EdE581",
	QuixoticContractAddress:  "0x20975Da6eB930d592b9d78f451a9156DB5E4C77b",
	ERC721Topic: map[string]Topic{
		"Transfer": {
			TopicHash:   "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
			InterfaceId: [4]byte{128, 172, 88, 205},
		},
	},
	PaintswapTopic: map[string]Topic{
		"CancelledSale": {
			TopicHash: "0x91ee9eb513284d27f660ec3907ad4fd27415782b1faa9b5fa46898dcb2f6983a",
		},
		"Sold": {
			TopicHash: "0x0cda439d506dbc3b73fe10f062cf285c4e75fe85d310decf4b8239841879ed61",
		},
		"NewSale": {
			TopicHash: "0xb3cd69c2fae51b3e21f25978f3604c96acf27abbf4fcb05685b102b4a567cb1d",
		},
		"SaleFinished": {
			TopicHash: "0xd1f4aab9b0d4b94e26beff68299512c1e11b250c148d243d4a9dd10bed2dc514",
		},
	},
	OpenseaTopic: map[string]Topic{
		"OrderFulfilled": {
			TopicHash: "0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31",
		},
		"OrderCancelled": {
			TopicHash: "0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d",
		},
	},

	X2Y2Topic: map[string]Topic{
		"EvInventory": {
			TopicHash: "0x3cbb63f144840e5b1b0a38a7c19211d2e89de4d7c5faf8b2d3c1776c302d1d33",
		},
	},
	NFTKeyTopic: map[string]Topic{
		"TokenBidAccepted": {
			TopicHash: "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
		},
		"TokenBidEntered": {
			TopicHash: "0xc547e24584f4dd2da70009d5141bf4344b59a0ce26cd8f7c5d5a28151a11f219",
		},
		"TokenBidWithdrawn": {
			TopicHash: "0xef9d84bc41d1a54361c4bf46f5e11b7c90a3fcb4e604b1b24e0e35d0fa371604",
		},
		"TokenBought": {
			TopicHash: "0x50a3cf2b1df7bd2994e752563ce6f85769fb50da66bbb9a9912d2d6066a6b4da",
		},
		"TokenDelisted": {
			TopicHash: "0x835a0a426c94e53ab00dd6cf617ba68d1fa6c9162ef7a036b80be930c9a32c53",
		},
		"TokenListed": {
			TopicHash: "0xfc8bd63d1c4480fdf7501f95b3dd0ba53542a02abbab92c0f579468341741abd",
		},
	},
	QuixoticTopic: map[string]Topic{
		"SellOrderFilled": {
			TopicHash: "0x70ba0d31158674eea8365d0f7b9ac70e552cc28b8bb848664e4feb939c6578f8",
		},
	},
}

type ChainId struct {
	Eth int64
	Ftm int64
	Opt int64
	Bsc int64
	Apt int64
}

type SaleVolumeAndFloorPriceTimeDuration string

var (
	SaleVolumeAndFloorPriceTimeDuration7Days    SaleVolumeAndFloorPriceTimeDuration = "7 Days"
	SaleVolumeAndFloorPriceTimeDuration30Days   SaleVolumeAndFloorPriceTimeDuration = "30 Days"
	SaleVolumeAndFloorPriceTimeDuration3Months  SaleVolumeAndFloorPriceTimeDuration = "3 Months"
	SaleVolumeAndFloorPriceTimeDuration12Months SaleVolumeAndFloorPriceTimeDuration = "12 Months"
)

func (t SaleVolumeAndFloorPriceTimeDuration) String() string {
	return string(t)
}

func (t SaleVolumeAndFloorPriceTimeDuration) IsValid() bool {
	return t == SaleVolumeAndFloorPriceTimeDuration7Days ||
		t == SaleVolumeAndFloorPriceTimeDuration30Days ||
		t == SaleVolumeAndFloorPriceTimeDuration3Months ||
		t == SaleVolumeAndFloorPriceTimeDuration12Months
}

func StringToSaleVolumeAndFloorPriceTimeDuration(str string) (SaleVolumeAndFloorPriceTimeDuration, bool) {
	duration, ok := map[string]SaleVolumeAndFloorPriceTimeDuration{
		"7d":  SaleVolumeAndFloorPriceTimeDuration7Days,
		"30d": SaleVolumeAndFloorPriceTimeDuration30Days,
		"3M":  SaleVolumeAndFloorPriceTimeDuration3Months,
		"12M": SaleVolumeAndFloorPriceTimeDuration12Months,
	}[str]

	return duration, ok
}

type TimeDuration string

var (
	TimeDuration24Hours TimeDuration = "24 hours"
	TimeDuration7Days   TimeDuration = "7 days"
	TimeDuration30Days  TimeDuration = "30 days"
	TimeDurationAllTime TimeDuration = "5 years"
)

func (t TimeDuration) String() string {
	return string(t)
}

func (t TimeDuration) IsValid() bool {
	return t == TimeDuration24Hours ||
		t == TimeDuration7Days ||
		t == TimeDuration30Days ||
		t == TimeDurationAllTime
}

func StringToTimeDuration(str string) (TimeDuration, bool) {
	duration, ok := map[string]TimeDuration{
		"24h": TimeDuration24Hours,
		"7d":  TimeDuration7Days,
		"30d": TimeDuration30Days,
		"all": TimeDurationAllTime,
	}[str]

	return duration, ok
}

type TimeInterval string

var (
	TimeInterval1Hour  TimeInterval = "1 hours"
	TimeInterval1Day   TimeInterval = "1 days"
	TimeInterval1Month TimeInterval = "1 months"
)

func (t TimeInterval) String() string {
	return string(t)
}

func (t TimeInterval) IsValid() bool {
	return t == TimeInterval1Hour ||
		t == TimeInterval1Day ||
		t == TimeInterval1Month
}

func StringToTimeInterval(str string) (TimeInterval, bool) {
	duration, ok := map[string]TimeInterval{
		"1h": TimeInterval1Hour,
		"1d": TimeInterval1Day,
		"1M": TimeInterval1Month,
	}[str]

	return duration, ok
}

type TimeRoundUp string

var (
	TimeRoundUpHours TimeRoundUp = "hours"
	TimeRoundUpDays  TimeRoundUp = "days"
	TimeRoundUpMonth TimeRoundUp = "months"
)

func (t TimeRoundUp) String() string {
	return string(t)
}

func (t TimeRoundUp) IsValid() bool {
	return t == TimeRoundUpHours ||
		t == TimeRoundUpDays ||
		t == TimeRoundUpMonth
}

func StringToTimeRoundUp(str string) (TimeRoundUp, bool) {
	duration, ok := map[string]TimeRoundUp{
		"h": TimeRoundUpHours,
		"d": TimeRoundUpDays,
		"M": TimeRoundUpMonth,
	}[str]

	return duration, ok
}

type Currency string

var (
	ETH Currency = "ETH"
	FTM Currency = "FTM"
)

func (c Currency) String() string {
	return string(c)
}

func (c Currency) IsValid() bool {
	return c == ETH ||
		c == FTM
}

func StringToCurrency(str string) (Currency, bool) {
	duration, ok := map[string]Currency{
		"eth": ETH,
		"ftm": FTM,
	}[str]

	return duration, ok
}
