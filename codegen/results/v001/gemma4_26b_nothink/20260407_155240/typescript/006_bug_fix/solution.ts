/**
 * The original code had two bugs:
 * 1. The loop started at 0, causing the first element to be "FizzBuzz" (0 % 15 === 0).
 * 2. The loop ended before reaching n (i < n), so it missed the final number.
 * 
 * Fixed version:
 * - Loop from 1 to n inclusive.
 * - Correct logic for Fizz, Buzz, and FizzBuzz.
 */

function fizzbuzz(n: number): string[] {
    const result: string[] = [];
    // Start loop at 1 and include n
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