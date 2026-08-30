/**
 * Fibonacci Sequence Generator
 * 
 * Task:
 * Accept a single integer N from the command line.
 * Print all Fibonacci numbers up to and including the largest Fibonacci number
 * that does not exceed N. One number per line. 
 * Sequence starts: 1, 1, 2, 3, 5, ...
 * If N < 1, print nothing.
 */

import process from 'node:process';

function main() {
    // Get the argument from command line
    const arg = process.argv[2];

    // Check if argument exists
    if (arg === undefined) {
        return;
    }

    // Parse the argument to a number
    const N = parseInt(arg, 10);

    // If N is not a number or N < 1, print nothing as per requirements
    if (isNaN(N) || N < 1) {
        return;
    }

    // Fibonacci sequence generation
    // We start with 1, 1 as per the prompt requirements
    let a = 1;
    let b = 1;

    // Print the first number in the sequence
    console.log(a);

    // If N is at least 1, the second 1 should also be printed if it's <= N
    // (In this case, N >= 1 is already checked)
    if (b <= N) {
        console.log(b);
    } else {
        // This case handles if N was 0, but we already returned for N < 1
        return;
    }

    // Calculate subsequent numbers
    while (true) {
        const next = a + b;
        
        // If the next number exceeds N, stop the loop
        if (next > N) {
            break;
        }

        console.log(next);

        // Update pointers for next iteration
        a = b;
        b = next;
    }
}

// Run the program
main();