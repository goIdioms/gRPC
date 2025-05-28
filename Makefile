.PHONY: proto

proto:
	@echo "Generating proto files..."
	@chmod +x scripts/generate-proto.sh
	@./scripts/generate-proto.sh