package score

import (
	"bytes"
	"github.com/qinsheng99/go-domain-web/api/score_api"
	"github.com/qinsheng99/go-domain-web/domain/score"
	"os"
	"os/exec"
	"strconv"
)

type scoreImpl struct {
	evaluate, calculate string
}

func NewScore(s1, s2 string) score.Score {
	return &scoreImpl{
		evaluate:  s1,
		calculate: s2,
	}
}

func (s *scoreImpl) Evaluate(col score_api.Score) (data []byte, err error) {
	args := []string{s.evaluate, "--pred_path", col.PredPath, "--true_path", col.TruePath, "--cls", strconv.Itoa(col.Cls), "--pos", strconv.Itoa(col.Pos)}
	data, err = exec.Command("python3", args...).Output()

	if err != nil {
		return
	}
	data = bytes.ReplaceAll(bytes.TrimSpace(data), []byte(`'`), []byte(`"`))
	return
}

func (s *scoreImpl) Calculate(col score_api.Score) (data []byte, err error) {
	args := []string{s.calculate, "--user_result", col.UserResult, "--unzip_path", os.Getenv("UPLOAD")}
	data, err = exec.Command("python3", args...).Output()

	if err != nil {
		return
	}
	data = bytes.ReplaceAll(bytes.TrimSpace(data), []byte(`'`), []byte(`"`))
	return
}
