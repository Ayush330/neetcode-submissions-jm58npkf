class Solution {
    public int[] productExceptSelf(int[] nums) {
        int[] pre = new int[nums.length];
        int[] post = new int[nums.length];
        int start = 1;
        for(int i=0; i<nums.length; i++){
            pre[i] = start * nums[i];
            start = pre[i];
        }
        start = 1;
        for(int i=nums.length-1; i>=0; i--){
            post[i] = start * nums[i];
            start = post[i];
        }
        // 2 <= nums.length <= 100,000
        int[] ans = new int[nums.length];
        for(int i=0; i<nums.length; i++){
            if(i==0){
                ans[i] = post[i+1];
            }else if(i==nums.length-1){
                ans[i] = pre[i-1];
            }else{
                ans[i] = pre[i-1]*post[i+1];
            }
        }
        return ans;
    }
}  
