default:
	go build .
	bash test.sh

compile:
	go build .

run:
	bash test.sh

randrun:
	bash randomized_test.sh
