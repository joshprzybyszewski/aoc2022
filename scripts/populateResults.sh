#!/bin/bash

rm results.md
cat scripts/default_results.md >> results.md

echo "|Puzzle|ns/op|B/op|allocs/op|" >> results.md
echo "|-|-:|-:|-:|" >> results.md

go test \
    -benchmem \
    -run=^$ \
    -bench "^(BenchmarkAll)$" \
    github.com/joshprzybyszewski/aoc2022 \
    -short \
| grep BenchmarkAll \
| sed \
    -e 's/-[0-9]\+\s\+[0-9]\+//g' \
    -e 's/ns\/op//' \
    -e 's/B\/op//' \
    -e 's/allocs\/op//' \
    -e 's/\s\+/|/g' \
    -e 's/BenchmarkAll\//|/' \
    -e 's/[_\/]/ /g' \
>> results.md