package logic

var (
	dbConName    = "default"
	cacheConName = "default"
)

func SetDbConName(name string) {
	dbConName = name
}
func SetCacheConName(name string) {
	cacheConName = name
}
func GetDbConName() string {
	return dbConName
}
func GetCacheConName() string {
	return cacheConName
}
