
build:
	docker image build . -t avantasia-txt

run: build
	docker run --env-file=.env -p 8080:8080 avantasia-txt 