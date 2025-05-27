import assert from "assert";

// Time Complexity: O(n)
// Space Complexity: O(n)
function twoSum(nums: number[], target: number): number[] {
    const numMap = new Map<number, number>();
    
    for (let i = 0; i < nums.length; i++) {
        const complement = target - nums[i];
        
        if (numMap.has(complement)) {
            return [numMap.get(complement)!, i];
        }
        
        numMap.set(nums[i], i);
    }
    
    return [];
}


if (require.main === module) {
    const result1 = twoSum([2, 7, 11, 15], 9);
    assert.deepStrictEqual(result1, [0, 1], "Test 1 failed");

    const result2 = twoSum([3, 2, 4], 6);
    assert.deepStrictEqual(result2, [1, 2], "Test 2 failed");
    
    const result3 = twoSum([3, 3], 6);
    assert.deepStrictEqual(result3, [0, 1], "Test 3 failed");
}