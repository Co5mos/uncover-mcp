package fofa

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/projectdiscovery/uncover"
	"github.com/projectdiscovery/uncover/sources"
)

// Handler processes FOFA search requests from MCP
func Handler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract parameters from request
	query, ok := request.Params.Arguments["query"].(string)
	if !ok {
		return nil, fmt.Errorf("query parameter must be a string")
	}

	limit := 100 // default value
	if limitArg, ok := request.Params.Arguments["limit"]; ok {
		switch v := limitArg.(type) {
		case float64:
			limit = int(v)
		case int:
			limit = v
		}
	}

	// Create uncover options
	opts := uncover.Options{
		Agents:   []string{"fofa"},
		Queries:  []string{query},
		Limit:    limit,
		MaxRetry: 2,
		Timeout:  30,
	}

	// Initialize uncover client
	client, err := uncover.New(&opts)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize uncover client: %v", err)
	}

	// Collect results
	var results []string
	resultCallback := func(result sources.Result) {
		results = append(results, result.IpPort())
	}

	// Execute query
	if err := client.ExecuteWithCallback(ctx, resultCallback); err != nil {
		return nil, fmt.Errorf("failed to execute FOFA query: %v", err)
	}

	// Return results
	resultText := strings.Join(results, "\n")
	return mcp.NewToolResultText(resultText), nil
}
