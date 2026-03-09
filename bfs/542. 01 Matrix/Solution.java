import java.util.*;

class Solution {
    public int[][] updateMatrix(int[][] mat) {
        int m = mat.length;
        int n = mat[0].length;
        Queue<int[]> q = new LinkedList<>();
        for (int i = 0; i < m;i++) {
            for (int j=0;j<n;j++) {
                if (mat[i][j]!=0) {
                    mat[i][j]=Integer.MAX_VALUE;
                    continue;
                }
                q.add(new int[]{i,j});
            }
        }
        int[][] dirs={{0,1},{1,0},{0,-1},{-1,0}};
        while(q.size()!=0){
            int[] pos=q.poll();
            for(int[] dir:dirs){
                int r=pos[0]+dir[0];
                int c=pos[1]+dir[1];
                if (r>=0&&c>=0&&r<m&&c<n&&mat[r][c]>mat[pos[0]][pos[1]]+1){
                    mat[r][c]=mat[pos[0]][pos[1]]+1;
                    q.add(new int[]{r,c});
                }
            }
        }
        return mat;
    }
}