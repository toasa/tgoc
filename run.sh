run () {
    input=$1
    expected=$2

    go run main.go $input > main.s
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

run "   30" 30
run " 255 " 255

echo "OK"