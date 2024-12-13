package ip2region

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"gohub/pkg/logger"
	"strings"
	"sync"
)

// Searcher 全局唯一
var Searcher *xdb.Searcher

// once 确保全局的 Searcher 对象只实例一次
var once sync.Once

func Setup() {
	once.Do(func() {
		cBuff, err := xdb.LoadContentFromFile("./storage/ip2region/ip2region.xdb")
		logger.LogIf(err)
		Searcher, err = xdb.NewWithBuffer(cBuff)
		logger.LogIf(err)
	})
}

func Search(ip string) string {
	region, err := Searcher.SearchByStr(ip)
	logger.LogIf(err)
	return parseRegion(region)
}

// parseRegion 解析 region 字符串并返回格式化结果
func parseRegion(region string) string {
	// 如果 region 是空，返回 "Unknown"
	if region == "" {
		return "Unknown"
	}

	// 按 "|" 分割字符串
	parts := strings.Split(region, "|")
	if len(parts) < 5 {
		return "Unknown" // 确保数据完整性
	}

	country := parts[0]
	province := parts[2]
	city := parts[3]
	network := parts[4]

	// 根据 country 的值返回不同的结果
	if country == "中国" {
		return fmt.Sprintf("%s-%s:%s", province, city, network)
	}
	if country == "0" {
		return "Unknown"
	}
	return country
}
