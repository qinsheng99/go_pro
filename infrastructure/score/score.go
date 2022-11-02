package score

import (
	"bytes"
	"github.com/qinsheng99/go-py/api/score_api"
	"github.com/qinsheng99/go-py/domain/score"
	"os"
	"os/exec"
	"strconv"
)

type scoreImpl struct {
	score1, score2 string
}

func NewScore(s1, s2 string) score.Score {
	return &scoreImpl{
		score1: s1,
		score2: s2,
	}
}

func (s *scoreImpl) Evaluate(col score_api.Score) (data []byte, err error) {
	args := []string{s.score1, "--pred_path", col.PredPath, "--true_path", col.TruePath, "--cls", strconv.Itoa(col.Cls), "--pos", strconv.Itoa(col.Pos)}
	data, err = exec.Command("python3", args...).Output()

	if err != nil {
		return
	}
	data = bytes.ReplaceAll(bytes.TrimSpace(data), []byte(`'`), []byte(`"`))
	return
}

func (s *scoreImpl) Calculate(col score_api.Score) (data []byte, err error) {
	args := []string{s.score2, "--user_result", col.UserResult, "--unzip_path", os.Getenv("UPLOAD")}
	data, err = exec.Command("python3", args...).Output()

	if err != nil {
		return
	}
	data = bytes.ReplaceAll(bytes.TrimSpace(data), []byte(`'`), []byte(`"`))
	return
}
