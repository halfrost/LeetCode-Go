package leetcode

func numMusicPlaylists(N int, L int, K int) int {
	dp, mod := make([][]int, L+1), 1000000007
	for i := 0; i < L+1; i++ {
		dp[i] = make([]int, N+1)
	}
	dp[0][0] = 1
	for i := 1; i <= L; i++ {
		for j := 1; j <= N; j++ {
			dp[i][j] = (dp[i-1][j-1] * (N - (j - 1))) % mod
			if j > K {
				dp[i][j] = (dp[i][j] + (dp[i-1][j]*(j-K))%mod) % mod
			}
		}
	}
	return dp[L][N]
}
