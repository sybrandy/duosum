build: duosum.8
	8l -o duosum duosum.8

duosum.8:
	8g duosum.go

clean:
	rm *.8
	rm duosum
