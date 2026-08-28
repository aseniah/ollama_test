/**
 * This program prints the Fibonacci sequence starting from 1, 1, 2, 3, 5...
 * up to and including the largest Fibonacci number that does not exceed N.
 */

function main() {
    // Retrieve the argument N from the command line
    const arg = process.argv[2];

    if (arg === undefined) {
        return;
    }

    // Use BigInt to handle potentially very large Fibonacci numbers 
    // beyond the safe integer limit of JavaScript (2^53 - 1).
    let n: bigint;
    try {
        n = BigInt(arg);
    } catch (e) {
        // If input is not a valid integer, we simply exit.
        return;
    }

    // If N < 1, print nothing as per requirements.
    if (n < 1n) {
        return;
    }

    let current = 1n;
    let next = 1n;

    // Generate Fibonacci sequence
    // Starting sequence: 1, 1, 2, 3, 5...
    while (current <= n) {
        process.stdout.write(current.toString() + '\n');
        
        const temp = current + next;
        current = next;
        next = temp;
    }
}

main();