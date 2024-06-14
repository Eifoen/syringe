tools: (install "golang.org/x/tools/cmd/goimports") (install "honnef.co/go/tools/cmd/staticcheck")

# ---- HOUSEKEEPING ----
imports: (install "golang.org/x/tools/cmd/goimports")
	@echo "Running goimports..."
	@goimports -w "{{justfile_directory()}}"

tidy:
	@echo "Running go mod tidy..."
	@go mod tidy

# ---- FORMATTING ----
format: imports tidy
	@echo "Running gofmt..."
	@gofmt -s -w "{{justfile_directory()}}"

# ---- CHECKS ----
check: format staticcheck vet

staticcheck: (install "honnef.co/go/tools/cmd/staticcheck") 
	@echo "Running staticcheck..."
	@staticcheck "{{justfile_directory()}}"
	
deadcode: (install "golang.org/x/tools/cmd/deadcode")
	@echo "Running deadcode..."
	@deadcode "{{justfile_directory()}}"

vet: 
	@echo "Running go vet..."
	@go vet "{{justfile_directory()}}"
	
# ---- UTILS ----
install tool:
	@echo -n "Installing {{tool}}... "
	@go install {{tool}}
	@echo "done!"
