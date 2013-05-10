package paging

type Paging interface {
	SetCurrentPage(uint)
	CurrentPage() uint
	LineSize() uint
	SetTotalPage()
	TotalPage() uint
	TotalCount() uint
}

type paging struct {
	currentPage uint // 当前页
	lineSize    uint // 每页几行
	totalPage   uint // 总页数
	totalCount  uint // 总行数
}

func New(lineSize, totalCount uint) Paging {
	return &paging{
		lineSize:   lineSize,
		totalCount: totalCount,
	}
}

func (p *paging) SetCurrentPage(currentPage uint) {
	p.currentPage = currentPage
}

func (p paging) CurrentPage() uint {
	return p.currentPage
}

func (p paging) LineSize() uint {
	return p.lineSize
}

func (p *paging) SetTotalPage() {
	totalPage := p.TotalCount() / p.LineSize()
	if p.TotalCount()%p.LineSize() == 0 {
		p.totalPage = totalPage
	} else {
		p.totalPage = totalPage + 1
	}
}

func (p paging) TotalPage() uint {
	return p.totalPage
}

func (p paging) TotalCount() uint {
	return p.totalCount
}
