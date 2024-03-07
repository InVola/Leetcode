func tribonacci(n int) int {
    var f, s, t, sum int = 0, 1, 1, 1

    if n == 0{
        s = 0
    }

    if n == 1{
        s = 1
    }

    i :=2
    for i <= n{
        sum = f+s+t
        f = s
        s = t
        t = sum
        i++
    }

    return s
}