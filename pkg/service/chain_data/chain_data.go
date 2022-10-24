package chaindata

import (
	"database/sql"

	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/model"
)

type Service struct {
	db *sql.DB
}

func New(db *sql.DB) IService {
	return &Service{
		db: db,
	}
}

func (s Service) GetEventsByAddressFromStartBlockToBlock(address string, start int64, to int64, chainId int64, ch chan *model.Event) (err error) {
	defer logger.L.Info("CLOSE CHANNEL")
	defer close(ch)
	var size int64 = 20000
	var offset int64 = 0
	// to optimize this query, we query by 20k block each
	for {
		logger.L.Fields(logger.Fields{
			"address": address,
			"chainId": chainId,
		}).Info("Querying events from clickhouse")
		rows, err := s.db.Query("SELECT distinct transaction_hash, block_number, data, topics FROM data.evm_events WHERE address = $1 AND block_number > $2 AND block_number <= $3 AND  chain_id = $4 ORDER BY block_number ASC, log_index ASC LIMIT $5 OFFSET $6", address, start, to, chainId, size, offset)
		defer func() {
			if rows != nil {
				rows.Close()
			}
		}()

		if err != nil {
			return err
		}

		var i int64 = 0
		for rows.Next() {
			var (
				transactionHash string
				blockNumber     int64
				data            []byte
				topics          []string
			)
			if err := rows.Scan(&transactionHash, &blockNumber, &data, &topics); err != nil {
				return err
			}
			i++
			ch <- &model.Event{
				ChainId:         chainId,
				TransactionHash: transactionHash,
				BlockNumber:     blockNumber,
				Address:         address,
				Data:            data,
				Topics:          topics,
			}
		}

		err = rows.Err()
		if err != nil {
			return err
		}
		if i < size {
			break
		} else {
			logger.L.Fields(logger.Fields{
				"contractAddress": address,
				"start":           start,
				"to":              to,
				"chainId":         chainId,
				"size":            size,
				"offset":          offset,
			}).Info("query next batch")
		}
		offset += size
	}
	return nil
}

func (s Service) GetTipOfChain() (*model.TipOfChain, error) {
	resp := &model.TipOfChain{}
	rows, err := s.db.Query("SELECT chain_id, latest_block FROM latest_block")
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		return nil, err
	}

	tip := make(map[int64]int64)
	for rows.Next() {
		var (
			chainId     int64
			latestBlock int64
		)
		if err := rows.Scan(&chainId, &latestBlock); err != nil {
			return nil, err
		}
		tip[chainId] = latestBlock
	}

	resp.BlockHeadNumber = tip
	return resp, rows.Err()
}

func (s Service) GetBlockStatsByChainId(chainId int64) (*model.BlockStats, error) {
	resp := &model.BlockStats{}
	rows, err := s.db.Query(`
		SELECT * FROM block_stats WHERE chain_id = $1
	`, chainId)

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			chainId         int64
			firstBlock      int64
			lastBlock       int64
			totalBlocks     int64
			actualBlocks    uint64
			remainingBlocks int64
			percentage      float64
		)
		if err := rows.Scan(&chainId, &firstBlock, &lastBlock, &totalBlocks, &actualBlocks, &remainingBlocks, &percentage); err != nil {
			return nil, err
		}
		resp = &model.BlockStats{
			ChainId:         chainId,
			FirstBlock:      firstBlock,
			LastBlock:       lastBlock,
			TotalBlocks:     totalBlocks,
			ActualBlocks:    actualBlocks,
			RemainingBlocks: remainingBlocks,
			Percentage:      percentage,
		}
	}
	return resp, rows.Err()
}
