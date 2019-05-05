package core

type BatchSelector struct {
	MinSize     int
	MaxSize     int
	ActiveBatch []string
	Batches     [][]string
}

func NewBatchSelector(minSize int, maxSize int) *BatchSelector {
	return &BatchSelector{
		MinSize:     minSize,
		MaxSize:     maxSize,
		ActiveBatch: make([]string, minSize),
		Batches:     make([][]string, 0),
	}
}


func (bs *BatchSelector) AddPlayer(playerID string) {
	if len(bs.ActiveBatch) >= bs.MaxSize {
		bs.Batches = append(bs.Batches, bs.ActiveBatch)
		bs.ActiveBatch = make([]string, bs.MinSize)
	}

	bs.ActiveBatch = append(bs.ActiveBatch, playerID)
}

func (bs *BatchSelector) GetSelectedBatches() [][]string {
	if len(bs.ActiveBatch) >= bs.MinSize {
		return append(bs.Batches, bs.ActiveBatch)
	} else {
		return bs.Batches
	}
}