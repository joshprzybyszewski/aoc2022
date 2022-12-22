#!/bin/bash

benchmarks="$(go test \
    -benchmem \
    -run=^$ \
    -bench "^(BenchmarkAll)$" \
    github.com/joshprzybyszewski/aoc2022 \
    -short)"

if [[ "$benchmarks" != *"ok"* ]] || [[ "$benchmarks" != *"PASS"* ]]; then
    echo "benchmarks did not pass. results.md was not updated"
    exit 1
fi

rm results.md
cat scripts/default_results.md >> results.md

echo "" >> results.md
echo "## Benchmark Specs" >> results.md
echo "" >> results.md

echo "$benchmarks" \
| grep ":" \
| sed \
    -e 's/BenchmarkAll.*\//|/' \
| sed \
    '0~1 a\\' \
>> results.md

echo "" >> results.md
echo "## Puzzles" >> results.md
echo "" >> results.md
echo "|Puzzle|Duration|Bytes allocated to Heap|# of Heap allocations|" >> results.md
echo "|-|-:|-:|-:|" >> results.md

echo "$benchmarks" \
| grep BenchmarkAll \
| sed \
    -r \
    -e 's/([[:digit:]]+)([[:digit:]]{2})([[:digit:]]{7})[[:space:]]ns/\1.\2_s/' \
    -e 's/([[:digit:]]+)([[:digit:]]{2})([[:digit:]]{4})[[:space:]]ns/\1.\2_ms/' \
    -e 's/([[:digit:]]+)([[:digit:]]{2})([[:digit:]]{1})[[:space:]]ns/\1.\2_µs/' \
    -e 's/([[:digit:]])[[:space:]]ns/\1_ns/' \
| sed \
    -e 's/-[0-9]\+\s\+[0-9]\+//g' \
    -e 's/B\/op//' \
    -e 's/allocs\/op//' \
    -e 's/\/op//' \
    -e 's/\s\+/|/g' \
    -e 's/BenchmarkAll\//|/' \
    -e 's/[_\/]/ /g' \
>> results.md