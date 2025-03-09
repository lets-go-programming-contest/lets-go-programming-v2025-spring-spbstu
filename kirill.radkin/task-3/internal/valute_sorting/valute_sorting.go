package valutesorting

import (
	"sort"

	"github.com/yanelox/task-3/internal/mytypes"
)

type less func(p1, p2 mytypes.Valute) bool

type valuteSorter struct {
	valutes []mytypes.Valute
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

func ValuteSort(valutes []mytypes.Valute, less less) {
	var vs = &valuteSorter{valutes, less}
	sort.Sort(vs)
}
