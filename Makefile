start:
	PORT="3000" DATABASE_URL="marcelo:password@tcp(localhost:3306)/banking" nodemon --exec go run main.go --signal SIGTERM