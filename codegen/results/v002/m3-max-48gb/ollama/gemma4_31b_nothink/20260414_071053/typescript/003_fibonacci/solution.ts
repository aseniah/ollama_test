/**
 * This program prints all Fibonacci numbers up to and including 
 * the largest Fibonacci number that does not exceed N.
 * 
 * Sequence start: 1, 1, 2, 3, 5, ...
 */

function main() {
    // Get the argument from the command line
    const arg = process.argv[2];

    if (arg === undefined) {
        return;
    }

    // Parse the argument as an integer
    const n = parseInt(arg, 10);

    // If N is not a number or N < 1, print nothing as per requirements
    if (isNaN(n) || n < 1) {
        return;
    }

    let a = 1;
    let b = 1;

    // Print the first number in the sequence
    console.log(a);

    // Handle the case where N is 1 separately to avoid double printing 1 
    // if only one 1 is intended, but the prompt specifies "1, 1, 2...", 
    // so we proceed to calculate the next.
    
    while (b <= n) {
        console.log(b);
        
        // Calculate next Fibonacci number
        const next = a + b;
        a = b;
        b = next;
    }
}

main();