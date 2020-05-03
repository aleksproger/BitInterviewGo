 const (
    r = iota
    l
    u
    d
)

func generateMatrix(A int )  ([][]int) {
    mat := make([][]int, A)
    for i := 0; i < A; i++ {
        mat[i] = make([]int, A)
    }
    direction := r
    cc := 0
    cr := 0
    for i := 1; i <= A*A; i++ {
        if cc < len(mat[0]) && cr < len(mat[0]) && mat[cr][cc] == 0 {
            mat[cr][cc] = i
            switch direction {
                case r:
                    if cc + 1 >= len(mat[0]) || mat[cr][cc+1] != 0 {
                        direction = d
                        cr += 1
                    } else {
                        cc += 1
                    }
                case d:
                    if cr + 1 >= len(mat[0]) || mat[cr+1][cc] != 0 {
                        direction = l
                        cc -= 1
                    } else {
                        cr += 1
                    }
                case l:
                    if cc - 1 < 0 || mat[cr][cc-1] != 0 {
                        direction = u
                        cr -= 1
                    } else {
                        cc -= 1
                    }
                case u:
                    if cr - 1 < 0 || mat[cr-1][cc] != 0 {
                        direction = r
                        cc += 1
                    } else {
                        cr -= 1
                    }
            }
        } else {
            return mat
        }
    }
    return mat
}
