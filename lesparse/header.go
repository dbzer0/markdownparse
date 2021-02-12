package lesparse

import "strings"

// headerMap вычленяет заголовок, если он есть. Возвращает map с ключ-значением.
func (p *Parser) headerMap(text string) map[string]string {
	const titleTagName = "title"

	// проверяем наличие заголовка по полю title
	// если его нет, то возвращаем пустую мапу
	if !strings.Contains(text, titleTagName) {
		return nil
	}

	kv := map[string]string{}
	for _, line := range strings.Split(text, "\n") {
		// строка не содержит key: value
		if !strings.Contains(line, ":") {
			continue
		}

		k, v, err := keyValue(line)
		if err != nil {
			continue
		}

		kv[k] = v
	}

	return kv
}

// header формирует из текста структуру заголовка урока.
func (p *Parser) header(text string) *Header {
	var header Header

	headerMap := p.headerMap(text)
	if headerMap == nil {
		return nil
	}

	if v, ok := headerMap["title"]; ok {
		header.Title = v
	}
	if v, ok := headerMap["title_meta"]; ok {
		header.TitleMeta = v
	}
	if v, ok := headerMap["description"]; ok {
		header.Description = v
	}
	if v, ok := headerMap["attachments"]; ok {
		// TODO: оттримить пробелы между запятыми
		header.Attachments = strings.Split(v, ",")
	}
	if v, ok := headerMap["slides_url"]; ok {
		header.SlidesURL = v
	}
	if v, ok := headerMap["free_preview"]; ok {
		if v == "true" {
			header.FreePreview = true
		}
	}

	return &header
}
