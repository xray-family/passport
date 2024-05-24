test:
	go test -count=1 ./...

cover:
	go test -coverprofile=./bin/cover.out --cover ./...

update-dict:
	goi18n merge -sourceLanguage en-US active.*.toml
