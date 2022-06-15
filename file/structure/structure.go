package structure

type Notestruct struct {
	Name        string       `json: "name"`
	FilePath    string       `json: "contents"`
	Owner       string       `json: "owner"`
	ID          string       `json: "id"`
	Permissions []Permission `json : "permissions`
}

type Permission struct {
	User     string `json: "user"`
	CanRead  string `json: "canRead"`
	CanWrite string `json: "canWrite"`
	CanAdmin string `json: "canAdmin"`
}
