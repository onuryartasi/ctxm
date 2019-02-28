package util

func GetCurrentContext() string {
	config := GetRawConfig()
	return config.CurrentContext
}
