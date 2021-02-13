package lesparse

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Chapter(text string) *Chapter {
	blocks := p.blocks(text)

	// обнаружились ли блоки текста?
	if len(blocks) == 0 {
		return nil
	}

	// начинаем формировать тело сущности Chapter
	var chapter Chapter
	chapter.Header = p.header(blocks[0])

	// был ли заголовок? Если не был, то считаем что урок оформлен
	// без него. В противном случае вырезаем первый блок так как
	// считаем его заголовком.
	if chapter.Header != nil {
		blocks = blocks[1:]
	}

	chapter.LessonsCount = len(blocks)

	for _, block := range blocks {
		chapter.Lessons = append(chapter.Lessons, p.Lesson(block))
	}

	return &chapter
}

func (p *Parser) Lesson(text string) Lesson {
	return *NewLesson(text).Parse()
}
