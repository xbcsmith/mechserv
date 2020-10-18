TOOLS    = $(CURDIR)/../tools

all:
	$(TOOLS)/go-bindata -o templates.go --pkg templates *.html *.css && \
		sed -i 's/Html/HTML/g' templates.go && \
		sed -i 's/Css/CSS/g' templates.go && \
		gofmt -w -s templates.go
