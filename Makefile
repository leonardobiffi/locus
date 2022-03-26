SOURCE=""
test:
	for i in {1..$(COUNT)}; do locus get -c 79831514 -s $(SOURCE); done
