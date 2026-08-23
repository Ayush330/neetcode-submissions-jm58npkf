class Solution {
    public int longestConsecutive(int[] nums) {
        Set<Integer> set = new HashSet<>();
        for(int num : nums){
            set.add(num);
        }
        int ans = 0;
        for(int num : set){
            int tmp = num;
            int curr = 0;
            if(!set.contains(tmp-1)){
                curr++;
                while(set.contains(tmp+1)){
                    curr++;
                    tmp++;
                }
                ans = Math.max(ans, curr);
            }
        }
        return ans;
    }
}
