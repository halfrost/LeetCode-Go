package leetcode

// 解法一 DFS
func solveNQueens(n int) [][]string {
	col, dia1, dia2, row, res := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1), []int{}, [][]string{}
	putQueen(n, 0, &col, &dia1, &dia2, &row, &res)
	return res
}

// 尝试在一个n皇后问题中, 摆放第index行的皇后位置
func putQueen(n, index int, col, dia1, dia2 *[]bool, row *[]int, res *[][]string) {
	if index == n {
		*res = append(*res, generateBoard(n, row))
		return
	}
	for i := 0; i < n; i++ {
		// 尝试将第index行的皇后摆放在第i列
		if !(*col)[i] && !(*dia1)[index+i] && !(*dia2)[index-i+n-1] {
			*row = append(*row, i)
			(*col)[i] = true
			(*dia1)[index+i] = true
			(*dia2)[index-i+n-1] = true
			putQueen(n, index+1, col, dia1, dia2, row, res)
			(*col)[i] = false
			(*dia1)[index+i] = false
			(*dia2)[index-i+n-1] = false
			*row = (*row)[:len(*row)-1]
		}
	}
	return
}

func generateBoard(n int, row *[]int) []string {
	board := []string{}
	res := ""
	for i := 0; i < n; i++ {
		res += "."
	}
	for i := 0; i < n; i++ {
		board = append(board, res)
	}
	for i := 0; i < n; i++ {
		tmp := []byte(board[i])
		tmp[(*row)[i]] = 'Q'
		board[i] = string(tmp)
	}
	return board
}

// 解法二 二进制操作法
// class Solution
// {
//     int n;
//     string getNq(int p)
//     {
//         string s(n, '.');
//         s[p] = 'Q';
//         return s;
//     }
//     void nQueens(int p, int l, int m, int r, vector<vector<string>> &res)
//     {
//         static vector<string> ans;
//         if (p >= n)
//         {
//             res.push_back(ans);
//             return ;
//         }
//         int mask = l | m | r;
//         for (int i = 0, b = 1; i < n; ++ i, b <<= 1)
//             if (!(mask & b))
//             {
//                 ans.push_back(getNq(i));
//                 nQueens(p + 1, (l | b) >> 1, m | b, (r | b) << 1, res);
//                 ans.pop_back();
//             }
//     }
// public:
//     vector<vector<string> > solveNQueens(int n)
//     {
//         this->n = n;
//         vector<vector<string>> res;
//         nQueens(0, 0, 0, 0, res);
//         return res;
//     }
// };
