# Uncover MCP

Uncover MCP是基于[uncover](https://github.com/projectdiscovery/uncover)工具的[MCP (Model Context Protocol)](https://github.com/mark3labs/mcp-go) 服务实现，用于快速发现互联网上暴露的主机。

## 功能特点

- 直接使用projectdiscovery/uncover库进行搜索
- 支持多个搜索引擎（当前实现Shodan和FOFA）
- 支持多种结果格式输出
- 简单易用的MCP服务接口
- 与AI助手自然交互

## 安装

```bash
# 克隆仓库
git clone https://github.com/Co5mos/uncover-mcp
cd uncover-mcp

# 构建项目
go build -o uncover-mcp ./cmd/uncover-mcp
```

## 使用方法

作为MCP服务运行：

```bash
{
    "mcpServers": {
        "uncover-mcp": {
            "command": "uncover-mcp.exe",
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

## 示例

在AI助手中使用如下格式调用：

```
查找使用Shodan搜索ssl:"Uber Technologies, Inc."
```

```
使用FOFA搜索app="ATLASSIAN-JIRA"
```

## 声明

本工具仅供安全研究使用，使用本工具时请遵守相关法律法规和uncover项目的使用条款。

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件
