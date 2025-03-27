# Uncover MCP

Uncover MCP is a [MCP (Model Context Protocol)](https://github.com/mark3labs/mcp-go) service implementation based on the [uncover](https://github.com/projectdiscovery/uncover) tool, designed for quickly discovering exposed hosts on the internet.

<h3>Engilsh | <a href='./README_CN.md'> 中文 </a></h3>

## Features

- Direct integration with projectdiscovery/uncover library for searching
- Support for multiple search engines (currently Shodan and FOFA)
- Multiple output format options
- Simple and easy-to-use MCP service interface
- Natural interaction with AI assistants

## Installation

```bash
# Clone repository
git clone https://github.com/Co5mos/uncover-mcp
cd uncover-mcp

# Build project
go build -o uncover-mcp ./cmd/uncover-mcp
```

## Usage

Run as an MCP service:

```bash
{
    "mcpServers": {
        "uncover-mcp": {
            "command": "./uncover-mcp",
            "args": [],
            "env": {
                "SHODAN_API_KEY": "",
                "FOFA_EMAIL": "",
                "FOFA_KEY": ""
            }
        }
    }
}
```

## Examples

Use the following format in AI assistant:

```
Search using Shodan for ssl:"Uber Technologies, Inc."
```

```
Search using FOFA for app="ATLASSIAN-JIRA"
```

## Declaration

This tool is intended for security research purposes only. Use this tool in compliance with relevant laws and regulations and the terms of use of the uncover project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file
