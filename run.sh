run () {
    input=$1
    expected=$2

    go run main.go "${input}" > main.s
    gcc -o main.o main.s && ./main.o

    output=$?

    if [ $output != $expected ]; then
        echo "${input}"
        echo "expected ${expected}, but got ${output}"
        exit 1
    else
        echo "${input} => ${output}"
    fi
}

run "0" 0
run "20" 20
run "255" 255

run "5 + 3" 8
run "5 - 3" 2
run "5 * 3" 15
run "15 / 3" 5
run "15 % 3" 0

run "5 + 3 + 10" 18
run "5 - 3 + 10" 12
run "2 * 3 + 4" 10
run "2 + 3 * 4" 14
run "2 * 3 + 4 * 5" 26
run "2 + 3 * 4 + 5" 19
run "6 / 3 * 10" 20
run "12 / 4 + 5" 8
run "6 + 30 / 10" 9

run "(2 + 3) * 4" 20
run "(2 + 3) * (4 + 5)" 45
run "26 / (10 + 3)" 2

echo "OK"