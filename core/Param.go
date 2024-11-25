package core

type PageParam struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (p *PageParam) GetPage() int {
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
	return (p.GetPage() - 1) * p.GetPageSize()
}
