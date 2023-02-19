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

type ChannelPipeExporter[TResult interface{}] struct {
	targetChan chan TResult
}

func NewChannelPipeExporter[TResult interface{}](target chan TResult) *ChannelPipeExporter[TResult] {
	return &ChannelPipeExporter[TResult]{targetChan: target}
}

func (cpe *ChannelPipeExporter[TResult]) Export(exports chan interface{}) error {
	go func() {
		for result := range exports {
			cpe.targetChan <- result.(TResult)
		}

		var emptyResult TResult
		cpe.targetChan <- emptyResult
	}()

	return nil
}
