func fib(n int) int {
    
    var f, s, sum int = 0, 1, 0
    if n == 0{
        s=0
    }
    i := 2
    for i <= n{
        sum = f+s
        f = s
        s = sum
        i++
    }
    return s
}