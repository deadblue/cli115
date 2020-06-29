GO_EXE=go
GOBUILD_EXE=$(GO_EXE) build

LDFLAGS="-s -w"

OUTPUT_DIR=output
CMDS = $(shell find ./cmd -mindepth 1 -maxdepth 1 -type d -exec basename {} \;)

all: build

build:
	@mkdir -p $(OUTPUT_DIR)
	@for cmd in $(CMDS) ; do \
  		echo "Building $$cmd ..."; \
  		$(GOBUILD_EXE) -o $(OUTPUT_DIR)/$$cmd -ldflags=$(LDFLAGS) ./cmd/$$cmd/*.go; \
	done

clean:
	rm -rf $(OUTPUT_DIR)
