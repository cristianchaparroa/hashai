package usecases

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/puzpuzpuz/xsync/v3"
	"hashtracker/internal/entities"
	"net/url"
	"sync"
)

type scanner struct {
	repo         ETHTransactionRepository
	ensRepo      ENSRepository
	ensValidator ENSValidator
}

func NewScanner(
	repo ETHTransactionRepository,
	ensRepo ENSRepository,
	ensValidator ENSValidator,
) ScannerUseCase {
	return &scanner{
		repo:         repo,
		ensValidator: ensValidator,
		ensRepo:      ensRepo,
	}
}

func (s *scanner) Scan(ctx context.Context, address string, level int) (*entities.TransactionList, error) {
	isValidENS := s.ensValidator.IsValid(ctx, address)
	fmt.Println("--> Is valid ENS: ", isValidENS)
	if isValidENS {
		result, err := s.ensRepo.Resolve(ctx, address)
		if err != nil {
			return nil, err
		}
		address = result.Address
	}

	return s.GetTransactions(ctx, address, level)
}

func (s *scanner) GetTransactions(ctx context.Context, address string, level int) (*entities.TransactionList, error) {
	fmt.Println("----> GetTransactions: ", address)
	transactions, getErr := s.repo.GetTransactions(ctx, address)
	if getErr != nil {
		return nil, getErr
	}
	m := xsync.NewMapOf[string, *entities.TransactionList]()
	m.Store(address, transactions)

	var wg sync.WaitGroup
	for i := 1; i < level; i++ {
		targets := transactions.GetTargets()
		for _, addr := range targets {
			wg.Add(1)
			go func(wg *sync.WaitGroup, m *xsync.MapOf[string, *entities.TransactionList], address string) {
				defer wg.Done()
				ts, getTxErr := s.repo.GetOutTransactions(ctx, address)
				if getTxErr != nil {
					return
				}
				m.Store(address, ts)
			}(&wg, m, addr)
		}
	}
	wg.Wait()

	txs := make([]*entities.Transaction, 0)
	m.Range(func(key string, l *entities.TransactionList) bool {
		txs = append(txs, l.List...)
		return true
	})
	g := entities.NewSankeyGraph(txs)

	urlEncoded := url.QueryEscape(g.ToMermaid())
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(urlEncoded))

	return &entities.TransactionList{
		List:       txs,
		Graph:      g,
		MermaidB64: base64Encoded,
	}, nil
}
