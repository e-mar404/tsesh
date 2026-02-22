package bookmark

type Entries struct {
	Data []Entry `json:"entries"`
}

type Entry struct {
	Url       string   `json:"url"`
	Directory string   `json:"directory"`
	Tags      []string `json:"tags"`
}
