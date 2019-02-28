package util

func GetCurrentContext() (string, string) {
	config := GetRawConfig()
	ns := config.Contexts[config.CurrentContext].Namespace
	if !(len(ns) > 0) {
		ns = "default"
	}
	return config.CurrentContext, ns
}
