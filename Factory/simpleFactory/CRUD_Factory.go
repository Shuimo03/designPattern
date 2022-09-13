package simpleFactory

/**
根据传入参数，创建不同的数据库类型并返回实例。
*/

var (
	MySQL string
	Redis string
)

func GetDB(DbName string) (DBCategory func() string, err error) {
	switch DbName {
	case MySQL:
		return NewMySQLClient, nil
	case Redis:
		return NewRedisClient, nil
	}
	return nil, err
}
