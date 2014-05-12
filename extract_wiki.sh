#!/bin/sh
go run extract_wiki.go | ruby to_space.rb | wordcut --dict dict.txt > wiki.txt
