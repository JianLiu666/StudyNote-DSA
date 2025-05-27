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

type input = {
    nums: number[];
    target: number;
}
    
type output = {
    ans: number[];
}

type testdata = {
    input: input;
    output: output;
}

const tds: testdata[] = [
    {
        input: {
            nums: [2, 7, 11, 15],
            target: 9,
        },
        output: {
            ans: [0, 1],
        },
    },
    {
        input: {
            nums: [3, 2, 4],
            target: 6,
        },
        output: {
            ans: [1, 2],
        },
    }, 
    {
        input: {
            nums: [3, 3],
            target: 6,
        },
        output: {
            ans: [0, 1],
        },
    }
]

if (require.main === module) {
    for (const td of tds) {
        const result = twoSum(td.input.nums, td.input.target);
        assert.deepStrictEqual(result, td.output.ans, `Test failed: expected ${td.output.ans}, but got ${result}`);
    }
}