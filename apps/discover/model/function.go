package model

func getKey(appid, env string) string {
	return appid + "|" + env
}
