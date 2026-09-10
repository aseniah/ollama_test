/**
 * This program prints all Fibonacci numbers up to and including the 
 * largest Fibonacci number that does not exceed the given integer N.
 * 
 * The Fibonacci sequence used here starts with: 1, 1, 2, 3, 5, ...
 */

function main() {
    // Get the argument from the command line
    const arg = process.argv[2];

    if (arg === undefined) {
        return;
    }

    let n;
    try {
        // Use BigInt to handle potentially very large numbers provided by the user
        n = BigInt(arg);
    } catch (e) {
        // If the argument is not a valid integer, we simply exit
        return;
    }

    // If N < 1, print nothing
    if (n < 1n) {
        return;
    }

    // Initialize the first two numbers of the sequence
    let current = 1n;
    let next = 1n;

    /**
     * Fibonacci sequence logic:
     * We print the current number, then calculate the next one in the sequence.
     * We stop once the current number to be printed exceeds N.
     */
    
    // The problem specifies starting with 1, 1...
    // Print the first 1
    console.log(current.toString());

    // Loop to calculate and print subsequent numbers while they are <= N
    while (next <= n) {
        console.log(next.toString());
        
        const temp = current + next;
        current = next;
        next = temp;
    }
}

// Execute the program
main();