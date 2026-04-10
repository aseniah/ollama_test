/**
 * Fixed FizzBuzz implementation.
 * 
 * The bug in the original code was:
 * 1. The loop started at 0. In FizzBuzz, 0 is usually not included, or 
 *    if it were, 0 % 3 and 0 % 5 are both 0, resulting in "FizzBuzz" 
 *    being the first element, which is not the standard sequence.
 * 2. The loop condition 'i < n' excluded the number 15.
 * 
 * The standard FizzBuzz sequence for n=15 starts at 1 and ends at 15.
 */

function fizzbuzz(n: number): string[] {
    const result: string[] = [];
    // Start from 1 and include n
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

// Execute for n = 15 as requested
const n = 15;
fizzbuzz(n).forEach(v => console.log(v));