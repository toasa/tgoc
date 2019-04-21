run () {
    input=$1
    expected=$2

    go run main.go "${input}" > main.s
    gcc -o main.o main.s && ./main.o

    output=$?

    if [ $output != $expected ]; then
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
run "5 - 3 + 10" 12

echo "OK"