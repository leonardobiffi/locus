SOURCE=""
test:
	for i in {1..$(COUNT)}; do go run main.go get -c 79831514 -s $(SOURCE); echo "Test: $$i"; done
