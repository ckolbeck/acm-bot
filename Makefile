GV=6

acmbot:
	$(GV)g main.go config.go xmlstructs.go
	$(GV)l -o acm-bot main.$(GV)

clean:
	-rm *.$(GV) *~ 2> /dev/null 