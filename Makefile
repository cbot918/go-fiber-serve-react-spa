PROJNAME=ui1

ui:
	yarn create vite $(PROJNAME) --template react && cd $(PROJNAME) && yarn

rest:
	go run .

.PHONY: ui