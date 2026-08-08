class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        HashMap<String, ArrayList<String>> map = new HashMap<>();
        for(int i=0; i<strs.length; i++){
            char[] input = strs[i].toCharArray();
            Arrays.sort(input);
            String s = String.valueOf(input);
            ArrayList<String> val = map.getOrDefault(s, new ArrayList<>());
            val.add(strs[i]);
            map.put(s, val);
        }
        return new ArrayList<>(map.values());
    }
}
