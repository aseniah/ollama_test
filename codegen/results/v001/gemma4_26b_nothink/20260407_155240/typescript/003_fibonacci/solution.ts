/**
 * Fibonacci sequence generator script.
 * 
 * Usage: tsx script.ts <N>
 * Output: All Fibonacci numbers up to N, one per line.
 * Sequence starts: 1, 1, 2, 3, 5...
 */

import process from 'node:process';

function main() {
    // Get the argument from command line
    const arg = process.argv[2];

    // Validate that an argument exists
    if (arg === undefined) {
        return;
    }

    // Parse the argument to an integer
    const n = parseInt(arg, 10);

    // If N < 1, the problem specifies to print nothing.
    // Also handle cases where parsing fails (NaN).
    if (isNaN(n) || n < 1) {
        return;
    }

    // Fibonacci sequence logic
    // Starting numbers: 1, 1
    let a = 1;
    let b = 1;

    // First number in sequence
    if (a <= n) {
        process.stdout.write(a.toString() + '\n');
    }

    // Second number in sequence
    if (b <= n) {
        process.stdout.write(b.toString() + '\n');
    }

    // Calculate subsequent numbers
    while (true) {
        const next = a + b;
        
        // Stop if the next Fibonacci number exceeds N
        if (next > n) {
            break;
        }

        process.stdout.write(next.toString() + '\n');

        // Update values for next iteration
        a = b;
        b = next;
    }
}

// Execute the main function
main();