# Define the path to your ABI JSON file and the output directory for the generated Go code
ABI_FILE := smart-contract/build/contracts/SecretStorage.json
ABI_EXTRACTED_FILE := smart-contract/build/contracts/SecretStorage_abi.json
GO_BINDINGS_OUTPUT_DIR := internal/ethereum

# Define the Go-Ethereum abigen tool path
ABIGEN := abigen

# Define the path to jq for extracting the ABI (make sure jq is installed)
JQ := jq

# Default target
all: generate-go-bindings

# Rule to extract ABI and generate Go code
generate-go-bindings: $(ABI_FILE)
	@mkdir -p $(GO_BINDINGS_OUTPUT_DIR)
	# Extract ABI array from JSON
	$(JQ) -r '.abi' $(ABI_FILE) > $(ABI_EXTRACTED_FILE)
	# Generate Go bindings from ABI file
	$(ABIGEN) --abi $(ABI_EXTRACTED_FILE) --pkg ethereum --out $(GO_BINDINGS_OUTPUT_DIR)/secret_storage.go
	# Clean up extracted ABI file
	rm $(ABI_EXTRACTED_FILE)

# Clean up generated files
clean:
	rm -rf $(GO_BINDINGS_OUTPUT_DIR)

# Phony targets
.PHONY: all generate-go-bindings clean
