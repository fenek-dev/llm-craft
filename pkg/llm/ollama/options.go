package ollama

type Option func(*Ollama)

func WithModel(model string) Option {
	return func(o *Ollama) {
		if model != "" {
			o.model = model
		}
	}
}
