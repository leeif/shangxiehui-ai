package dto

type ChatCompletionTextStreamRequest struct {
	Messages []*ChatMessage `json:"messages"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Message string `json:"message"`
}

type ChatTextCompletionStream struct {
	Role  string `json:"role"`
	Chunk string `json:"chunk"`
}
