package media

type Mood string

const (
	Chill    Mood = "chill"
	Dramatic Mood = "dramatic"
	Upbeat   Mood = "upbeat"
	Sad      Mood = "sad"
	MainMenu Mood = "main-menu"
)

var InGameMoods []Mood = []Mood{Chill, Dramatic, Upbeat, Sad}
