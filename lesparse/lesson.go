package lesparse

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"

	"github.com/yuin/goldmark"
)

type Lesson struct {
	markDownRaw     string
	Title           string     `json:"title"`
	Type            string     `json:"type"`
	Lang            string     `json:"lang"`
	DurationMinutes int        `json:"duration_minutes"`
	XP              int        `json:"xp"`
	Skill           int        `json:"skill"`
	AnswerIndex     int        `json:"answer"`
	Code            LessonCode `json:"code"`
	HTML            LessonHtml `json:"html"`
}

type LessonHtml struct {
	Instructions    string   `json:"instructions"`
	Hints           []string `json:"hints"`
	Body            string   `json:"markdown"`
	PossibleAnswers []string `json:"possible_answers"`
}

type LessonCode struct {
	PEC      string `json:"pec"`
	Sample   string `json:"sample"`
	Solution string `json:"solution"`
	Test     string `json:"test"`
}

var (
	lessonRegexp          = regexp.MustCompile("(?s)`{3}lesson(.*?)`{3}")
	instructionsRegexp    = regexp.MustCompile("(?s)`@instructions(.*?)\\.\\n{2}")
	hintsRegexp           = regexp.MustCompile("(?s)`@hint(.*?)\\n{2,}")
	pecRegexp             = regexp.MustCompile("(?s)`@pre_exercise_code`\n(`{3})(.*?)(`{3})")
	sampleCodeRegexp      = regexp.MustCompile("(?s)`@sample_code`\n(`{3})(.*?)(`{3})")
	solutionCodeRegexp    = regexp.MustCompile("(?s)`@solution`\n(`{3})(.*?)(`{3})")
	testCodeRegexp        = regexp.MustCompile("(?s)`@test`\n(`{3})(.*?)(`{3})")
	possibleAnswersRegexp = regexp.MustCompile("(?s)`@possible_answers(.*?)\\n{2,}")
	answerIndexRegexp     = regexp.MustCompile("(?s)`@answer(.*?)\\n{2,}")
)

func NewLesson(text string) *Lesson {
	return &Lesson{
		markDownRaw: text,
	}
}

func (l *Lesson) Parse() *Lesson {
	l.parseHeader()
	l.parseTitle()

	l.contentPEC()
	l.contentSampleCode()
	l.contentSolution()
	l.contentTest()

	l.renderPossibleAnswers()
	l.renderAnswer()
	l.renderInstructions()
	l.renderHints()
	l.renderMarkdown()

	return l
}

// contentMarkdown очищает текст от технических блоков, и возвращает
// чистый Markdown который должен прочитать пользователь.
func (l *Lesson) contentMarkdown() string {
	md := l.markDownRaw

	md = lessonRegexp.ReplaceAllString(md, "")
	md = instructionsRegexp.ReplaceAllString(md, "")
	md = sampleCodeRegexp.ReplaceAllString(md, "")
	md = solutionCodeRegexp.ReplaceAllString(md, "")
	md = testCodeRegexp.ReplaceAllString(md, "")
	md = pecRegexp.ReplaceAllString(md, "")
	md = hintsRegexp.ReplaceAllString(md, "")
	md = strings.ReplaceAll(md, "\n\n\n", "\n")

	return md
}

// renderMarkdown рендерит очищенный от технических блоков Markdown
// в HTML контент для пользователя.
func (l *Lesson) renderMarkdown() {
	l.HTML.Body = l.htmlRender(l.contentMarkdown())
}

// renderInstructions построчно рендерит инструкции в HTML.
func (l *Lesson) renderInstructions() {
	l.HTML.Instructions = l.htmlRender(l.contentInstructions())
}

// renderHints построчно рендерит подсказки в HTML.
func (l *Lesson) renderHints() {
	var hints []string
	for _, line := range l.contentHints() {
		hints = append(hints, l.htmlRender(line))
	}

	l.HTML.Hints = hints
}

// htmlRender занимается рендерингом markdown в html.
func (l *Lesson) htmlRender(text string) string {
	if len(text) == 0 {
		return ""
	}

	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(text), &buf); err != nil {
		return ""
	}

	return buf.String()
}

// contentTest находит блок текста @test и обновляет поле Test.
func (l *Lesson) contentTest() {
	var cont string
	if cont = testCodeRegexp.FindString(l.markDownRaw); len(cont) == 0 {
		return
	}

	l.Code.Test = l.trimCodeBlock(cont, "@test")
}

// contentSolution возвращает блок текста @solution.
func (l *Lesson) contentSolution() {
	var cont string
	if cont = solutionCodeRegexp.FindString(l.markDownRaw); len(cont) == 0 {
		return
	}

	l.Code.Solution = l.trimCodeBlock(cont, "@solution")
}

func (l *Lesson) renderAnswer() {
	l.AnswerIndex = l.contentAnswer()
}

func (l *Lesson) contentAnswer() int {
	var cont string
	if cont = answerIndexRegexp.FindString(l.markDownRaw); len(cont) == 0 {
		return 0
	}

	block := l.trimCodeBlock(cont, "@answer")
	block = strings.TrimSpace(block)

	i, err := strconv.Atoi(block)
	if err != nil {
		return 0
	}

	return i
}

func (l *Lesson) renderPossibleAnswers() {
	var answers []string

	for _, answer := range l.contentPossibleAnswers() {
		answers = append(answers, l.htmlRender(answer))
	}

	l.HTML.PossibleAnswers = answers
}

func (l *Lesson) contentPossibleAnswers() []string {
	var cont string
	if cont = possibleAnswersRegexp.FindString(l.markDownRaw); len(cont) == 0 {
		return []string{}
	}

	block := l.trimCodeBlock(cont, "@possible_answers")
	var possAnswers []string
	for _, line := range strings.Split(block, "\n") {
		possAnswers = append(possAnswers, line)
	}

	return possAnswers
}

func (l *Lesson) parseTitle() {
	lines := strings.Split(l.markDownRaw, "\n")
	if len(lines) == 0 {
		return
	}

	l.Title = strings.TrimLeft(lines[0], "# ")
}

// contentPEC возвращает блок текста @pre_exercise_code.
func (l *Lesson) contentPEC() {
	var cont string
	if cont = pecRegexp.FindString(l.markDownRaw); len(cont) == 0 {
		return
	}

	l.Code.PEC = l.trimCodeBlock(cont, "@pre_exercise_code")
}

// contentSampleCode возвращает код примераю
func (l *Lesson) contentSampleCode() {
	var cont string
	if cont = sampleCodeRegexp.FindString(l.markDownRaw); len(cont) == 0 {
		return
	}

	l.Code.Sample = l.trimCodeBlock(cont, "@sample_code")
}

// trimCodeBlock - тримит весь лишний текст из кодового блока с именем
// blockID.
// Например: trimCodeBlock(code, "@pre_exercise_code")
func (l *Lesson) trimCodeBlock(code, blockID string) string {
	res := strings.TrimPrefix(code, "`"+blockID+"`")
	res = strings.TrimLeft(res, " `\n")
	res = strings.TrimPrefix(res, "{"+l.Lang+"}")
	res = strings.TrimRight(res, " `\n")
	return strings.TrimSpace(res)
}

// contentHints возвращает блок текста @hint.
func (l *Lesson) contentHints() []string {
	var cont string
	if cont = hintsRegexp.FindString(l.markDownRaw); len(cont) == 0 {
		return []string{}
	}

	res := strings.TrimPrefix(cont, "`@hint`\n")
	res = strings.TrimSuffix(res, "`")
	res = strings.TrimSpace(res)

	var hints []string
	for _, line := range strings.Split(res, "\n") {
		hints = append(hints, line)
	}

	return hints
}

// contentInstructions возвращает блок текста @instructions.
func (l *Lesson) contentInstructions() string {
	var cont string
	if cont = instructionsRegexp.FindString(l.markDownRaw); len(cont) == 0 {
		return ""
	}

	res := strings.Trim(cont, "`@instructions")
	res = strings.TrimSpace(res)

	return res
}

// contentHeader возвращает блок текста lesson.
func (l *Lesson) contentHeader() string {
	var cont string
	if cont = lessonRegexp.FindString(l.markDownRaw); len(cont) == 0 {
		return ""
	}

	res := strings.TrimPrefix(cont, "```lesson")
	res = strings.TrimSuffix(res, "```")
	res = strings.TrimSpace(res)

	return res
}

// parseHeader заполняет структуру заголовка из блока текста lesson.
func (l *Lesson) parseHeader() {
	for _, line := range strings.Split(l.contentHeader(), "\n") {
		k, v, err := keyValue(line)
		if err != nil {
			continue
		}

		switch strings.ToLower(k) {
		case "duration_minutes":
			i, _ := strconv.Atoi(v)
			l.DurationMinutes = i
		case "type":
			l.Type = v
		case "lang":
			l.Lang = v
		case "xp":
			i, _ := strconv.Atoi(v)
			l.XP = i
		case "skills":
			i, _ := strconv.Atoi(v)
			l.Skill = i
		}
	}
}
