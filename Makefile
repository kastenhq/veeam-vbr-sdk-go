golang_spec = spec/openapi_spec.yaml
vbr_spec = spec/swagger.json

default: generate 
cleanup:
	@echo "Cleaning up..."
	rm -f ./vbrclient/types.go
	rm -f ./vbrclient/client.go

generate: cleanup
	@echo "Generating types..."
	oapi-codegen -generate types -o ./vbrclient/types.go -package vbrclient ${golang_spec}
	@echo "Generating client..."
	oapi-codegen -generate client -o ./vbrclient/client.go -package vbrclient ${golang_spec}
	cd vbrclient && go mod tidy

convert:
	@echo "Converting spec..."
	cd tools/oapifixer && go build
	./tools/oapifixer/oapifixer -input ${vbr_spec} -output ${golang_spec}