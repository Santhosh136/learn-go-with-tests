package main

const (
	tamil = "Tamil"
	hindi = "Hindi"

	englishHelloPrefix = "Hello "
	tamilHelloPrefix   = "Vanakkam "
	hindiHelloPrefix   = "Namaste "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) string {
	prefix := englishHelloPrefix

	switch language {
	case tamil:
		prefix = tamilHelloPrefix
	case hindi:
		prefix = hindiHelloPrefix
	}

	return prefix
}
