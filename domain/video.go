package domain

type Classes struct {
	Modulos []struct {
		Name   string   `json:"name"`
		Folder string   `json:"folder"`
		Videos []string `json:"videos"`
	} `json:"modulos"`
}
