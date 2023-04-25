package database

var connection string

// will be executed when database package
// is called at the first time
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
