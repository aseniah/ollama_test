/**
 * This script accepts a single integer argument N from the command line
 * and prints the Fibonacci sequence starting from 1, 1, 2, 3, 5...
 * up to the largest Fibonacci number that does not exceed N.
 * 
 * Constraints:
 * - N < 1: Print nothing.
 * - One number per line.
 */

function main() {
    // Access the command-line arguments
    // process.argv[0] is the node executable
    // process.argv[1] is the script file path
    // process.argv[2] is the first user-provided argument
    const arg = process.argv[2];

    // If no argument is provided, exit silently
    if (arg === undefined) {
        return;
    }

    // Parse the argument as an integer
    const n = parseInt(arg, 10);

    // If N is not a valid number or N < 1, print nothing
    if (isNaN(n) || n < 1) {
        return;
    }

    /**
     * We initialize the sequence with the first two terms: 1 and 1.
     * 'current' represents the term to be printed in the current iteration.
     * 'nextTerm' represents the next number in the Fibonacci sequence.
     */
    let current = 1;
    let nextTerm = 1;

    /**
     * The loop continues as long as the current Fibonacci number
     * is less than or equal to the input N.
     */
    while (current <= n) {
        // Write the current number to stdout followed by a newline
        process.stdout.write(current.toString() + '\n');

        // Calculate the next term in the sequence
        const temp = current + nextTerm;
        
        // Shift values: the next number in the sequence becomes the current number
        // and the newly calculated sum becomes the next term.
        current = nextTerm;
        nextTerm = temp;
    }
}

// Run the main function
main();