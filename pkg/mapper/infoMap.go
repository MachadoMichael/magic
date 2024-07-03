package mapper

type InfoMap struct {
	Resource string `json:"resource"`
	Path     string `json:"path"`
	CreateAt string `json:"createAt"`
	Builds   int    `json:"builds"`
	IsFile   bool   `json:"isFile"`
}
