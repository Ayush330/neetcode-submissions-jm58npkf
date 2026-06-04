func minWindow(s string, t string) string {
    if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
        return ""
    }

    // 1. The Single Map (The Debt Ledger)
    charCount := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        charCount[t[i]]++
    }

    // 2. The Single Integer Counter
    missingChars := len(t) 

    // Window Pointers and Tracking
    start := 0
    minStart := 0
    minLen := len(s) + 1 // Set to an impossible large value initially

    // Phase 1: Aggressive Expand
    for end := 0; end < len(s); end++ {
        charRight := s[end]

        // If the count is > 0, we actually needed this character.
        // It pays off our missing character debt.
        if charCount[charRight] > 0 {
            missingChars--
        }
        
        // Unconditionally reduce the count in the map.
        // If it goes negative, it represents a surplus of that character.
        charCount[charRight]--

        // Phase 2: Defensive Shrink (Triggered when all debt is paid)
        for missingChars == 0 {
            
            // Record the current window if it's the smallest we've seen
            currentWindowLen := end - start + 1
            if currentWindowLen < minLen {
                minLen = currentWindowLen
                minStart = start
            }

            charLeft := s[start]
            
            // Reclaim the character as we shrink the window.
            charCount[charLeft]++

            // If the count goes ABOVE 0, we just dropped a required character.
            // Our debt increases, breaking the shrink loop.
            if charCount[charLeft] > 0 {
                missingChars++
            }

            start++
        }
    }

    // If minLen never changed, no valid window was found
    if minLen > len(s) {
        return ""
    }

    return s[minStart : minStart+minLen]
}