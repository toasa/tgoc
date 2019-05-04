#!/bin/bash

run () {
    input=$1
    expected=$2

    go run main.go "${input}"
    gcc -o main.o main.s

    output=$(./main.o)

    if [ $output != $expected ]; then
        echo "${input}"
        echo "expected ${expected}, but got ${output}"
        exit 1
    else
        echo "${input} => ${output}"
    fi
}

# run "0;" 0
# run "20" 20
# run "255" 255

# run "5 + 3;" 8
# run "5 - 3" 2
# run "5 * 3;" 15
# run "15 / 3" 5
# run "15 % 3" 0

# run "5 + 3 + 10;" 18
# run "5 - 3 + 10" 12
# run "2 * 3 + 4" 10
# run "2 + 3 * 4" 14
# run "2 * 3 + 4 * 5;" 26
# run "2 + 3 * 4 + 5" 19
# run "6 / 3 * 10" 20
# run "12 / 4 + 5" 8
# run "6 + 30 / 10" 9

# run "(2 + 3) * 4" 20
# run "(2 + 3) * (4 + 5)" 45
# run "26 / (10 + 3)" 2

# run "a := 20 * 2; a" 40
# run "abc := 30 + 4 * 2; xyz := abc * 2; xyz;" 76
# run "a := 1; b := 1; c := a + b; d := b + c; e := c + d; e;" 5
# run "a := 1; b := 1; c := a + b; d := b + c; e := c + d; return e;" 5
# run "a := 10; a = 1 + 10; return a;" 11
# run "a := 10; a = a + 10; return a;" 20
# run "a := 5; b := a * a; b = 10 + b; return b" 35

# run "a := 2 * 3; b := a + 40;" 46

# run "+10" 10
# run "-2 + 6;" 4
# run "-(10 - 16)" 6
# run "-3*+5*-2" 30
# run "a := -20; return -a;" 20

# run "1 << 7" 128
# run "1024 >> 3" 128

# run "20 == 20" 1
# run "20 != 20" 0
# run "1 + 2 + 3 == 1 * 2 * 3" 1
# run "0 == 1" 0
# run "4 / 2 != 2" 0
# run "4 / 2 == 2" 1
# run "a:=10; a!=11;" 1

# run "true;" 1
# run "false;" 0
# run "1; return false;" 0

# run "1 < 20;" 1
# run "20 < 20" 0
# run "20 < 1" 0
# run "1 > 20;" 0
# run "20 > 1" 1

# run "4 <= 2" 0
# run "10 <= 10" 1
# run "5 >= 10" 0

# run "1 || 0" 1
# run "(1 < 20) || false" 1
# run "true || 1" 1
# run "1 && 0" 0
# run "(1 < 20) && false" 0
# run "true && 1" 1

# run "!0" 1
# run "!1" 0
# run "!true" 0
# run "!false" 1
# run "!(20 >= 10*3)" 1

# run "if true {20;}" 20
# run "if (1 * 3 <= 2 + 4) {return 10;}" 10
# run "if (1*3<=2+4) { return 10; } return 20;" 10
# run "if 1 * 3 <= 2 + 4 { return 20; } else { return 30; }" 20
# run "if false { return 20; } else { return 30; }" 30

# run "6 & 10" 2
# run "6 | 10" 14
# run "6 ^ 10" 12
# run "6 &^ 10" 4
# run "10 &^ 6" 8
# run "30 | 20 & 10" 30
# run "30 & 20 | 10" 30

# run "a := 1; a = 2 * a; a = 2 * a; a = 2 * a; a = 2 * a; a;" 16
# run "i := 1; for i < 200 { i = i + 2; } i;" 201
# run "a := 1; for a < 40000 { a = a * 2; } a;" 65536
# run "i := 1; a := 10 for i < 10 { a = a + i; i = i + 1;} a;" 55

# run "a := 10; for i := 0; i < 10; i = i + 1 { a = a + 1; } return a;" 20

run "var a int; a = 10; return a" 10
# run "var abc int = 200; return abc" 200
# run "var a int; a = 20; var b *int; b = &a; return *b" 20

echo "OK"