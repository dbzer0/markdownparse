package lesparse

type Chapter struct {
	Header     *Header  `json:"header"`
	Lessons    []Lesson `json:"lessons"`
	PagesCount int      `json:"pages_count"`
}

type Header struct {
	Title       string   `json:"title"`
	TitleMeta   string   `json:"title_meta"`
	Description string   `json:"description"`
	Attachments []string `json:"attachments"`
	SlidesURL   string   `json:"slides_url"`
	FreePreview bool     `json:"free_preview"`
}
