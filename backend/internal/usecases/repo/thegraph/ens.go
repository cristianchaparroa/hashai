package thegraph

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"hashtracker/internal/entities/thegraph"
	"hashtracker/internal/usecases"
	"io"
	"net/http"
	"strings"
	"time"
)

const getENSQuery = `
	query GetENSAddress($name: String!) {
		domains(where: { name: $name }) {
			id
			name
			resolvedAddress {
				id
			}
			owner {
				id
			}
			resolver {
				address
			}
		}
	}`

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type ENSQueryVariables struct {
	Name string `json:"name"`
}

type ensRepository struct {
	apiKey     string
	httpClient http.Client // TODO: point this to an interfaces
}

func NewENSRepository(apiKey string) usecases.ENSRepository {
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	return &ensRepository{
		apiKey:     apiKey,
		httpClient: client,
	}
}

func (e *ensRepository) Resolve(ctx context.Context, ensName string) (*thegraph.ENSResponse, error) {
	// Construct the URL with the API key
	endpoint := fmt.Sprintf("https://api.thegraph.com/subgraphs/name/ensdomains/ens")

	variables := ENSQueryVariables{
		Name: strings.ToLower(ensName), // ENS names are case insensitive
	}
	request := GraphQLRequest{
		Query: getENSQuery,
		Variables: map[string]interface{}{
			"name": variables.Name,
		},
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	body := bytes.NewBuffer(jsonBody)
	req, err := http.NewRequest(http.MethodPost, endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", e.apiKey))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := e.httpClient.Do(req)
	if err != nil {
		return nil, nil
	}
	defer res.Body.Close()

	var response *ENSResponse
	bs, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	unmarshalErr := json.Unmarshal(bs, &response)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	if response == nil {
		return nil, fmt.Errorf("no result found for address: %s", ensName)
	}

	if len(response.Data.Domains) == 0 {
		return nil, fmt.Errorf("there is no record for the ENS: %s", ensName)
	}

	return &thegraph.ENSResponse{
		Address: response.Data.Domains[0].ResolvedAddress.ID,
	}, nil
}
