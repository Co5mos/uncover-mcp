package main

import (
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"github.com/projectdiscovery/uncover"
	"github.com/projectdiscovery/uncover-mcp/sources/fofa"
	"github.com/projectdiscovery/uncover-mcp/sources/shodan"
)

func main() {
	// Initialize gologger
	gologger.DefaultLogger.SetMaxLevel(levels.LevelInfo)

	// Create a new MCP server
	s := server.NewMCPServer(
		"Uncover MCP",
		"0.0.1",
		server.WithLogging(),
	)

	// Add Shodan tool
	shodanTool := mcp.NewTool("shodan",
		mcp.WithDescription("Use Shodan search engine to find exposed hosts"),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("Shodan search query"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Limit the number of results"),
		),
		mcp.WithString("field",
			mcp.Description("Return field format (ip:port, host, ip, port)"),
		),
	)
	s.AddTool(shodanTool, shodan.Handler)

	// Add FOFA tool
	fofaTool := mcp.NewTool("fofa",
		mcp.WithDescription("Use FOFA search engine to find exposed hosts"),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("FOFA search query"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Limit the number of results"),
		),
		mcp.WithString("field",
			mcp.Description("Return field format (ip:port, host, ip, port)"),
		),
	)
	s.AddTool(fofaTool, fofa.Handler)

	// Display available uncover engines
	u, err := uncover.New(&uncover.Options{})
	if err == nil {
		agents := u.AllAgents()
		gologger.Info().Msgf("Available uncover search engines: %v", agents)
	}

	// Start the server
	fmt.Println("Starting Uncover MCP service...")
	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}
