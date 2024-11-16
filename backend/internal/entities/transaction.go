package entities

import (
	"fmt"
	"time"
)

type Transaction struct {
	Hash      string    `json:"hash"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

func (t Transaction) String() string {
	return fmt.Sprintf("[from: %s, to: %s, value: %s timestamp:%s]",
		t.From, t.To, t.Value, t.Timestamp)
}

type TransactionList struct {
	List  []*Transaction `json:"transactions"`
	Graph *Graph         `json:"graph"`
}

func NewTransactionList(l []*Transaction) *TransactionList {
	return &TransactionList{
		List: l,
	}
}

func (l *TransactionList) GetTargets() []string {
	targets := make([]string, 0)
	for _, tx := range l.List {
		targets = append(targets, tx.To)
	}
	return targets
}
