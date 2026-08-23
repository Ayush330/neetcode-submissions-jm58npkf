class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        List<List<Integer>> ans = new ArrayList<>();
        Map<String, Boolean> map = new HashMap<>();
        for(int i=0; i<nums.length-2; i++){
            int target = -nums[i];
            int low = i+1;
            int high = nums.length-1;
            while(low < high){
                int sum = nums[low]+nums[high];
                if(sum==target){
                    ArrayList<Integer> data = new ArrayList<>(Arrays.asList(nums[i], nums[high], nums[low]));
                    Boolean exists = map.getOrDefault(data.toString(), false);
                    if(!exists){
                        ans.add(Arrays.asList(nums[i], nums[high], nums[low]));
                        map.put(data.toString(), true);
                    }
                    low++;
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
