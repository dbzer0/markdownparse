package lesparse

type Chapter struct {
	Header       *Header  `json:"header"`
	Lessons      []Lesson `json:"lessons"`
	LessonsCount int      `json:"lessons_count"`
}

type Header struct {
	Title         string   `json:"title"`
	TitleMeta     string   `json:"title_meta"`
	TitleImageURL string   `json:"title_image_url"`
	Description   string   `json:"description"`
	Attachments   []string `json:"attachments"`
	SlidesURL     string   `json:"slides_url"`
	FreePreview   bool     `json:"free_preview"`
}
