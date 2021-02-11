package lesparse

import "strings"

// keyValue из строки вычленяет ключ и значение.
func keyValue(line string) (string, string) {
	kv := strings.SplitN(line, ":", 2)
	return strings.Trim(kv[0], " "), strings.Trim(kv[1], " ")
}

func (p *Parser) blocks(text string) []string {
	var blocks []string
	var pos int

	const pattern = "---\n"

	cur := text
	for {
		pos = strings.Index(cur, pattern)

		if pos > 0 {
			blocks = append(blocks, strings.Trim(cur[:pos+len(pattern)], " \n"))
		}
		if pos < 0 {
			blocks = append(blocks, strings.Trim(cur[pos+1:], " \n"))
			break
		}
		cur = cur[pos+len(pattern):]
	}

	return blocks
}
