package mapper

type InfoMap struct {
	Path     string `json:"path"`
	CreateAt string `json:"createAt"`
	Builds   int    `json:"builds"`
}
