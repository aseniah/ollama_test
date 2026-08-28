import * as fs from 'fs';

// Get the integer argument N
const argIndex = 1;
if (argIndex < process.argv.length && !isNaN(process.argv[argIndex])) {
    const n = parseInt(process.argv[argIndex], 10);

    // If N < 1, print nothing
    if (n < 1) {
        process.exit(0);
    }

    let a = 1;
    let b = 1;

    // Print the first two Fibonacci numbers (since we start with 1, 1)
    // We need to handle the case where N is exactly 1, 2, etc.
    
    while (a <= n) {
        console.log(a);
        
        // Calculate the next Fibonacci number
        const next = a + b;
        
        // If the next number exceeds N, stop. 
        // However, the loop condition 'a <= n' handles the final print,
        // so we must ensure we don't print numbers > N.
        // The sequence is generated: 1, 1, 2, 3, 5...
        // In each iteration, we print 'a', then compute 'next'.
        // If 'next' > n, the loop will terminate next time because 'a' becomes 'next'.
        
        a = next;
        b = next - a; // update b to the previous 'a'
        
        // A safer loop structure to avoid complex logic inside:
        // Reset and regenerate clearly.
    }
    
    // Let's rewrite the loop for clarity and correctness based on standard generation
    let fib1 = 1;
    let fib2 = 1;
    
    // We know the sequence starts: 1, 1, 2, 3, 5...
    // We iterate until the next number would exceed N.
    
    if (fib1 <= n) {
        console.log(fib1);
    }
    
    if (fib2 <= n) {
        console.log(fib2);
    }
    
    while (true) {
        let fibNext = fib1 + fib2;
        if (fibNext > n) {
            break;
        }
        console.log(fibNext);
        fib1 = fib2;
        fib2 = fibNext;
    }
}