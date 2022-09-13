package simpleFactory

/**
实现不同的数据库连接方式，比如MySQL/redis/mogonDB等等，也可以拆分出去。
*/

type DBConfig struct {
	DBName    string
	TableName string
	Key       string
	DBURL     string
	Port      string
}

func (db *DBConfig) DBCommon(DBURL, Port string) error {
	panic("implement me")
}

func (db *DBConfig) Create(dbName, table, key string) {
	panic("implement me")
}

func (db *DBConfig) Read(dbName, table, key string) {
	panic("implement me")
}

func (db *DBConfig) Update(dbName, table, key string) {
	panic("implement me")
}

func (db *DBConfig) Delete(dbName, table, key string) {
	panic("implement me")
}

func NewMySQLClient() string {
	return MySQL
}

func NewRedisClient() string {
	return Redis
}
