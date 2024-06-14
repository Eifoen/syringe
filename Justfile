tools: (install "golang.org/x/tools/cmd/goimports") (install "honnef.co/go/tools/cmd/staticcheck") (install "gotest.tools/gotestsum")

# ---- TESTS ----
test: (install "gotest.tools/gotestsum")
	@echo "Running tests"
	@gotestsum --format pkgname `go list -e ./... | grep -v github.com/eifoen/syringe/tools`


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
check: staticcheck vet

staticcheck: (install "honnef.co/go/tools/cmd/staticcheck") 
	@echo "Running staticcheck..."
	@staticcheck `go list -e ./... | grep -v github.com/eifoen/syringe/tools`
	
vet: 
	@echo "Running go vet..."
	@go vet "{{justfile_directory()}}"
	
# ---- UTILS ----
install tool:
	@echo -n "Installing {{tool}}... "
	@go install {{tool}}
	@echo "done!"
