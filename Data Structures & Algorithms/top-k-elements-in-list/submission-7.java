class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        Map<Integer, Integer> map = new HashMap<>();
        for(int num : nums){
            map.put(num, map.getOrDefault(num, 0)+1);
        }
        // Create Bucket
        List<Integer>[] bucket = new ArrayList[nums.length+1];
        for(Map.Entry<Integer, Integer> m : map.entrySet()){
            if(bucket[m.getValue()] == null){
                bucket[m.getValue()] = new ArrayList<>();
            }
            bucket[m.getValue()].add(m.getKey());
        }
        int []res = new int[k];
        int count = 0;
        for(int i=bucket.length-1; i>=0; i--){
            if(bucket[i] == null){
                continue;
            }
            for(int j=0; j<bucket[i].size(); j++){
                if(count == k){
                    return res;
                }
                res[count] = bucket[i].get(j);
                count++;
            }
        }
        return res;
    }
}
