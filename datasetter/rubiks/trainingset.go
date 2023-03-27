package rubiks

import (
	"bufio"
	"github.com/owulveryck/lstm/datasetter"
	"io"
)

type TrainingSet struct {
	rs        io.ReadSeeker
	buf       *bufio.Reader
	runeToIdx func(r rune) (int, error)
	batchSize int
	vocabSize int
	step      int
	pass      int
}

// NewTrainingSet from a ReadSeeker
func NewTrainingSet() *TrainingSet {
	return &TrainingSet{}
}

func (t *TrainingSet) GetTrainer() (datasetter.Trainer, error) {
	return nil, nil
}
