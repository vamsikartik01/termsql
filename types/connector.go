package types

type Connection struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type Help struct {
	Key   string
	Title string
}
