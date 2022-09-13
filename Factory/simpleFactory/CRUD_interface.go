package simpleFactory

// DBOperate 共同类型,具有相同行为
// 这里需要共同类型行为,如果出现某个数据库支持但是某个数据库不支持就会很麻烦，所以这里应该是共用最多的，特殊的可以在对应的方法中实现。
type DBOperate interface {
	Create(dbName, table, key string)
	Read(dbName, table, key string)
	Update(dbName, table, key string)
	Delete(dbName, table, key string)
}
