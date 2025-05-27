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

type input = {
    word1: string;
    word2: string;
}
    
type output = {
    ans: string;
}

type testdata = {
    input: input;
    output: output;
}

const tds: testdata[] = [
    {
        input: {
            word1: "abc",
            word2: "pqr",
        },
        output: {
            ans: "apbqcr",
        },
    },
    {
        input: {
            word1: "ab",
            word2: "pqrs",
        },
        output: {
            ans: "apbqrs",
        },
    },
    {
        input: {
            word1: "abcd",
            word2: "pq",
        },
        output: {
            ans: "apbqcd",
        },
    }
]

if (require.main === module) {
    for (const td of tds) {
        const result = mergeAlternately(td.input.word1, td.input.word2);
        assert.equal(result, td.output.ans, `Test failed: expected ${td.output.ans}, but got ${result}`);
    }
}

