BINARY_NAME=myapp.exe
DB_URL=postgres://root:secret@localhost:5432/concurrency?sslmode=disable
REDIS=127.0.0.1:6379

## build: builds all binaries
build:
	@go build -o ${BINARY_NAME} ./cmd/web
	@echo back end built!

## run: builds and runs the application
run: build
	@echo Starting...
	set DB_URL=${DB_URL}
	set REDIS=${REDIS}
	set DB_URL
	set REDIS
	start /min cmd /c ${BINARY_NAME} &
	@echo back end started!

## clean: runs go clean and deletes binaries
clean:
	@echo Cleaning...
	@DEL ${BINARY_NAME}
	@go clean
	@echo Cleaned!

## start: an alias to run
start: run

## stop: stops the running application
stop:
	@echo "Stopping..."
	@taskkill /IM ${BINARY_NAME} /F
	@echo Stopped back end

## restart: stops and starts the application
restart: stop start

## test: runs all tests
test:
	@echo "Testing..."
	go test -v ./...