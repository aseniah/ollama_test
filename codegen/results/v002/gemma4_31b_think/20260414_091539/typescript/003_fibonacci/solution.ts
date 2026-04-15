/**
 * This program generates the Fibonacci sequence starting from 1, 1, 2, 3, 5...
 * and prints all elements up to and including the largest Fibonacci number 
 * that does not exceed the provided integer argument N.
 */

function main() {
    // Access the command-line argument N
    const arg = process.argv[2];

    // If no argument is provided, do nothing
    if (arg === undefined) {
        return;
    }

    let N: bigint;
    try {
        // Use BigInt to support values larger than Number.MAX_SAFE_INTEGER
        N = BigInt(arg);
    } catch (error) {
        // If the argument is not a valid integer, exit silently or ignore
        return;
    }

    // Requirement: If N < 1, print nothing
    if (N < 1n) {
        return;
    }

    // The sequence starts with 1, 1, 2, 3, 5...
    let current = 1n;
    let next = 1n;

    /**
     * Fibonacci Logic:
     * Starting from current=1, next=1:
     * 1. Print current if current <= N.
     * 2. Calculate the next Fibonacci number: temp = current + next.
     * 3. Shift current to next and next to temp.
     */
    while (current <= N) {
        process.stdout.write(current.toString() + '\n');
        
        const temp = current + next;
        current = next;
        next = temp;
    }
}

// Run the program
main();