/**
 * Problem: Print all Fibonacci numbers up to and including the largest 
 * Fibonacci number that does not exceed N.
 * Sequence starts: 1, 1, 2, 3, 5, ...
 * If N < 1, print nothing.
 */

import process from 'node:process';

function solve(): void {
    // Get command line arguments
    const args = process.argv.slice(2);

    if (args.length === 0) {
        return;
    }

    const input = args[0];
    const n = parseInt(input, 10);

    // If N is not a number or N < 1, print nothing as per requirements
    if (isNaN(n) || n < 1) {
        return;
    }

    // Special case for N >= 1
    // The sequence starts with 1, 1, 2...
    // We use BigInt to avoid overflow for very large N, 
    // though standard Number works for most practical integer inputs.
    let a: bigint = 1n;
    let b: bigint = 1n;
    const limit: bigint = BigInt(n);

    // Print the first number
    process.stdout.write(a.toString() + '\n');

    // Check if the second number (also 1) should be printed
    // Since the prompt asks for the sequence 1, 1, 2... and N >= 1,
    // the second 1 is always part of the sequence if N >= 1.
    process.stdout.write(b.toString() + '\n');

    // Calculate subsequent Fibonacci numbers
    while (true) {
        let next = a + b;
        if (next > limit) {
            break;
        }
        process.stdout.write(next.toString() + '\n');
        a = b;
        b = next;
    }
}

// Execute the solver
solve();