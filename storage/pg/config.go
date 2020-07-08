package pg

type Config struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
	Schema       string
	SSL          string
	MaxIdleConns int
	MaxOpenConns int
}
