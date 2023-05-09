package controller

type UvpData struct {
	Id         string           `json:"id"         binding:"required"`
	Desc       string           `json:"desc"       binding:"required"`
	Source     string           `json:"source"     binding:"required"`
	Pushed     string           `json:"pushed"     binding:"required"`
	PushType   string           `json:"push_type"  binding:"required"`
	Affected   []string         `json:"affected"   binding:"required"`
	Published  string           `json:"published"  binding:"required"`
	Severity   []severity       `json:"severity"`
	References []referencesData `json:"references"`
	Patch      []patch          `json:"patch"`
}

type severity struct {
	Type   string `json:"type"`
	Score  string `json:"score"`
	Vector string `json:"vector"`
}

type referencesData struct {
	Url  string `json:"url"`
	Type string `json:"type"`
}

type patch struct {
	Package    string `json:"package"`
	FixVersion string `json:"fix_version"`
	FixPatch   string `json:"fix_patch"`
	BreakPatch string `json:"break_patch"`
	Source     string `json:"source"`
	Branch     string `json:"branch"`
}
