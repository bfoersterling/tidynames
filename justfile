binary := "tidynames"

default:
	@just -l

build:
    go build

test:
	go test -v

install: default
	install -v -m 755 {{ binary }} /usr/local/bin/.

create_unicode_file:
	touch "test_files/c$(printf "\u00a9 File.txt")"

remove_unicode_file:
	rm -fv -- "test_files/c$(printf "\u00a9 File.txt")"
