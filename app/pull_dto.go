package app

import "github.com/qinsheng99/go-domain-web/domain"

type pullRequestDTO struct {
	List  []domain.PullInfo `json:"list"`
	Total int64             `json:"total"`
}

func toPullRequestDTO(v []domain.PullInfo, total int64) pullRequestDTO {
	return pullRequestDTO{
		List:  v,
		Total: total,
	}
}
