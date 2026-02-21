class Solution {
    public List<String> generateParenthesis(int n) {
        List<String> result = new ArrayList<String>();
        helper("", 0,0, n, result);
        return result;
    }

    private void helper(String generated, int open, int close, int n, List<String> result) {
        if (generated.length() == 2*n) {
            result.add(generated);
            return;
        }

        if (open < n){
            helper(generated + "(", open+1, close, n, result);
        }

        if (close < open) {
            helper(generated + ")",open, close+1, n, result);
        }
    }
}