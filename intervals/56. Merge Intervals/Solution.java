class Solution {
    public int[][] merge(int[][] intervals) {
        if(intervals.length<=1) return intervals;
        List<int[]> arrList=new ArrayList<>();
        Arrays.sort(intervals,(a,b)->a[0]-b[0]);
        int[] currentInterval=intervals[0];
        arrList.add(currentInterval);
        for(int[] interval:intervals){
            int currentEnd=currentInterval[1];
            int intervalStart=interval[0];
            int intervalEnd=interval[1];
            if (currentEnd>=intervalStart){
                currentInterval[1]=Math.max(currentEnd,intervalEnd);
            } else{
                currentInterval=interval;
                arrList.add(currentInterval);
            }
        }
        return arrList.toArray(new int[arrList.size()][]);
    }
}