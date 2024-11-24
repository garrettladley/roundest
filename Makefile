NODE_BIN := ./node_modules/.bin

.PHONY:build
build: gen-css gen-templ
	@go build -tags dev -o bin/roundest cmd/server/main.go

.PHONY:build-prod
build-prod: gen-css gen-templ
	@go build -tags prod -o bin/roundest cmd/server/main.go

.PHONY:run
run: build
	@./bin/roundest

.PHONY: install
install: install-templ gen-templ
	@go get ./...
	@go mod tidy
	@go mod download
	@mkdir -p cmd/server/deps
	@wget -q -O cmd/server/deps/htmx-2.0.3.min.js.gz https://unpkg.com/htmx.org@2.0.3/dist/htmx.min.js.gz
	@gunzip -f cmd/server/deps/htmx-2.0.3.min.js.gz
	@npm install -D daisyui@latest
	@npm install -D tailwindcss


.PHONY: gen-css
gen-css:
	@$(NODE_BIN)/tailwindcss build -i internal/views/css/app.css -o cmd/server/public/styles.css --minify

.PHONY: watch-css
watch-css:
	@$(NODE_BIN)/tailwindcss -i internal/views/css/app.css -o cmd/server/public/styles.css --minify --watch

.PHONY: install-templ
install-templ:
	@go install github.com/a-h/templ/cmd/templ@latest

.PHONY: gen-templ
gen-templ:
	@templ generate

.PHONY: watch-templ
watch-templ:
	@templ generate --watch --proxy=http://127.0.0.1:8000

.PHONY: ci-scaffold
ci-scaffold:
	@mkdir -p cmd/server/deps
	@echo "hello world" > cmd/server/public/hello.txt
	@mkdir -p cmd/server/public
	@echo "hello world" > cmd/server/public/hello.txt
