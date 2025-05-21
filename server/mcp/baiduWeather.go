package mcpTool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool(&BaiduWeatherTool{})
}

// BaiduWeather 百度天气数据模型
type BaiduWeather struct {
	ID              uint   `json:"id" gorm:"primarykey"`                            // 主键ID
	DistrictID      string `json:"district_id" gorm:"column:district_id"`           // 区县ID
	Province        string `json:"province" gorm:"column:province"`                 // 省份
	City            string `json:"city" gorm:"column:city"`                         // 城市
	CityGeocode     string `json:"city_geocode" gorm:"column:city_geocode"`         // 城市地理编码
	District        string `json:"district" gorm:"column:district"`                 // 区县
	DistrictGeocode string `json:"district_geocode" gorm:"column:district_geocode"` // 区县地理编码
	Lon             string `json:"lon" gorm:"column:lon"`                           // 经度
	Lat             string `json:"lat" gorm:"column:lat"`                           // 纬度
}

// TableName 设置表名
func (BaiduWeather) TableName() string {
	return "baidu_weather"
}

type BaiduWeatherAbroad struct {
	ID             uint   `json:"id" gorm:"primarykey"`                            // 主键ID
	DistrictID     string `json:"district_id" gorm:"column:district_id"`           // 区县ID
	CountryEng     string `json:"country_eng" gorm:"column:country_eng"`           // 英文国家名
	Level1Eng      string `json:"level1_eng" gorm:"column:level1_eng"`             // 英文一级行政区
	Level2Eng      string `json:"level2_eng" gorm:"column:level2_eng"`             // 英文二级行政区
	Level3Eng      string `json:"level3_eng" gorm:"column:level3_eng"`             // 英文三级行政区
	CountryChn     string `json:"country_chn" gorm:"column:country_chn"`           // 中文国家名
	Level1Chn      string `json:"level1_chn" gorm:"column:level1_chn"`             // 中文一级行政区
	Level2Chn      string `json:"level2_chn" gorm:"column:level2_chn"`             // 中文二级行政区
	Level3Chn      string `json:"level3_chn" gorm:"column:level3_chn"`             // 中文三级行政区
	IsoCountryCode string `json:"iso_country_code" gorm:"column:iso_country_code"` // ISO国家代码
	Lon            string `json:"lon" gorm:"column:lon"`                           // 经度
	Lat            string `json:"lat" gorm:"column:lat"`                           // 纬度
	Capital        bool   `json:"capital" gorm:"column:capital"`                   // 是否首府
}

func (BaiduWeatherAbroad) TableName() string {
	return "baidu_weather_abroad"
}

// BaiduWeatherResponse 百度天气API响应数据结构
type BaiduWeatherResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  struct {
		Location struct {
			Country  string `json:"country"`
			Province string `json:"province"`
			City     string `json:"city"`
			Name     string `json:"name"`
			ID       string `json:"id"`
		} `json:"location"`
		Now struct {
			Temp      int     `json:"temp"`       // 温度
			FeelsLike int     `json:"feels_like"` // 体感温度
			RH        int     `json:"rh"`         // 相对湿度
			WindClass string  `json:"wind_class"` // 风力等级
			WindDir   string  `json:"wind_dir"`   // 风向
			Text      string  `json:"text"`       // 天气状况文字
			Prec1h    float64 `json:"prec_1h"`    // 1小时降水量
			Clouds    int     `json:"clouds"`     // 云量
			Vis       int     `json:"vis"`        // 能见度
			AQI       int     `json:"aqi"`        // 空气质量指数
			PM25      int     `json:"pm25"`       // PM2.5
			PM10      int     `json:"pm10"`       // PM10
			NO2       int     `json:"no2"`        // 二氧化氮
			SO2       int     `json:"so2"`        // 二氧化硫
			O3        int     `json:"o3"`         // 臭氧
			CO        float64 `json:"co"`         // 一氧化碳
			Uptime    string  `json:"uptime"`     // 更新时间
		} `json:"now"`
		Indexes []struct {
			Name   string `json:"name"`   // 指数名称
			Brief  string `json:"brief"`  // 简要说明
			Detail string `json:"detail"` // 详细说明
		} `json:"indexes"`
		Alerts []struct {
			Type  string `json:"type"`  // 预警类型
			Level string `json:"level"` // 预警级别
			Title string `json:"title"` // 预警标题
			Desc  string `json:"desc"`  // 预警详情
		} `json:"alerts"`
		Forecasts []struct {
			Date      string `json:"date"`       // 日期
			Week      string `json:"week"`       // 星期
			High      int    `json:"high"`       // 最高温度
			Low       int    `json:"low"`        // 最低温度
			WCDay     string `json:"wc_day"`     // 白天风力
			WCNight   string `json:"wc_night"`   // 夜间风力
			WDDay     string `json:"wd_day"`     // 白天风向
			WDNight   string `json:"wd_night"`   // 夜间风向
			TextDay   string `json:"text_day"`   // 白天天气状况
			TextNight string `json:"text_night"` // 夜间天气状况
			AQI       int    `json:"aqi"`        // 空气质量指数
		} `json:"forecasts"`
		ForecastHours []struct {
			Text      string `json:"text"`       // 天气状况
			TempFC    int    `json:"temp_fc"`    // 温度
			WindClass string `json:"wind_class"` // 风力等级
			WindDir   string `json:"wind_dir"`   // 风向
			RH        int    `json:"rh"`         // 相对湿度
			Prec1h    int    `json:"prec_1h"`    // 1小时降水量
			Clouds    int    `json:"clouds"`     // 云量
			DataTime  string `json:"data_time"`  // 数据时间
		} `json:"forecast_hours"`
	} `json:"result"`
}

// BaiduWeatherTool 实现McpTool接口的百度天气工具
type BaiduWeatherTool struct {
}

// Handle 处理百度天气查询请求
func (t *BaiduWeatherTool) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 从请求参数中获取查询方式和对应的值
	queryType, ok := request.Params.Arguments["query_type"].(string)
	if !ok {
		return nil, errors.New("参数错误：query_type 必须是字符串类型")
	}

	var weatherResponse *BaiduWeatherResponse
	var err error

	switch queryType {
	case "district":
		district, ok := request.Params.Arguments["district"].(string)
		if !ok {
			return nil, errors.New("参数错误：district 必须是字符串类型")
		}
		weatherResponse, err = GetBaiduWeather(district)
	case "location":
		lon, ok := request.Params.Arguments["lon"].(string)
		if !ok {
			return nil, errors.New("参数错误：lon 必须是字符串类型")
		}
		lat, ok := request.Params.Arguments["lat"].(string)
		if !ok {
			return nil, errors.New("参数错误：lat 必须是字符串类型")
		}
		weatherResponse, err = QueryWeatherByLocation(lon, lat)
	case "keyword":
		keyword, ok := request.Params.Arguments["keyword"].(string)
		if !ok {
			return nil, errors.New("参数错误：keyword 必须是字符串类型")
		}
		weatherResponse, err = QueryWeatherByKeyword(keyword)
	default:
		return nil, errors.New("不支持的查询类型，请使用 'district'、'location' 或 'keyword'")
	}

	if err != nil {
		return nil, err
	}

	// 格式化天气信息
	now := weatherResponse.Result.Now
	location := weatherResponse.Result.Location

	fmt.Println(location)

	// 当前天气信息
	weatherText := fmt.Sprintf("【%s %s %s %s】\n当前天气：%s，温度：%d℃（体感 %d℃），湿度：%d%%，风向：%s %s\n\n",
		location.Country, location.Province, location.City, location.Name,
		now.Text, now.Temp, now.FeelsLike, now.RH, now.WindDir, now.WindClass)

	// 空气质量信息
	if now.AQI > 0 {
		weatherText += fmt.Sprintf("空气质量：AQI %d，PM2.5 %d，PM10 %d\n\n",
			now.AQI, now.PM25, now.PM10)
	}

	// 天气预警信息
	if len(weatherResponse.Result.Alerts) > 0 {
		weatherText += "预警信息：\n"
		for _, alert := range weatherResponse.Result.Alerts {
			weatherText += fmt.Sprintf("【%s %s】%s\n",
				alert.Type, alert.Level, alert.Title)
		}
		weatherText += "\n"
	}

	// 天气指数信息
	if len(weatherResponse.Result.Indexes) > 0 {
		weatherText += "生活指数：\n"
		for _, index := range weatherResponse.Result.Indexes {
			weatherText += fmt.Sprintf("【%s】%s，%s\n",
				index.Name, index.Brief, index.Detail)
		}
		weatherText += "\n"
	}

	// 未来天气预报
	if len(weatherResponse.Result.Forecasts) > 0 {
		weatherText += "未来天气预报：\n"
		for _, forecast := range weatherResponse.Result.Forecasts {
			weatherText += fmt.Sprintf("- %s（%s）：%s到%s，%d℃～%d℃，%s %s\n",
				forecast.Date, forecast.Week,
				forecast.TextDay, forecast.TextNight,
				forecast.Low, forecast.High,
				forecast.WDDay, forecast.WCDay)
		}
	}

	// 更新时间
	updateTime, err := formatUptime(now.Uptime)
	if err == nil {
		weatherText += fmt.Sprintf("\n更新时间：%s", updateTime)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: weatherText,
			},
		},
	}, nil
}

// formatUptime 格式化更新时间
func formatUptime(uptime string) (string, error) {
	if len(uptime) != 14 {
		return uptime, errors.New("时间格式不正确")
	}

	year := uptime[0:4]
	month := uptime[4:6]
	day := uptime[6:8]
	hour := uptime[8:10]
	minute := uptime[10:12]
	second := uptime[12:14]

	return fmt.Sprintf("%s-%s-%s %s:%s:%s", year, month, day, hour, minute, second), nil
}

// New 返回工具注册信息
func (t *BaiduWeatherTool) New() mcp.Tool {
	return mcp.NewTool("baiduWeather",
		mcp.WithDescription("获取百度天气信息"),
		mcp.WithString("query_type",
			mcp.Required(),
			mcp.Description("查询类型"),
			//mcp.Enum("district", "location", "keyword"),
			mcp.Enum("keyword"),
		),
		// mcp.WithString("district",
		// 	mcp.Description("区县名称，当query_type为district时必填"),
		// ),
		// mcp.WithString("lon",
		// 	mcp.Description("经度，当query_type为location时必填"),
		// ),
		// mcp.WithString("lat",
		// 	mcp.Description("纬度，当query_type为location时必填"),
		// ),
		mcp.WithString("keyword",
			mcp.Description("关键字，当query_type为keyword时必填"),
		),
	)
}

// SearchDistrictID 根据关键字搜索district_id
func SearchDistrictID(keyword string) (string, error) {
	var weatherInfo BaiduWeather

	var weatherAbroadInfo BaiduWeatherAbroad

	// 先在city字段中搜索关键字
	result := global.GVA_DB.Where("city LIKE ?", "%"+keyword+"%").First(&weatherInfo)
	if result.Error == nil {
		return weatherInfo.DistrictID, nil
	}

	// 如果city中没有找到，在district字段中搜索关键字
	result = global.GVA_DB.Where("district LIKE ?", "%"+keyword+"%").First(&weatherInfo)
	if result.Error == nil {
		return weatherInfo.DistrictID, nil
	}

	// 海外
	result = global.GVA_DB.Where("level2_chn LIKE ?", "%"+keyword+"%").First(&weatherAbroadInfo)
	if result.Error == nil {
		return weatherAbroadInfo.DistrictID, nil
	}

	// 海外
	result = global.GVA_DB.Where("level3_chn LIKE ?", "%"+keyword+"%").First(&weatherAbroadInfo)
	if result.Error == nil {
		return weatherAbroadInfo.DistrictID, nil
	}

	// 如果都没有找到，返回错误
	return "", fmt.Errorf("未找到与关键字 '%s' 相关的地区信息", keyword)
}

// QueryWeatherByKeyword 根据关键字查询天气
func QueryWeatherByKeyword(keyword string) (*BaiduWeatherResponse, error) {
	// 根据关键字搜索district_id和判断是否为海外地区
	districtID, isAbroad, err := SearchDistrictIDWithType(keyword)
	if err != nil {
		return nil, err
	}

	// 从环境变量获取百度API密钥
	baiduApiKey := os.Getenv("BAIDU_API_KEY")
	if baiduApiKey == "" {
		return nil, errors.New("未设置BAIDU_API_KEY环境变量")
	}

	// 构建API请求URL，根据是否是海外地区选择不同的API
	var url string
	if isAbroad {
		// 海外天气API
		url = fmt.Sprintf("https://api.map.baidu.com/weather_abroad/v1/?district_id=%s&data_type=all&ak=%s", districtID, baiduApiKey)
	} else {
		// 国内天气API
		url = fmt.Sprintf("https://api.map.baidu.com/weather/v1/?district_id=%s&data_type=all&ak=%s", districtID, baiduApiKey)
	}

	// 发送HTTP请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON响应
	var weatherResponse BaiduWeatherResponse
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		return nil, err
	}

	// 检查API响应状态
	if weatherResponse.Status != 0 {
		return nil, fmt.Errorf("百度天气API错误: %s", weatherResponse.Message)
	}

	return &weatherResponse, nil
}

// SearchDistrictIDWithType 根据关键字搜索district_id并判断是否为海外地区
func SearchDistrictIDWithType(keyword string) (string, bool, error) {
	var weatherInfo BaiduWeather
	var weatherAbroadInfo BaiduWeatherAbroad
	var weatherAbroadInfos []BaiduWeatherAbroad

	// 检查关键字是否包含空格，如果包含则进行分割处理
	parts := strings.Split(keyword, " ")
	if len(parts) == 2 {
		// 国内地区：第一部分作为province，第二部分作为district
		result := global.GVA_DB.Where("province LIKE ? AND district LIKE ?",
			"%"+parts[0]+"%", "%"+parts[1]+"%").First(&weatherInfo)
		if result.Error == nil {
			return weatherInfo.DistrictID, false, nil
		}

		// 海外地区：优化搜索逻辑
		// 1. 先查找匹配的国家
		result = global.GVA_DB.Where("country_chn = ?",
			parts[0]).Find(&weatherAbroadInfos)
		if result.Error == nil && len(weatherAbroadInfos) > 0 {
			// 2. 在匹配的国家中查找匹配的城市
			for _, info := range weatherAbroadInfos {
				if strings.Contains(info.Level2Chn, parts[1]) || strings.Contains(info.Level3Chn, parts[1]) {
					return info.DistrictID, true, nil
				}
			}
		}
	}

	// 常规搜索逻辑
	// 先在city字段中搜索关键字
	result := global.GVA_DB.Where("city LIKE ?", "%"+keyword+"%").First(&weatherInfo)
	if result.Error == nil {
		return weatherInfo.DistrictID, false, nil
	}

	// 如果city中没有找到，在district字段中搜索关键字
	result = global.GVA_DB.Where("district LIKE ?", "%"+keyword+"%").First(&weatherInfo)
	if result.Error == nil {
		return weatherInfo.DistrictID, false, nil
	}

	// 海外
	result = global.GVA_DB.Where("level2_chn LIKE ?", "%"+keyword+"%").First(&weatherAbroadInfo)
	if result.Error == nil {
		return weatherAbroadInfo.DistrictID, true, nil
	}

	// 海外
	result = global.GVA_DB.Where("level3_chn LIKE ?", "%"+keyword+"%").First(&weatherAbroadInfo)
	if result.Error == nil {
		return weatherAbroadInfo.DistrictID, true, nil
	}

	// 如果都没有找到，返回错误
	return "", false, fmt.Errorf("未找到与关键字 '%s' 相关的地区信息", keyword)
}

// GetBaiduWeather 获取百度天气信息
func GetBaiduWeather(district string) (*BaiduWeatherResponse, error) {
	var weatherInfo BaiduWeather
	result := global.GVA_DB.Where("district = ?", district).First(&weatherInfo)
	if result.Error != nil {
		return nil, errors.New("地区信息不存在")
	}

	// 从环境变量获取百度API密钥
	baiduApiKey := os.Getenv("BAIDU_API_KEY")
	if baiduApiKey == "" {
		return nil, errors.New("未设置BAIDU_API_KEY环境变量")
	}

	// 构建API请求URL
	url := fmt.Sprintf("https://api.map.baidu.com/weather/v1/?district_id=%s&data_type=all&ak=%s", weatherInfo.DistrictID, baiduApiKey)

	// 发送HTTP请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON响应
	var weatherResponse BaiduWeatherResponse
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		return nil, err
	}

	// 检查API响应状态
	if weatherResponse.Status != 0 {
		return nil, fmt.Errorf("百度天气API错误: %s", weatherResponse.Message)
	}

	return &weatherResponse, nil
}

// QueryWeatherByLocation 根据经纬度查询天气
func QueryWeatherByLocation(lon, lat string) (*BaiduWeatherResponse, error) {
	// 从环境变量获取百度API密钥
	baiduApiKey := os.Getenv("BAIDU_API_KEY")
	if baiduApiKey == "" {
		return nil, errors.New("未设置BAIDU_API_KEY环境变量")
	}

	// 构建API请求URL
	url := fmt.Sprintf("https://api.map.baidu.com/weather/v1/?location=%s,%s&data_type=all&ak=%s", lat, lon, baiduApiKey)

	// 发送HTTP请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON响应
	var weatherResponse BaiduWeatherResponse
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		return nil, err
	}

	// 检查API响应状态
	if weatherResponse.Status != 0 {
		return nil, fmt.Errorf("百度天气API错误: %s", weatherResponse.Message)
	}

	return &weatherResponse, nil
}
