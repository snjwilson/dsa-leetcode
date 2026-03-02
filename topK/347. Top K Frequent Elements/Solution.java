class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        Map<Integer, Integer> mp = new HashMap<>();
        for (int num: nums) {
            mp.put(num, mp.getOrDefault(num, 0)+1);
        }

        PriorityQueue<Integer> pq = new PriorityQueue<>((a,b) -> mp.get(a) - mp.get(b));
        for (int currKey: mp.keySet()){
            pq.add(currKey);
            if (pq.size() > k) {
                pq.poll();
            }
        }
        
        return pq.stream().mapToInt(i -> i).toArray();
    }
}