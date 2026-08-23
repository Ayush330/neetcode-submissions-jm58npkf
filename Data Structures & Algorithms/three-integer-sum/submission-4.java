class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        List<List<Integer>> ans = new ArrayList<>();
        for(int i=0; i<nums.length-2; i++){
            if(i > 0 && nums[i]==nums[i-1]) continue;
            int target = -nums[i];
            int low = i+1;
            int high = nums.length-1;
            while(low < high){
                int sum = nums[low]+nums[high];
                if(sum==target){
                    ArrayList<Integer> data = new ArrayList<>(Arrays.asList(nums[i], nums[high], nums[low]));
                    ans.add(Arrays.asList(nums[i], nums[high], nums[low]));
                    while(low+1 < high && nums[low]==nums[low+1]) low++;
                    while(high-1 > low && nums[high]==nums[high-1]) high--;
                    low++;
                    high--;
                }else if(sum < target){
                    low++;
                }else{
                    high--;
                }
            }
        }
        return ans;
    }
}
