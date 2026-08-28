/**
 * This program takes a single integer argument N from the command line
 * and prints all Fibonacci numbers in the sequence (starting 1, 1, 2, 3, 5, ...)
 * up to and including the largest Fibonacci number that does not exceed N.
 * 
 * If N < 1, the program prints nothing.
 * 
 * The program uses BigInt to ensure accuracy for very large values of N.
 */

import process from 'node:process';

function solve(): void {
    // Get the command-line arguments.
    const args = process.argv;

    // Ensure there is an argument provided.
    if (args.length < 3) {
        return;
    }

    const input = args[2];
    let n: bigint;

    // Attempt to parse the input as a BigInt.
    try {
        n = BigInt(input);
    } catch (err) {
        // If the input is not a valid integer, terminate.
        return;
    }

    // If N < 1, print nothing as per instructions.
    if (n < 1n) {
        return;
    }

    // The Fibonacci sequence as defined in the prompt starts: 1, 1, 2, 3, 5, ...
    // We initialize the first two terms of the sequence.
    let a = 1n;
    let b = 1n;

    /**
     * Logic:
     * The sequence is defined by F(1)=1, F(2)=1, and F(i) = F(i-1) + F(i-2).
     * We want to print all F(i) such that F(i) <= N.
     * Since N >= 1, the first 1 is always included.
     */
    
    // Print the first term.
    process.stdout.write(a.toString() + '\n');

    // Check if the second term (1) is within the limit N.
    // Since we already checked N >= 1, the second '1' is always included if N >= 1.
    if (b <= n) {
        process.stdout.write(b.toString() + '\n');

        // Prepare to iterate through the rest of the sequence.
        let prev = a;
        let curr = b;

        while (true) {
            const next = prev + curr;
            
            // If the next Fibonacci number is within the limit, print it and continue.
            if (next <= n) {
                process.stdout.write(next.toString() + '\n');
                prev = curr;
                curr = next;
            } else {
                // If the next number exceeds N, stop the loop.
                break;
            }
        }
    }
}

// Execute the solver.
solve();