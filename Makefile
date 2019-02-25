dev-install:
	@command -v terraform >/dev/null 2>&1 && { printf '%s\n' 'Terraform installed: OK'; exit 0; } || { printf '%s\n' 'Terraform not installed, installing with brew...'; brew install terraform; exit 0; }
	@command -v dep >/dev/null 2>&1 && { printf '%s\n' 'GoDep installed: OK'; exit 0; } || { printf '%s\n' 'GoDep not installed, installing with brew...'; brew install dep; exit 0; }
	@printf "Running dep ensure...\n" && dep ensure -v
tests:
	@go test -v | grep -v `date '+%Y'`