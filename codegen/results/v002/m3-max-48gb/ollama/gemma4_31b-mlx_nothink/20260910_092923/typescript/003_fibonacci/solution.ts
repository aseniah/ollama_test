/**
 * This program calculates and prints Fibonacci numbers up to and including
 * the largest Fibonacci number that does not exceed a given integer N.
 * 
 * The sequence starts: 1, 1, 2, 3, 5, ...
 */

function printFibonacciUpToN(n: number): void {
    // If N < 1, the problem specifies to print nothing.
    if (n < 1) {
        return;
    }

    let prev = 1;
    let curr = 1;

    // The sequence starts with 1, 1. 
    // Since we know n >= 1, we print the first 1.
    console.log(prev);

    // If n is exactly 1, we only print the first 1 (though the 
    // sequence starts 1, 1, so we handle the second 1 below).
    // However, the standard Fibonacci sequence definition for this 
    // task requires us to handle the growth.
    
    // Print the second 1 if it doesn't exceed N.
    if (curr <= n) {
        console.log(curr);
    } else {
        return;
    }

    // Calculate subsequent numbers
    while (true) {
        const next = prev + curr;
        if (next > n) {
            break;
        }
        console.log(next);
        prev = curr;
        curr = next;
    }
}

function main() {
    // Get the argument from command line (process.argv[0] is node, [1] is script path)
    const arg = process.argv[2];

    if (arg === undefined) {
        return;
    }

    const n = parseInt(arg, 10);

    // Handle cases where the argument is not a valid number
    if (isNaN(n)) {
        return;
    }

    printFibonacciUpToN(n);
}

main();