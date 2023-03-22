package config

import (
	"github.com/spf13/viper"
)

// Loader load config from reader into Viper
type Loader interface {
	Load(viper.Viper) (*viper.Viper, error)
}

type Config struct {
	// db
	Postgres   DBConnection
	Clickhouse DBConnection

	// server
	ApiServer  ApiServer
	GrpcServer GrpcServer
	Kafka      Kafka

	// blockchain
	RpcUrl              RpcUrl
	ChainExplorerApiKey ChainExplorerApiKey
	MarketplaceApiKey   MarketplaceApiKey

	// cloud service
	GoogleServiceAccountKey string
	GoogleNftImageBucket    string
	Debug                   bool

	// ipfs server
	IpfsServer string

	// coingecko
	CoingeckoApiKey string
}

type DBConnection struct {
	Host string
	Port string
	User string
	Name string
	Pass string

	SSLMode string
}

type ApiServer struct {
	Port           string
	AllowedOrigins string
}

type GrpcServer struct {
	Port string
	Host string
}

type ChainExplorerApiKey struct {
	Eth      string
	Ftm      string
	Bsc      string
	Optimism string
}
type RpcUrl struct {
	Eth      string
	Ftm      string
	Optimism string
}

type MarketplaceApiKey struct {
	Opensea  string
	Quixotic string
}

type Kafka struct {
	Servers               string
	IndexerTopic          string
	ConsumerGroup         string
	EvmConsumerGroup      string
	AptosConsumerGroup    string
	RoninConsumerGroup    string
	SolanaConsumerGroup   string
	PriorityConsumerGroup string
	SuiConsumerGroup      string
	OnusConsumerGroup     string
}

func generateConfigFromViper(v *viper.Viper) *Config {
	return &Config{
		Debug: v.GetBool("DEBUG"),

		GrpcServer: GrpcServer{
			Port: v.GetString("GRPC_PORT"),
			Host: v.GetString("GRPC_HOST"),
		},

		ApiServer: ApiServer{
			Port:           v.GetString("PORT"),
			AllowedOrigins: v.GetString("ALLOWED_ORIGINS"),
		},

		Kafka: Kafka{
			Servers:               v.GetString("KAFKA_SERVERS"),
			IndexerTopic:          v.GetString("KAFKA_INDEXER_TOPIC"),
			ConsumerGroup:         v.GetString("KAFKA_CONSUMER_GROUP"),
			EvmConsumerGroup:      v.GetString("KAFKA_EVM_CONSUMER_GROUP"),
			AptosConsumerGroup:    v.GetString("KAFKA_APTOS_CONSUMER_GROUP"),
			RoninConsumerGroup:    v.GetString("KAFKA_RONIN_CONSUMER_GROUP"),
			SolanaConsumerGroup:   v.GetString("KAFKA_SOL_CONSUMER_GROUP"),
			PriorityConsumerGroup: v.GetString("KAFKA_PRIORITY_CONSUMER_GROUP"),
			SuiConsumerGroup:      v.GetString("KAFKA_SUI_CONSUMER_GROUP"),
			OnusConsumerGroup:     v.GetString("KAFKA_ONUS_CONSUMER_GROUP"),
		},

		GoogleServiceAccountKey: v.GetString("GOOGLE_SERVICE_ACCOUNT_KEY"),
		GoogleNftImageBucket:    v.GetString("GOOGLE_NFT_IMAGE_BUCKET"),

		Clickhouse: DBConnection{
			Host: v.GetString("CLICKHOUSE_HOST"),
			Port: v.GetString("CLICKHOUSE_PORT"),
			User: v.GetString("CLICKHOUSE_USER"),
			Name: v.GetString("CLICKHOUSE_NAME"),
			Pass: v.GetString("CLICKHOUSE_PASS"),
		},

		Postgres: DBConnection{
			Host:    v.GetString("DB_HOST"),
			Port:    v.GetString("DB_PORT"),
			User:    v.GetString("DB_USER"),
			Name:    v.GetString("DB_NAME"),
			Pass:    v.GetString("DB_PASS"),
			SSLMode: v.GetString("DB_SSL_MODE"),
		},

		RpcUrl: RpcUrl{
			Eth:      v.GetString("ETH_RPC"),
			Ftm:      v.GetString("FTM_RPC"),
			Optimism: v.GetString("OPTIMISM_RPC"),
		},

		ChainExplorerApiKey: ChainExplorerApiKey{
			Eth:      v.GetString("ETHERSCAN_API_KEY"),
			Ftm:      v.GetString("FTMSCAN_API_KEY"),
			Bsc:      v.GetString("BSCSCAN_API_KEY"),
			Optimism: v.GetString("OPTIMISM_API_KEY"),
		},

		MarketplaceApiKey: MarketplaceApiKey{
			Opensea:  v.GetString("OPENSEA_API_KEY"),
			Quixotic: v.GetString("QUIXOTIC_API_KEY"),
		},

		IpfsServer:      v.GetString("IPFS_SERVER"),
		CoingeckoApiKey: v.GetString("COINGECKO_API_KEY"),
	}
}

func DefaultConfigLoaders() []Loader {
	loaders := []Loader{}
	fileLoader := NewFileLoader(".env", ".")
	loaders = append(loaders, fileLoader)
	loaders = append(loaders, NewENVLoader())

	return loaders
}

// LoadConfig load config from loader list
func LoadConfig(loaders []Loader) *Config {
	v := viper.New()
	v.SetDefault("PORT", "8080")
	v.SetDefault("GRPC_PORT", "8081")
	v.SetDefault("ENV", "local")
	v.SetDefault("FTM_RPC", "https://rpc.fantom.network")
	v.SetDefault("ALLOWED_ORIGINS", "*")

	for idx := range loaders {
		newV, err := loaders[idx].Load(*v)

		if err == nil {
			v = newV
		}
	}
	return generateConfigFromViper(v)
}

func LoadTestConfig() Config {
	return Config{
		Debug: true,
		ApiServer: ApiServer{
			Port: "8080",
		},
		GrpcServer: GrpcServer{
			Host: "grpc_host",
			Port: "8081",
		},
		Postgres: DBConnection{
			Host:    "127.0.0.1",
			Port:    "15432",
			User:    "postgres",
			Pass:    "postgres",
			Name:    "indexer_local",
			SSLMode: "disable",
		},
		Clickhouse: DBConnection{
			Host: "clickhouse_host",
			Port: "clickhouse_port",
			User: "clickhouse_user",
			Pass: "clickhouse_pass",
			Name: "clickhouse_name",
		},
		ChainExplorerApiKey: ChainExplorerApiKey{
			Eth: "eth",
			Ftm: "ftm",
			Bsc: "bsc",
		},
		RpcUrl: RpcUrl{
			Eth: "eth_rpc",
			Ftm: "ftm_rpc",
		},
		IpfsServer: "ipfs.io",
	}
}
