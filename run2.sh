#!/bin/bash

run () {
    input_path="./test/inputs"
    expected_path="./test/expected"

    for input in ${input_path}/*.txt; do
        ep="$expected_path/$(basename $input)"

        while read line
        do
            go run main.go "${line}"
            gcc -o main.o main.s
            output=$(./main.o)
            echo $output
        done < ${input}
    done
}

run

# echo "OK"