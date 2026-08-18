class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        List<Integer> [] bucket = new ArrayList[nums.length+1];
        Map<Integer, Integer> map = new HashMap<>();
        for(int el : nums){
            map.put(el, map.getOrDefault(el, 0)+1);
        }
        for(Map.Entry<Integer, Integer> entry : map.entrySet()){
            Integer key = entry.getKey();
            Integer value = entry.getValue();
            if(bucket[value]==null){
                bucket[value] = new ArrayList<>();
            }
            bucket[value].add(key);
        }

        int[] answer = new int[k];
        int track = 0;
        for(int i=bucket.length-1; i>=0 && track < k; i--){
            if(bucket[i]!=null){
                for(Integer z : bucket[i]){
                    answer[track++] = z;
                    if(track==k){
                        return answer;
                    }
                }
            }
        }
        return answer;
    }
}
