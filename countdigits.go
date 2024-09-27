package main

func countDigits(n int) int {
    count := 0
    if n < 0 {
        n = -n
    }
    if n == 0 {
        return 1
    }
    for n != 0 {
        n /= 10
        count++
    }

    return count
}
