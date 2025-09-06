package domain

type Lesson struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	VideoURL string `json:"videoUrl"`
}

type Module struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Lessons []Lesson `json:"lessons"`
}

type Course struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ImageURL    string   `json:"imageUrl"`
	Modules     []Module `json:"modulos"`
}

type Classes struct {
	Modulos []struct {
		Name   string   `json:"name"`
		Folder string   `json:"folder"`
		Videos []string `json:"videos"`
	} `json:"modulos"`
}