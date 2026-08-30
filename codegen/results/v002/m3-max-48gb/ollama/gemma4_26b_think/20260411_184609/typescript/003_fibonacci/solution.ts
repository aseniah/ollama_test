/**
 * This program calculates and prints the Fibonacci sequence up to a given integer N.
 * The sequence starts: 1, 1, 2, 3, 5, ...
 * Each Fibonacci number in the sequence that is less than or equal to N is printed on a new line.
 * If N < 1, the program prints nothing.
 */

import * as process from 'process';

function solve(): void {
    // Retrieve the first command-line argument
    const arg = process.argv[2];

    // If no argument is provided, terminate
    if (arg === undefined) {
        return;
    }

    let N: bigint;

    try {
        // Use BigInt to handle potentially very large numbers without precision loss
        // BigInt is suitable for handling Fibonacci numbers which grow exponentially
        N = BigInt(arg);
    } catch (error) {
        // If the argument is not a valid integer, terminate
        return;
    }

    // Requirement: If N < 1, print nothing.
    if (N < 1n) {
        return;
    }

    // Initialize the first two numbers of the Fibonacci sequence
    let a = 1n;
    let b = 1n;

    // Print the first number in the sequence (1)
    // We use .toString() because BigInts append an 'n' when interpolated in some environments
    console.log(a.toString());

    /**
     * The sequence is 1, 1, 2, 3, 5...
     * We have already printed the first '1'.
     * We now iterate through the sequence as long as the current number 'b' 
     * does not exceed N.
     */
    while (b <= N) {
        console.log(b.toString());
        
        // Calculate the next Fibonacci number
        const next = a + b;
        
        // Shift the window: the current 'b' becomes the new 'a',
        // and the newly calculated 'next' becomes the new 'b'.
        a = b;
        b = next;
    }
}

// Execute the solver
solve();