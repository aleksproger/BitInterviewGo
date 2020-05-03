//https://www.interviewbit.com/problems/min-steps-in-infinite-grid/

func coverPoints(A []int , B []int )  (n int) {
    length := len(A)
    if len(B) < length {
        length = len(B)
    }
    n = 0
    for i := 0; i < length - 1; i++ {
        x0 := A[i]; y0 := B[i]
        x1 := A[i+1]; y1 := B[i+1]
        x := abs(x0 - x1)
        y := abs(y0 - y1)
        n += max(x,y)
    }
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    } 
    return a
}
