class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> map = new HashMap<>();
        for(String str : strs){
            char[] input = str.toCharArray();
            Arrays.sort(input);
            String s = String.valueOf(input);
            map.computeIfAbsent(s, k-> new ArrayList<>()).add(str);
        }
        return new ArrayList<>(map.values());
    }
}
