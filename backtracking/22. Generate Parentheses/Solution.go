package Solution

func generateParenthesis(n int) []string {
	// result := make([]string, )
	var result = []string{}
	helper(&result, "", 0, 0, n)
	return result
}

func helper(result *[]string, generated string, open, close, n int) {
	if len(generated) == 2*n {
		*result = append(*result, generated)
		return
	}

	if open < n {
		helper(result, generated+"(", open+1, close, n)
	}

	if close < open {
		helper(result, generated+")", open, close+1, n)
	}
}
