package app

import (
	"bytes"
	"encoding/json"
	"github.com/qinsheng99/go-domain-web/api/score_api"
	"github.com/qinsheng99/go-domain-web/domain/score"
)

type scoreService struct {
	s score.Score
}

type ScoreServiceImpl interface {
	Evaluate(score_api.Score, *score_api.ScoreRes) error
	Calculate(score_api.Score, *score_api.ScoreRes) error
}

func NewScoreService(s score.Score) ScoreServiceImpl {
	return &scoreService{
		s: s,
	}
}

func (s *scoreService) Evaluate(col score_api.Score, res *score_api.ScoreRes) error {
	bys, err := s.s.Evaluate(col)
	if err != nil {
		return err
	}

	return json.NewDecoder(bytes.NewBuffer(bys)).Decode(res)
}

func (s *scoreService) Calculate(col score_api.Score, res *score_api.ScoreRes) error {
	bys, err := s.s.Calculate(col)
	if err != nil {
		return err
	}

	return json.NewDecoder(bytes.NewBuffer(bys)).Decode(res)
}
