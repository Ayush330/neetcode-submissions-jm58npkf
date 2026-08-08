class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        Map<Integer, Integer> map = new HashMap<>();
        for(int num : nums){
            map.put(num, map.getOrDefault(num, 0)+1);
        }
        // create an array
        List<List<Integer>> l = new ArrayList<>();
        for(Map.Entry<Integer, Integer> entry : map.entrySet()){
            l.add(new ArrayList<>(List.of(entry.getKey(), entry.getValue())));
        }
        l.sort((a, b) -> Integer.compare(b.get(1), a.get(1)));
        int[] res = new int[k];
        for(int i=0; i<k; i++){
            res[i] = l.get(i).get(0);
        }
        return res;
    }
}
