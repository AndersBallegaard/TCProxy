package main

func main() {
	instance := loadConfig("settings.yaml")
	instance.server()
}
