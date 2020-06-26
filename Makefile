swagger=docker run --rm -it -e GOPATH=$(HOME)/go:/go -v $(HOME):$(HOME) -w $$(pwd) quay.io/goswagger/swagger:v0.24.0
generate:
	$(swagger) generate client --target=./netbox --spec=./swagger.json --copyright-file=./copyright_header.txt

clean:
	rm -rf netbox/client netbox/models

integration:
	go test ./... -tags=integration
