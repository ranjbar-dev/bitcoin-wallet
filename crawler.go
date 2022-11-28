package bitcoinWallet

import (
	"errors"
	"fmt"
	"github.com/Amirilidan78/bitcoin-wallet/blockDaemon"
	"github.com/Amirilidan78/bitcoin-wallet/blockDaemon/response"
	"github.com/Amirilidan78/bitcoin-wallet/enums"
	"sync"
	"time"
)

type Crawler struct {
	Node      enums.Node
	Addresses []string
}

type CrawlResult struct {
	Address      string
	Transactions []CrawlTransaction
}

type CrawlTransaction struct {
	TxId          string
	Confirmations int64
	FromAddress   string
	ToAddress     string
	Amount        uint64
	Symbol        string
}

func (c *Crawler) client() blockDaemon.BlockDaemon {
	return blockDaemon.NewBlockDaemonService(c.Node.Config)
}

func (c *Crawler) ScanBlocks(count int) ([]CrawlResult, error) {

	var wg sync.WaitGroup

	var allTransactions [][]CrawlTransaction

	client := c.client()

	number, err := client.CurrentBlockNumber()
	if err != nil {
		return nil, err
	}

	blockNumber := int64(number)

	go c.getBlockData(&wg, client, &allTransactions, blockNumber)

	for i := count; i > 0; i-- {
		wg.Add(1)
		blockNumber = blockNumber - 1
		// sleep to avoid 503 error
		time.Sleep(100 * time.Millisecond)
		go c.getBlockData(&wg, client, &allTransactions, blockNumber)
	}

	wg.Wait()

	return c.prepareCrawlResultFromTransactions(allTransactions), nil
}

func (c *Crawler) ScanBlocksFromTo(from int, to int) ([]CrawlResult, error) {

	if to-from < 1 {
		return nil, errors.New("to number should be more than from number")
	}

	client := c.client()

	var wg sync.WaitGroup

	var allTransactions [][]CrawlTransaction

	blockNumber := int64(to)

	for i := int(blockNumber); i > from; i-- {
		wg.Add(1)
		// sleep to avoid 503 error
		time.Sleep(100 * time.Millisecond)
		go c.getBlockData(&wg, client, &allTransactions, int64(i))
	}

	wg.Wait()

	return c.prepareCrawlResultFromTransactions(allTransactions), nil
}

func (c *Crawler) getBlockData(wg *sync.WaitGroup, client blockDaemon.BlockDaemon, allTransactions *[][]CrawlTransaction, num int64) {

	defer wg.Done()

	block, err := client.BlockByNumber(num)
	if err != nil {
		fmt.Println(err)
		return
	}

	// check block for transaction
	*allTransactions = append(*allTransactions, c.extractOurTransactionsFromBlock(block))
}

func (c *Crawler) extractOurTransactionsFromBlock(block response.BlockResponse) []CrawlTransaction {

	var txs []CrawlTransaction

	for _, transaction := range block.Txs {

		symbol := "BTC"

		fromAddress := ""
		for _, item := range c.getTxInputs(transaction) {
			fromAddress = item.Source
		}

		toAddress := ""
		amount := 0
		for _, item := range c.getTxOutputs(transaction) {
			for _, ourAddress := range c.Addresses {
				if ourAddress == item.Destination {
					toAddress = item.Destination
					amount = item.Amount
				}
			}
		}

		txId := transaction.Id
		confirmations := transaction.Confirmations

		for _, ourAddress := range c.Addresses {
			if ourAddress == toAddress || ourAddress == fromAddress {
				txs = append(txs, CrawlTransaction{
					TxId:          txId,
					FromAddress:   fromAddress,
					ToAddress:     toAddress,
					Amount:        uint64(amount),
					Confirmations: int64(confirmations),
					Symbol:        symbol,
				})
			}
		}
	}

	return txs
}

func (c *Crawler) prepareCrawlResultFromTransactions(transactions [][]CrawlTransaction) []CrawlResult {

	var result []CrawlResult

	for _, transaction := range transactions {
		for _, tx := range transaction {

			if c.addressExistInResult(result, tx.ToAddress) {
				id, res := c.getAddressCrawlInResultList(result, tx.ToAddress)
				res.Transactions = append(res.Transactions, tx)
				result[id] = res

			} else {
				result = append(result, CrawlResult{
					Address:      tx.ToAddress,
					Transactions: []CrawlTransaction{tx},
				})
			}
		}
	}

	return result
}

func (c *Crawler) addressExistInResult(result []CrawlResult, address string) bool {
	for _, res := range result {
		if res.Address == address {
			return true
		}
	}
	return false
}

func (c *Crawler) getAddressCrawlInResultList(result []CrawlResult, address string) (int, CrawlResult) {
	for id, res := range result {
		if res.Address == address {
			return id, res
		}
	}
	panic("crawl result not found")
}

func (c *Crawler) getTxInputs(tx response.Transaction) []response.TransactionEvent {

	var records []response.TransactionEvent

	for _, event := range tx.Events {
		if event.Type == "utxo_input" {
			records = append(records, event)
		}
	}

	return records
}

func (c *Crawler) getTxOutputs(tx response.Transaction) []response.TransactionEvent {

	var records []response.TransactionEvent

	for _, event := range tx.Events {
		if event.Type == "utxo_output" {
			records = append(records, event)
		}
	}

	return records
}
