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

function arraysEqual(a: number[], b: number[]): boolean {
    return a.length === b.length && a.every((val, index) => val === b[index]);
}

function assert(condition: boolean, message: string): void {
    if (!condition) {
        throw new Error(`Assertion failed: ${message}`);
    }
}

if (require.main === module) {
    const result1 = twoSum([2, 7, 11, 15], 9);
    assert(arraysEqual(result1, [0, 1]), `Test 1 failed: expected [0, 1], got [${result1}]`);

    const result2 = twoSum([3, 2, 4], 6);
    assert(arraysEqual(result2, [1, 2]), `Test 2 failed: expected [1, 2], got [${result2}]`);
    
    const result3 = twoSum([3, 3], 6);
    assert(arraysEqual(result3, [0, 1]), `Test 3 failed: expected [0, 1], got [${result3}]`);
}