package hello_world

// languages
const spanish = "Spanish"
const french = "French"
const japanese = "Japanese"
const chinese = "Chinese"

// prefixes
const englishHelloPrefix = "Hello, "
const spanishHelloprefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const japaneseHelloPrefix = "こんにちは、"
const chineseHelloPrefix = "你好, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPractice(language) + name
}

func greetingPractice(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloprefix
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix

}
