package entities

import "strconv"

type GraphNode struct {
	Name string `json:"name"`
}

type GraphLink struct {
	Source int     `json:"source"`
	Target int     `json:"target"`
	Value  float64 `json:"value"`
}

type Graph struct {
	Nodes []GraphNode `json:"nodes"`
	Links []GraphLink `json:"links"`
}

func NewSankeyGraph(transactions []*Transaction) *Graph {
	addresses := make(map[string]int)
	links := make([]GraphLink, 0)

	// First pass: create nodes
	index := 0
	for _, tx := range transactions {
		_, existFrom := addresses[tx.From]
		if !existFrom {
			addresses[tx.From] = index
			index++
		}

		_, existTo := addresses[tx.To]
		if !existTo {
			addresses[tx.To] = index
			index++
		}
	}

	nodes := make([]GraphNode, 0)
	for hash := range addresses {
		node := GraphNode{
			Name: hash,
		}
		nodes = append(nodes, node)
	}

	// Second pass: create links
	for _, tx := range transactions {
		if tx.From == tx.To {
			continue
		}

		amount, _ := strconv.ParseFloat(tx.Value, 64)
		sourceIndex := addresses[tx.From]
		targetIndex := addresses[tx.To]

		// Create link even if source and target are the same
		link := GraphLink{
			Source: sourceIndex,
			Target: targetIndex,
			Value:  amount,
		}
		links = append(links, link)
	}

	return &Graph{
		Nodes: nodes,
		Links: links,
	}
}
