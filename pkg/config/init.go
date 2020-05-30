package config

func init() {
	Server = (&server{}).Load("conf/app.ini").Init()
}
