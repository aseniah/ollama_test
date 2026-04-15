/**
 * Fibonacci Sequence Generator
 * 
 * This program prints the Fibonacci sequence starting from 1, 1, 2...
 * up to the largest Fibonacci number that does not exceed the provided 
 * integer N.
 */

function main() {
    // Get the argument from the command line
    const arg = process.argv[2];

    if (arg === undefined) {
        return;
    }

    // Parse N as an integer
    const n = parseInt(arg, 10);

    // Handle cases where N is not a number or N < 1
    if (isNaN(n) || n < 1) {
        return;
    }

    let a = 1;
    let b = 1;

    // First number in the sequence
    console.log(a);

    // If N is at least 1, and since we start 1, 1, 2... 
    // we need to handle the second '1' explicitly or via loop.
    // The sequence defined is 1, 1, 2, 3, 5...
    
    // If N is 1, we only print the first 1 and the second 1.
    // But our loop logic needs to be precise.
    
    // Special case for the first '1' is already done.
    // Now handle the second '1' and subsequent numbers.
    
    // The loop continues as long as the current Fibonacci number (b) <= N
    while (b <= n) {
        console.log(b);
        
        // Calculate next Fibonacci number
        const next = a + b;
        a = b;
        b = next;
    }
}

main();