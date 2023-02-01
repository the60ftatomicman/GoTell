burn:
	pkill -f *gotell | echo done
	ps

#stty -icanon && nc localhost 9002
run:
	- $(MAKE) burn
	go run .