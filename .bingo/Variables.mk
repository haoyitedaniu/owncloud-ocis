# Auto generated binary variables helper managed by https://github.com/bwplotka/bingo v0.3.1. DO NOT EDIT.
# All tools are designed to be build inside $GOBIN.
BINGO_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
GOPATH ?= $(shell go env GOPATH)
GOBIN  ?= $(firstword $(subst :, ,${GOPATH}))/bin
GO     ?= $(shell which go)

# Bellow generated variables ensure that every time a tool under each variable is invoked, the correct version
# will be used; reinstalling only if needed.
# For example for bingo variable:
#
# In your main Makefile (for non array binaries):
#
#include .bingo/Variables.mk # Assuming -dir was set to .bingo .
#
#command: $(BINGO)
#	@echo "Running bingo"
#	@$(BINGO) <flags/args..>
#
BINGO := $(GOBIN)/bingo-v0.3.1
$(BINGO): $(BINGO_DIR)/bingo.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/bingo-v0.3.1"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=bingo.mod -o=$(GOBIN)/bingo-v0.3.1 "github.com/bwplotka/bingo"

BUF := $(GOBIN)/buf-v0.37.1
$(BUF): $(BINGO_DIR)/buf.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/buf-v0.37.1"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=buf.mod -o=$(GOBIN)/buf-v0.37.1 "github.com/bufbuild/buf/cmd/buf"

CALENS := $(GOBIN)/calens-v0.2.0
$(CALENS): $(BINGO_DIR)/calens.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/calens-v0.2.0"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=calens.mod -o=$(GOBIN)/calens-v0.2.0 "github.com/restic/calens"

FILEB0X := $(GOBIN)/fileb0x-v1.1.4
$(FILEB0X): $(BINGO_DIR)/fileb0x.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/fileb0x-v1.1.4"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=fileb0x.mod -o=$(GOBIN)/fileb0x-v1.1.4 "github.com/UnnoTed/fileb0x"

FLAEX := $(GOBIN)/flaex-v0.2.0
$(FLAEX): $(BINGO_DIR)/flaex.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/flaex-v0.2.0"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=flaex.mod -o=$(GOBIN)/flaex-v0.2.0 "github.com/owncloud/flaex"

GOLANGCI_LINT := $(GOBIN)/golangci-lint-v1.35.2
$(GOLANGCI_LINT): $(BINGO_DIR)/golangci-lint.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/golangci-lint-v1.35.2"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=golangci-lint.mod -o=$(GOBIN)/golangci-lint-v1.35.2 "github.com/golangci/golangci-lint/cmd/golangci-lint"

GOLINT := $(GOBIN)/golint-v0.0.0-20201208152925-83fdc39ff7b5
$(GOLINT): $(BINGO_DIR)/golint.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/golint-v0.0.0-20201208152925-83fdc39ff7b5"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=golint.mod -o=$(GOBIN)/golint-v0.0.0-20201208152925-83fdc39ff7b5 "golang.org/x/lint/golint"

GOVERAGE := $(GOBIN)/goverage-v0.0.0-20180129164344-eec3514a20b5
$(GOVERAGE): $(BINGO_DIR)/goverage.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/goverage-v0.0.0-20180129164344-eec3514a20b5"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=goverage.mod -o=$(GOBIN)/goverage-v0.0.0-20180129164344-eec3514a20b5 "github.com/haya14busa/goverage"

GOX := $(GOBIN)/gox-v1.0.1
$(GOX): $(BINGO_DIR)/gox.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/gox-v1.0.1"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=gox.mod -o=$(GOBIN)/gox-v1.0.1 "github.com/mitchellh/gox"

HUGO := $(GOBIN)/hugo-v0.80.0
$(HUGO): $(BINGO_DIR)/hugo.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/hugo-v0.80.0"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=hugo.mod -o=$(GOBIN)/hugo-v0.80.0 "github.com/gohugoio/hugo"

PROTOC_GEN_DOC := $(GOBIN)/protoc-gen-doc-v1.4.1
$(PROTOC_GEN_DOC): $(BINGO_DIR)/protoc-gen-doc.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/protoc-gen-doc-v1.4.1"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=protoc-gen-doc.mod -o=$(GOBIN)/protoc-gen-doc-v1.4.1 "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"

PROTOC_GEN_GO := $(GOBIN)/protoc-gen-go-v1.25.0
$(PROTOC_GEN_GO): $(BINGO_DIR)/protoc-gen-go.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/protoc-gen-go-v1.25.0"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=protoc-gen-go.mod -o=$(GOBIN)/protoc-gen-go-v1.25.0 "google.golang.org/protobuf/cmd/protoc-gen-go"

PROTOC_GEN_MICRO := $(GOBIN)/protoc-gen-micro-v3.0.0-20210217182006-0f0ace1a44a9
$(PROTOC_GEN_MICRO): $(BINGO_DIR)/protoc-gen-micro.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/protoc-gen-micro-v3.0.0-20210217182006-0f0ace1a44a9"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=protoc-gen-micro.mod -o=$(GOBIN)/protoc-gen-micro-v3.0.0-20210217182006-0f0ace1a44a9 "github.com/asim/go-micro/cmd/protoc-gen-micro/v3"

PROTOC_GEN_MICROWEB := $(GOBIN)/protoc-gen-microweb-v0.0.0-20210224131655-d9b1137a84d4
$(PROTOC_GEN_MICROWEB): $(BINGO_DIR)/protoc-gen-microweb.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/protoc-gen-microweb-v0.0.0-20210224131655-d9b1137a84d4"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=protoc-gen-microweb.mod -o=$(GOBIN)/protoc-gen-microweb-v0.0.0-20210224131655-d9b1137a84d4 "github.com/owncloud/protoc-gen-microweb"

PROTOC_GEN_OPENAPIV2 := $(GOBIN)/protoc-gen-openapiv2-v2.2.0
$(PROTOC_GEN_OPENAPIV2): $(BINGO_DIR)/protoc-gen-openapiv2.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/protoc-gen-openapiv2-v2.2.0"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=protoc-gen-openapiv2.mod -o=$(GOBIN)/protoc-gen-openapiv2-v2.2.0 "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"

REFLEX := $(GOBIN)/reflex-v0.3.0
$(REFLEX): $(BINGO_DIR)/reflex.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/reflex-v0.3.0"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=reflex.mod -o=$(GOBIN)/reflex-v0.3.0 "github.com/cespare/reflex"

STATICCHECK := $(GOBIN)/staticcheck-v0.1.2
$(STATICCHECK): $(BINGO_DIR)/staticcheck.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/staticcheck-v0.1.2"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=staticcheck.mod -o=$(GOBIN)/staticcheck-v0.1.2 "honnef.co/go/tools/cmd/staticcheck"

