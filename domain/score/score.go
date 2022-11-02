package score

import (
	"github.com/qinsheng99/go-py/api/score_api"
)

type Score interface {
	Evaluate(score_api.Score) ([]byte, error)
	Calculate(score_api.Score) ([]byte, error)
}
