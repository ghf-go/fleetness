package core

import (
	"time"

	"github.com/ghf-go/fleetness/core/utils"
)

// 接口请求数据
type ApiParam struct {
	AppVer   string `json:"app_ver"`
	WgtVer   string `json:"wgt_ver"`
	Platform string `json:"platform"`
	OsVer    string `json:"os_ver"`
	OsLang   string `json:"os_lang"`
	ID       uint64 `json:"id"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}
type PageParam struct {
	ID        uint64   `json:"id"`
	Page      int      `json:"page"`
	PageSize  int      `json:"page_size"`
	RangeDate []string `json:"range_date"`
	SearchKey string   `json:"key"`
	TabName   string   `json:"tab"`
}

// 是否有日期范围
func (p *PageParam) HasDateRange() bool {
	return len(p.RangeDate) == 2
}

// 开始时间
func (p *PageParam) StartDate() time.Time {
	r, _ := time.Parse(utils.T_DATE, p.RangeDate[0])
	return r
}

// 结束时间
func (p *PageParam) EndDate() time.Time {
	r, _ := time.Parse(utils.T_DATE, p.RangeDate[1])
	return r
}
func (p *PageParam) getPage() int {
	if p.Page < 1 {
		p.Page = 1
	}
	return p.Page
}

func (p *PageParam) GetPageSize() int {
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	return p.PageSize
}
func (p *PageParam) GetOffset() int {
	return (p.getPage() - 1) * p.GetPageSize()
}

func (p *ApiParam) getPage() int {
	if p.Page < 1 {
		p.Page = 1
	}
	return p.Page
}

func (p *ApiParam) GetPageSize() int {
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	return p.PageSize
}
func (p *ApiParam) GetOffset() int {
	return (p.getPage() - 1) * p.GetPageSize()
}
