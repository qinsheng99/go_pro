package app

import (
	"bytes"
	"encoding/json"
	"github.com/qinsheng99/go-py/api/score_api"
	"github.com/qinsheng99/go-py/domain/score"
)

type scoreService struct {
	s score.Score
}

type ScoreService interface {
	Evaluate(score_api.Score, *score_api.ScoreRes) error
	Calculate(score_api.Score, *score_api.ScoreRes) error
}

func NewScoreService(s score.Score) ScoreService {
	return &scoreService{
		s: s,
	}
}

func (s *scoreService) Evaluate(col score_api.Score, res *score_api.ScoreRes) error {
	bys, err := s.s.Evaluate(col)
	if err != nil {
		return err
	}

	err = json.NewDecoder(bytes.NewBuffer(bys)).Decode(res)
	if err != nil {
		return err
	}
	return nil
}

func (s *scoreService) Calculate(col score_api.Score, res *score_api.ScoreRes) error {
	bys, err := s.s.Calculate(col)
	if err != nil {
		return err
	}

	err = json.NewDecoder(bytes.NewBuffer(bys)).Decode(res)
	if err != nil {
		return err
	}
	return nil
}
