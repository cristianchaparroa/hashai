package entities

import (
	"fmt"
	"strconv"
	"strings"
)

type GraphNode struct {
	Name string `json:"name"`
}

// Edge represents a processed edge with combined values
type Edge struct {
	Source string
	Target string
	Value  float64
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

// ToMermaid converts a Graph structure to Mermaid diagram format
func (g *Graph) ToMermaid() string {
	var mermaidLines []string
	mermaidLines = append(mermaidLines, "graph LR")

	// Create a mapping of index to node names
	nodeMap := make(map[int]string)
	for i, node := range g.Nodes {
		nodeMap[i] = node.Name
	}

	// Process edges and combine values for same source-target pairs
	processedEdges := make(map[string]*Edge)

	for _, link := range g.Links {
		source := nodeMap[link.Source]
		target := nodeMap[link.Target]

		// Skip self-referential edges
		if source == target {
			continue
		}

		// Create a unique key for this edge
		edgeKey := fmt.Sprintf("%s-%s", source, target)

		if edge, exists := processedEdges[edgeKey]; exists {
			edge.Value += link.Value
		} else {
			processedEdges[edgeKey] = &Edge{
				Source: source,
				Target: target,
				Value:  link.Value,
			}
		}
	}

	// Convert edges to Mermaid format
	for _, edge := range processedEdges {
		if edge.Value > 0 { // Only show edges with positive values
			source := shortenAddress(edge.Source)
			target := shortenAddress(edge.Target)
			value := formatValue(edge.Value)

			// Format the edge with the value as a label
			edgeLine := fmt.Sprintf("    %s-->|%s|%s", source, value, target)
			mermaidLines = append(mermaidLines, edgeLine)
		}
	}

	return strings.Join(mermaidLines, "\n")
}
