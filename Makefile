BRANCH	     = main
PWD          = `pwd`
PACKAGE      = `basename $(PWD)`
REV          = `git rev-list --tags --max-count=1`
VERSION      = `git describe --tags $(REV)`
DATE	       = `date "+%Y.%m%d%"`
NEXT         = `autotag -b $(BRANCH) -n`
RELEASE_DIR  = ../dist
RELEASE_FILE = $(PACKAGE)-$(VERSION)

.SILENT: all
.SILENT: docs
.SILENT: release-next
.SILENT: _patch
.SILENT: _packages

# Default target.
all:
	echo "Hello $(LOGNAME), nothing to do by default"

docs:
	mkdir -p docs
	godoc-static -site-name="$(RELEASE_FILE)" -destination=./docs . 

release:
	autotag -b $(BRANCH)
	goreleaser release

release-next:
	echo "Next release: v$(NEXT)"

release-build:
	goreleaser release --snapshot --rm-dist

.PHONY:release-all release docs

_patch:
	# fetch all tags and history:
	git fetch --tags --prune

	if [ `git rev-parse --abbrev-ref HEAD` != "$(BRANCH)" ]; then
		# ensure a local branch exists at 'refs/heads/master'
		git branch --track master origin/$(BRANCH)
	fi

_packages:
	go mod download
