import assert from "assert";

// Time Complexity: O(n)
// Space Complexity: O(n)
function mergeAlternately(word1: string, word2: string): string {
    let result = "";
    
    let i = 0;
    for (; i < word1.length && i < word2.length; i++) {
        result += word1[i] + word2[i];
    }

    if (i < word1.length) {
        result += word1.slice(i);
    }
    if (i < word2.length) {
        result += word2.slice(i);
    }

    return result;
};

if (require.main === module) {
    assert(mergeAlternately("abc", "pqr") === "apbqcr", "Test 1 failed");
    assert(mergeAlternately("ab", "pqrs") === "apbqrs", "Test 2 failed");
    assert(mergeAlternately("abcd", "pq") === "apbqcd", "Test 3 failed");
}