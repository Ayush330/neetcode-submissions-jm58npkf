class Solution {
    public int[] twoSum(int[] nums, int target) {
        int[] soln = new int[]{-1, -1};
        Map<Integer, Integer> map = new HashMap<>();
        for(int i=0; i<nums.length; i++){
            int comp = target - nums[i];
            if (map.get(comp) == null){
                map.put(nums[i], i);
            }else{
                soln[0] = map.get(comp);
                soln[1] = i;
                return soln;
            }
        }
        return soln;
    }
}
