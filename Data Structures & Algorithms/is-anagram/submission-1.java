class Solution {
    public boolean isAnagram(String s, String t) {
        char[] lookup = new char[26];
        for(int i=0; i<s.length(); i++){
            int posn = s.charAt(i) - 'a';
            lookup[posn]++;
        }
        for(int i=0; i<t.length(); i++){
            int posn = t.charAt(i) - 'a';
            lookup[posn]--;
        }
        // verify now
        for(int i=0; i<lookup.length; i++){
            if(lookup[i] != 0){
                return false;
            }
        }
        return true;
    }
}
