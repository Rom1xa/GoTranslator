ANTLR_VERSION = 4.13.2
ANTLR_JAR     = tools/antlr-$(ANTLR_VERSION)-complete.jar
GRAMMAR       = grammar/GoSubset.g4
GEN_DIR       = internal/gen
BIN           = bin/goterp

.PHONY: tools generate build test docker-build docker-test clean

tools: $(ANTLR_JAR)

$(ANTLR_JAR):
	mkdir -p tools
	curl -L "https://www.antlr.org/download/antlr-$(ANTLR_VERSION)-complete.jar" -o $(ANTLR_JAR)

generate: tools
	rm -rf $(GEN_DIR)
	mkdir -p $(GEN_DIR)
	java -jar $(ANTLR_JAR) -Dlanguage=Go -package gen -visitor -o $(GEN_DIR) -lib grammar $(GRAMMAR)
	mv $(GEN_DIR)/grammar/*.go $(GEN_DIR)/ 2>/dev/null || true
	rm -rf $(GEN_DIR)/grammar

build:
	mkdir -p bin
	go build -o $(BIN) ./cmd/goterp

test:
	go test ./...

docker-build:
	docker build -t goterp .

docker-test: docker-build
	docker run --rm goterp make test

clean:
	rm -rf bin $(GEN_DIR)
