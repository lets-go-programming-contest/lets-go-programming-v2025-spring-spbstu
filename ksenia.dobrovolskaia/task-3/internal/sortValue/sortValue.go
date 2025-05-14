package sortValue

import "github.com/kseniadobrovolskaia/task-3/internal/valute"

type ByValue []valute.Valute

func (x ByValue) Len() int           { return len(x) }
func (x ByValue) Less(i, j int) bool { return x[j].Value < x[i].Value }
func (x ByValue) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
