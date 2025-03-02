package valutesorting

import (
	"sort"

	vl "github.com/yanelox/task-3/internal/valute"
)

type less func(p1, p2 vl.Valute) bool

type valuteSorter struct {
	valutes []vl.Valute
	less    less
}

func (vs *valuteSorter) Len() int {
	return len(vs.valutes)
}

func (vs *valuteSorter) Swap(i, j int) {
	vs.valutes[i], vs.valutes[j] = vs.valutes[j], vs.valutes[i]
}

func (vs *valuteSorter) Less(i, j int) bool {
	return vs.less(vs.valutes[i], vs.valutes[j])
}

func ValuteSort(valutes []vl.Valute, less less) {
	var vs = &valuteSorter{valutes, less}
	sort.Sort(vs)
}
