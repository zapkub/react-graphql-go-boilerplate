
prepare:
	@echo Fetch go dependencies...
	dep ensure -v
	@echo Fetch node dependencies...
	cd client&&yarn
	@echo Done !

build-client:
	cd client&&yarn build

serve-client: build-client
	cd client&&yarn serve

serve-server:
	go run main.go