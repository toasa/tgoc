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

run "0;" 0
run "20" 20
run "255" 255

run "5 + 3;" 8
run "5 - 3" 2
run "5 * 3;" 15
run "15 / 3" 5
run "15 % 3" 0

run "5 + 3 + 10;" 18
run "5 - 3 + 10" 12
run "2 * 3 + 4" 10
run "2 + 3 * 4" 14
run "2 * 3 + 4 * 5;" 26
run "2 + 3 * 4 + 5" 19
run "6 / 3 * 10" 20
run "12 / 4 + 5" 8
run "6 + 30 / 10" 9

run "(2 + 3) * 4" 20
run "(2 + 3) * (4 + 5)" 45
run "26 / (10 + 3)" 2

run "a := 20 * 2; a" 40
run "abc := 30 + 4 * 2; xyz := abc * 2; xyz;" 76
run "a := 1; b := 1; c := a + b; d := b + c; e := c + d; e;" 5
run "a := 1; b := 1; c := a + b; d := b + c; return d; e := c + d; e;" 3

run "a := 2 * 3; return a; b := 40;" 6

run "+10" 10
run "-2 + 6;" 4
run "-(10 - 16)" 6
run "-3*+5*-2" 30
run "a := -20; return -a;" 20

run "1 << 7" 128
run "1024 >> 3" 128

run "20 == 20" 1
run "20 != 20" 0
run "1 + 2 + 3 == 1 * 2 * 3" 1
run "0 == 1" 0
run "4 / 2 != 2" 0
run "4 / 2 == 2" 1
run "a:=10; a!=11;" 1

run "true;" 1
run "false;" 0
run "1; return false; 3 * 4;" 0

run "1 < 20;" 1
run "20 < 20" 0
run "20 < 1" 0
run "1 > 20;" 0
run "20 > 1" 1

run "4 <= 2" 0
run "10 <= 10" 1
run "5 >= 10" 0

run "1 || 0" 1
run "(1 < 20) || false" 1
run "true || 1" 1
run "1 && 0" 0
run "(1 < 20) && false" 0
run "true && 1" 1

run "!0" 1
run "!1" 0
run "!true" 0
run "!false" 1
run "!(20 >= 10*3)" 1


echo "OK"