package libs

import (
	"fmt"
	"math"
)

type Pages struct {
	Count   int
	Page    int
	PrePage int
	Url     string
}

var pageTpl string = `
<nav aria-label="Page navigation">
		<ul class="pagination pull-right">
			%s %s %s
		</ul>
	</nav>`

func (p *Pages) Get() string {
	pageCount := int(math.Ceil(float64(p.Count) / float64(p.PrePage)))
	var previousActive, previousPage, nextActive, nextPage string
	if p.Page-1 <= 0 {
		previousActive = `class="disabled"`
		previousPage = "javascript:void(0)"
	} else {
		previousPage = fmt.Sprintf("%s?p=%d", p.Url, p.Page-1)
	}
	previous := fmt.Sprintf(`<li %s>
		<a href="%s" aria-label="Previous">
			<span aria-hidden="true">&laquo;</span>
		</a>
	</li>`, previousActive, previousPage)
	if p.Page+1 > pageCount {
		nextActive = `class="disabled"`
		nextPage = "javascript:void(0)"
	} else {
		nextPage = fmt.Sprintf("%s?p=%d", p.Url, p.Page+1)
	}

	next := fmt.Sprintf(`<li %s>
		<a href="%s" aria-label="Next">
			<span aria-hidden="true">&raquo;</span>
		</a>
	</li>`, nextActive, nextPage)
	pages := ""
	var pDotLi, nDotLi bool
	for i := 1; i <= pageCount; i++ {
		active := ""
		if p.Page == i {
			active = `class="active"`
		}
		if i == 1 || i == pageCount || (p.Page-3 <= i && i <= p.Page+3) {
			pages += fmt.Sprintf(`<li %s ><a href="%s?p=%d">%d</a></li>`, active, p.Url, i, i)
		} else if !pDotLi && i < p.Page-3 {
			pDotLi = true
			pages += `<li class="disabled"><a href="javascript:void(0)">...</a></li>`
		} else if !nDotLi && i > p.Page+3 {
			nDotLi = true
			pages += `<li class="disabled"><a href="javascript:void(0)">...</a></li>`
		}

	}
	pageTpl := fmt.Sprintf(pageTpl, previous, pages, next)

	return pageTpl

}
