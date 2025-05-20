package mcpTool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&ExChangeRate{})
}

type ExChangeRate struct {
}

type Currency struct {
	CurrencyCode string `json:"Currency Code"`
	CurrencyName string `json:"Currency Name"`
	Country      string `json:"Country"`
}

type ExchangeRateResponse struct {
	Result         string  `json:"result"`
	Documentation  string  `json:"documentation"`
	TermsOfUse     string  `json:"terms_of_use"`
	TimeLastUpdate int64   `json:"time_last_update_unix"`
	TimeLastUTC    string  `json:"time_last_update_utc"`
	TimeNextUpdate int64   `json:"time_next_update_unix"`
	TimeNextUTC    string  `json:"time_next_update_utc"`
	BaseCode       string  `json:"base_code"`
	TargetCode     string  `json:"target_code"`
	ConversionRate float64 `json:"conversion_rate"`
}

// 返回给前端的响应结构
type ExchangeRateResult struct {
	Rate                 float64 `json:"rate"`
	TimeLastUpdateUnix   int64   `json:"time_last_update_unix"`
	TimeNextUpdateUnix   int64   `json:"time_next_update_unix"`
	TimeLastUpdateFormat string  `json:"time_last_update_format"`
	TimeNextUpdateFormat string  `json:"time_next_update_format"`
}

// 硬编码的备用货币列表 - 仅在无法加载coinlist.json时使用
var fallbackCurrencyCodes = []string{
	"USD", "EUR", "CNY", "JPY", "GBP", "AUD", "CAD", "CHF", "HKD", "NZD",
	"SEK", "KRW", "SGD", "NOK", "MXN", "INR", "RUB", "ZAR", "TRY", "BRL",
	"TWD", "DKK", "PLN", "THB", "IDR", "HUF", "CZK", "ILS", "CLP", "PHP",
	"AED", "COP", "SAR", "MYR", "RON", "BGN", "HRK", "ISK",
}

// 加载货币列表
func loadCurrencyCodes() ([]string, error) {
	// 尝试不同的路径加载文件
	possiblePaths := []string{
		"./server/mcp/coinlist.json",
		"./mcp/coinlist.json",
		"./coinlist.json",
		"../mcp/coinlist.json",
		"../../mcp/coinlist.json",
	}

	var fileData []byte
	var err error
	var loadedPath string

	// 尝试所有可能的路径
	for _, path := range possiblePaths {
		if _, statErr := os.Stat(path); statErr == nil {
			if fileData, err = os.ReadFile(path); err == nil {
				loadedPath = path
				break
			}
		}
	}

	// 如果没有找到文件
	if fileData == nil {
		log.Printf("警告: 无法找到或读取货币列表文件，将使用内置的备用货币列表")
		return fallbackCurrencyCodes, errors.New("无法加载货币列表文件")
	}

	log.Printf("成功从 %s 加载货币列表文件", loadedPath)

	// 解析 JSON 数据
	var currencies []Currency
	if err := json.Unmarshal(fileData, &currencies); err != nil {
		log.Printf("解析货币列表JSON失败: %s", err.Error())
		return fallbackCurrencyCodes, err
	}

	// 提取货币代码
	currencyCodes := make([]string, len(currencies))
	for i, currency := range currencies {
		currencyCodes[i] = currency.CurrencyCode
	}

	log.Printf("成功加载 %d 个货币代码", len(currencyCodes))
	return currencyCodes, nil
}

// 格式化Unix时间戳为易读格式
func formatUnixTime(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	return t.Format("2006-01-02 15:04:05")
}

// 汇率查询工具
func (t *ExChangeRate) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 1. 参数验证
	from, ok := request.Params.Arguments["from"].(string)
	if !ok {
		return nil, errors.New("参数错误：from 必须是字符串类型")
	}

	to, ok := request.Params.Arguments["to"].(string)
	if !ok {
		return nil, errors.New("参数错误：to 必须是字符串类型")
	}

	// 2. 记录操作日志
	global.GVA_LOG.Info(fmt.Sprintf("汇率查询工具被调用：from=%s, to=%s", from, to))

	// 3. 调用汇率API
	apiKey := os.Getenv("EXCHANGE_RATE_API_KEY") // 替换为实际的API密钥
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/pair/%s/%s", apiKey, from, to)

	// 发起HTTP请求
	resp, err := http.Get(url)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("汇率API请求失败: %s", err.Error()))
		return nil, errors.New("汇率API请求失败")
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("读取API响应失败: %s", err.Error()))
		return nil, errors.New("读取API响应失败")
	}

	// 解析JSON响应
	var rateResp ExchangeRateResponse
	if err := json.Unmarshal(body, &rateResp); err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("解析API响应失败: %s", err.Error()))
		return nil, errors.New("解析API响应失败")
	}

	// 检查API响应结果
	if rateResp.Result != "success" {
		global.GVA_LOG.Error(fmt.Sprintf("API返回错误: %s", string(body)))
		return nil, errors.New("汇率查询失败")
	}

	// 4. 构造返回结果
	percentage := rateResp.ConversionRate

	// 创建新的结果结构，包含汇率和时间戳
	result := ExchangeRateResult{
		Rate:                 percentage,
		TimeLastUpdateUnix:   rateResp.TimeLastUpdate,
		TimeNextUpdateUnix:   rateResp.TimeNextUpdate,
		TimeLastUpdateFormat: formatUnixTime(rateResp.TimeLastUpdate),
		TimeNextUpdateFormat: formatUnixTime(rateResp.TimeNextUpdate),
	}

	// 将结果转为JSON字符串
	resultJSON, err := json.Marshal(result)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("序列化结果失败: %s", err.Error()))
		return nil, errors.New("序列化结果失败")
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: string(resultJSON),
			},
		},
	}, nil
}

func (t *ExChangeRate) New() mcp.Tool {
	// 加载货币代码列表 - 即使出错也会返回备用货币列表
	currencyCodes, _ := loadCurrencyCodes()

	// 即使loadCurrencyCodes出错也会返回fallbackCurrencyCodes，所以这里不需要再次检查错误

	return mcp.NewTool("exchangeRate",
		mcp.WithDescription("汇率查询工具"),
		mcp.WithString("from", mcp.Required(),
			mcp.Description("当前货币，比如：USD"),
			mcp.Enum(currencyCodes...),
		),
		mcp.WithString("to", mcp.Required(),
			mcp.Description("目标货币，比如：CNY"),
			mcp.Enum(currencyCodes...),
		),
	)
}
