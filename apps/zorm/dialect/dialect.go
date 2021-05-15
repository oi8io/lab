package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

//Dialect 适配各种数据之间的差异，像没个地方有自己的方言一样，表达吃饭用不懂的读音和文字
type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
