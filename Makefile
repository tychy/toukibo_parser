BUCKET_NAME=toukibo-parser-samples
URL=https://pub-a26a7972d1ea437b983bf6696a7d847e.r2.dev
DATA_DIR=testdata
NUM_SAMPLE=778

build:
	mkdir -p bin
	go build -o bin/toukibo-parser main.go

run: build
	./bin/toukibo-parser -path=$(TARGET).pdf

run/sample: build
	./bin/toukibo-parser -path="$(DATA_DIR)/pdf/$(TARGET).pdf"
	
run/all: build
	./run-all-sample.sh
# ちょっとおかしいもの
# 133 住所が途中で切れている
# 770 住所のパースがおかしい

test: build
	go test -coverprofile=coverage.out -shuffle=on ./...

zip/sample:
	zip -r testdata.zip testdata

put/sample: zip/sample
	wrangler r2 object delete $(BUCKET_NAME)/testdata.zip
	wrangler r2 object put $(BUCKET_NAME)/testdata.zip --file testdata.zip
	
get/sample: clean/data
	mkdir -p $(DATA_DIR)
	curl -o testdata.zip $(URL)/testdata.zip
	unzip testdata.zip

clean: clean/bin clean/data

clean/bin:
	rm -rf bin

clean/data:
	rm -rf $(DATA_DIR)
	rm -rf testdata.zip
