package notifier

const (
	PRIORITY_MAX  = 5
	PRIORITY_HIGH = 4
	PRIORITY_MED  = 3
	PRIORITY_LOW  = 2
	PRIORITY_ZERO = 1
)

type Payload struct {
	Title    string   `json:"Title"`
	Priority int      `json:"Priority"`
	Message  string   `json:"message"`
	Tags     []string `json:"Tags"`
}

func BuildPayload(name, when string) Payload {
	tag := []string{"rotating_light"}
	title := "PLANTÃO HYDRA"
	message := "Não há ninguém escalado para amanhã. Atualizar planilha!"

	if name != "" {
		message = name + ", " + when + " é o seu sobreaviso."
	}

	return Payload{
		Title:    title,
		Priority: PRIORITY_MAX,
		Message:  message,
		Tags:     tag,
	}
}
