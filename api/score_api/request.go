package score_api

type Score struct {
	PredPath   string `json:"y_pred_path"`
	TruePath   string `json:"y_true_path"`
	UserResult string `json:"user_result"`
	Cls, Pos   int
}
