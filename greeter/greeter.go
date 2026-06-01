package greeter

// GreeterService is the INTERFACE — a contract that says:
// "any type that has a Greet(name) method qualifies as a GreeterService"
// This is Go's version of polymorphism. No 'implements' keyword needed.
type GreeterService interface {
	Greet(name string) string
	StyleName() string
}

// FormalGreeter is a STRUCT — Go's equivalent of a class
type FormalGreeter struct {
	Prefix string // field, like a class property
}

// Greet is a METHOD on FormalGreeter
// (g FormalGreeter) is the "receiver" — like 'this' in Java/JS
func (g FormalGreeter) Greet(name string) string {
	return g.Prefix + " " + name + ", welcome to our platform. We are delighted to have you."
}

func (g FormalGreeter) StyleName() string {
	return "formal"
}

// CasualGreeter — a second implementation of the same interface
type CasualGreeter struct{}

func (g CasualGreeter) Greet(name string) string {
	return "Hey " + name + "! Great to have you here 🎉"
}

func (g CasualGreeter) StyleName() string {
	return "casual"
}

// MultilingualGreeter — a third implementation using COMPOSITION
// It embeds CasualGreeter (reuse without inheritance)
type MultilingualGreeter struct {
	CasualGreeter        // embedded struct — inherits Greet and StyleName
	Language      string
}

// Override Greet — same idea as method overriding
func (g MultilingualGreeter) Greet(name string) string {
	greetings := map[string]string{
		"spanish":    "¡Hola " + name + "! Bienvenido 🌮",
		"french":     "Bonjour " + name + "! Bienvenue 🥐",
		"hindi":      "नमस्ते " + name + "! स्वागत है 🙏",
		"portuguese": "Olá " + name + "! Bem-vindo 🇧🇷",
	}
	if msg, ok := greetings[g.Language]; ok {
		return msg
	}
	return g.CasualGreeter.Greet(name) // fall back to embedded CasualGreeter
}

func (g MultilingualGreeter) StyleName() string {
	return "multilingual-" + g.Language
}

// NewGreeter is a FACTORY FUNCTION — Go's equivalent of a constructor
// It returns the interface type, so callers don't depend on a concrete struct
func NewGreeter(style string) GreeterService {
	switch style {
	case "formal":
		return FormalGreeter{Prefix: "Good day,"}
	case "spanish", "french", "hindi", "portuguese":
		return MultilingualGreeter{Language: style}
	default:
		return CasualGreeter{}
	}
}
