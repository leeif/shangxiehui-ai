package jsonutils

type JsonStreamDecoder struct {
	text       []rune
	tokenStack []rune
	offset     int
}

func (dec *JsonStreamDecoder) Write(text string) []string {
	result := make([]string, 0)

	for _, r := range text {
		dec.text = append(dec.text, r)
		if dec.enqueueToken(r) {
			// if token has closed
			result = append(result, string(dec.text[dec.offset:]))
		}
	}

	return result
}

func (dec *JsonStreamDecoder) enqueueToken(r rune) bool {
	switch r {
	case rune('{'):
		if len(dec.tokenStack) == 0 {
			dec.offset = len(dec.text) - 1
		}
		dec.tokenStack = append(dec.tokenStack, r)
	case rune('}'):
		if len(dec.tokenStack) == 0 {
			return false
		}
		lastIndex := len(dec.tokenStack) - 1
		top := dec.tokenStack[lastIndex]
		if top == rune('{') {
			dec.tokenStack = dec.tokenStack[:lastIndex]
		} else {
			dec.tokenStack = append(dec.tokenStack, r)
		}

		if len(dec.tokenStack) == 0 {
			return true
		}
	}

	return false
}

func NewJsonStreamDecoder() *JsonStreamDecoder {
	return &JsonStreamDecoder{
		text:       make([]rune, 0),
		tokenStack: make([]rune, 0),
		offset:     0,
	}
}
