package common

type MemoryExport[TResult any] struct {
	Results []TResult
}

func NewMemoryExport[TResult any]() *MemoryExport[TResult] {
	return &MemoryExport[TResult]{
		Results: make([]TResult, 0),
	}
}

func (me *MemoryExport[TResult]) Export(exports chan interface{}) error {
	for result := range exports {
		me.Results = append(me.Results, result.(TResult))
	}

	return nil
}
