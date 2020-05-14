package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {

	return &Block{
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
		previousHash: previousHash,
		transactions: transactions,
	}
}

func (b *Block) Print() {

	fmt.Printf("timestamp      %d\n", b.timestamp)
	fmt.Printf("nonce          %d\n", b.nonce)
	fmt.Printf("previousHash   %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}

}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Printf(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json: timestamp`
		Nonce        int            `json: nonce`
		PreviousHash [32]byte       `json: previous_hash`
		Transactions []*Transaction `json: transactions`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

type Blockchain struct {
	transactionsPool []*Transaction
	chain            []*Block
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionsPool)
	bc.chain = append(bc.chain, b)
	bc.transactionsPool = []*Transaction{}
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%sChain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {

	t := NewTransaction(sender, recipient, value)
	bc.transactionsPool = append(bc.transactionsPool, t)
}

type Transaction struct {
	senderBlockAddress    string
	recipientBlockAddress string
	value                 float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{senderBlockAddress: sender, recipientBlockAddress: recipient, value: value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("=", 40))
	fmt.Printf(" sender_blockchain_adress      %s\n", t.senderBlockAddress)
	fmt.Printf(" recipient_blockchain_adress   %s\n", t.recipientBlockAddress)
	fmt.Printf(" value                         %.1f\n", t.value)
}

func (t *Transaction) Marshal() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json: sender_block_address`
		Recipient string  `json: recipient_block_address`
		Value     float32 `json: value`
	}{
		Sender:    t.senderBlockAddress,
		Recipient: t.recipientBlockAddress,
		Value:     t.value,
	})
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchain := NewBlockchain()
	//blockchain.Print()

	blockchain.AddTransaction("A", "V", 1.0)

	previousHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(5, previousHash)
	blockchain.Print()

	blockchain.AddTransaction("A", "BB", 5.0)
	blockchain.AddTransaction("BB", "V", 9.0)
	previousHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(6, previousHash)
	blockchain.Print()

}
