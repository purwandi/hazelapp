DATE=$(shell date +%I:%M%p)
CHECK=\033[32m✔\033[39m
HR=\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#\#

run:
	@echo "\n${HR}"
	@echo "Building knods platform..."
	@echo "${HR}\n"
	@cd gateway && go-bindata -ignore=\.go -pkg=schema -o=schema/bindata.go schema/...
	@echo "Parsing graphql schema...             ${CHECK} Done"
	@go run gateway/cmd/http/main.go
