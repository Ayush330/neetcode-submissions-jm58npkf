func characterReplacement(s string, k int) int {
    left := 0
    maxLength := 0
    maxFrequency := 0 // Track the frequency of the most common char in the window
    freqMap := make(map[byte]int)

    // 1. The main loop iterates through the string with the `right` pointer.
    for right := 0; right < len(s); right++ {
        
        // 2. Expand the window: Add the new character at `right`.
        charRight := s[right]
        freqMap[charRight]++
        maxFrequency = max(maxFrequency, freqMap[charRight])

        // 3. Check if the window is now invalid.
        windowLength := right - left + 1
        if (windowLength - maxFrequency) > k {
            // If invalid, shrink the window from the left.
            charLeft := s[left]
            freqMap[charLeft]--
            left++
        }

        // 4. The current window is guaranteed to be valid or just became valid.
        // Update the maxLength with the current window's size.
        maxLength = max(maxLength, right - left + 1)
    }
    return maxLength
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}