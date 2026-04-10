/**
 * Fixed FizzBuzz implementation.
 * The bug in the original code was:
 * 1. The loop started at 0, which caused "FizzBuzz" to be pushed first (0 % 15 === 0).
 * 2. The loop condition 'i < n' stopped at 14 instead of 15.
 * 
 * Correct logic:
 * - Loop from 1 to n (inclusive).
 * - Check divisibility by 3 and 5.
 */

function fizzbuzz(n: number): string[] {
    const result: string[] = [];
    // Start from 1 and go up to and including n
    for (let i = 1; i <= n; i++) {
        if (i % 15 === 0) {
            result.push("FizzBuzz");
        } else if (i % 3 === 0) {
            result.push("Fizz");
        } else if (i % 5 === 0) {
            result.push("Buzz");
        } else {
            result.push(String(i));
        }
    }
    return result;
}

// Execute for n=15 and print each value on a new line
const n = 15;
fizzbuzz(n).forEach(v => console.log(v));