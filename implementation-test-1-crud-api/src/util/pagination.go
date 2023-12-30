package util

const DateFormat = "02/01/2006"

type PaginParam struct {
	Limit   uint
	Page    uint
	OrderBy string
}

func (p *PaginParam) GetLimit() uint {
	if p.Limit == 0 || p.Limit > 100 {
		return 100
	}
	return p.Limit
}

func (p *PaginParam) GetOffset() uint {
	if p.Page < 1 {
		p.Page = 1
	}
	return p.Limit*p.Page - p.Limit
}

func (p *PaginParam) GetOrderBy(def string) string {
	if p.OrderBy == "" {
		return def
	}
	return p.OrderBy
}
