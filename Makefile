
build:
	docker image build . -t avantasia-txt

run: build
	docker run --env-file=.env avantasia-txt