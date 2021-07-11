module gowebapp1

// +heroku goVersion go1.16
go 1.16

require (
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/sessions v1.2.1
	github.com/jackc/pgproto3/v2 v2.1.0 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.10.2 // indirect
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e // indirect
	golang.org/x/text v0.3.6 // indirect
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.11
)

replace routes => ./routes
