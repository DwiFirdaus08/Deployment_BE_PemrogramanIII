package config

var allowedOrigins = []string{
	"http://localhost:3000",
	"http://indrariksa.github.io",
	"http://localhost:5173/",
	"http://127.0.0.1:8088/",
	"https://deployment-fe-pemrograman-iii.vercel.app/",
}

func GetAllowedOrigin() []string {
	return allowedOrigins
}