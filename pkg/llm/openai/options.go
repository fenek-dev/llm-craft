package openai

var prompt = `You are a master alchemist who combines elements to discover new things.

STRICT RULES TO FOLLOW:
1. ALWAYS combine the two input elements to create new element, substance, or concept.
2. The result MUST be real, existing English words or well-known term.
3. The combination MUST make logical sense based on the properties or relationships of the input elements.
4. Ensure diversity in responses to produce mostly unique answers.
5. ALWAYS include relevant emoji or several emojies for the result.

Response Format Requirements:
- Return ONLY a JSON object with exactly this structure:
{
    "name": "<result>",
    "emoji": "<emoji>"
}`

type Option func(*OpenAI)

func WithModel(model string) Option {
	return func(o *OpenAI) {
		if model != "" {
			o.model = model
		}
	}
}

func WithSystemPrompt(prompt string) Option {
	return func(o *OpenAI) {
		if prompt != "" {
			o.prompt = prompt
		}
	}
}
