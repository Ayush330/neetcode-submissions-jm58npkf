func minDistance(word1 string, word2 string) int {
    return helper(word1, word2, 0, 0)   
}


func helper(word1, word2 string, i, j int) int{
  if i == len(word1) {
        return len(word2) - j 
    }
    if j == len(word2) {
        return len(word1) - i 
    }

   if word1[i] == word2[j]{
    return helper(word1, word2, i+1, j+1)
   }else{
        insert := helper(word1, word2, i, j+1)
        delet := helper(word1, word2, i+1, j)
        replace := helper(word1, word2, i+1, j+1)
        return min(insert, min(delet, replace)) + 1
   }
}


func min(a, b int)int{
    if a > b{
        return b
    }else{
        return a
    }
}