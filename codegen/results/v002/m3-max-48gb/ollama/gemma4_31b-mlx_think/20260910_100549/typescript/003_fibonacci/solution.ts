import * as process from 'node:process';

/**
 * This program prints Fibonacci numbers up to and including the largest 
 * Fibonacci number that does not exceed a given integer N.
 * The sequence starts: 1, 1, 2, 3, 5, ...
 */
function solve() {
    // Retrieve the first command-line argument
    const inputArg = process.argv[2];

    // If no argument is provided, do nothing
    if (inputArg === undefined) {
        return;
    }

    try {
        // Parse input as BigInt to handle integers larger than Number.MAX_SAFE_INTEGER
        const N = BigInt(inputArg);

        // If N is less than 1, print nothing as per requirements
        if (N < 1n) {
            return;
        }

        // Initialize the first two numbers of the sequence
        let a = 1n;
        let b = 1n;

        // The sequence is specified to start: 1, 1, 2, 3, 5...
        // Since N >= 1, the first Fibonacci number (1) always fits.
        console.log(a.toString());

        // The second Fibonacci number (1) also always fits since N >= 1.
        console.log(b.toString());

        // Calculate subsequent Fibonacci numbers
        while (true) {
            const next = a + b;

            // If the next number exceeds N, we stop.
            if (next > N) {
                break;
            }

            // Print the number and update sequence pointers
            console.log(next.toString());
            a = b;
            b = next;
        }
    } catch (error) {
        // In case of invalid integer input, the program will fail silently
        // or can be adjusted to print an error. Based on the prompt, 
        // we assume the input is a valid integer.
    }
}

// Execute the solver
solve();