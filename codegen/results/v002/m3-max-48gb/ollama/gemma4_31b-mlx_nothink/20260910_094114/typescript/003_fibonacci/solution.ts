/**
 * This program prints the Fibonacci sequence starting from 1, 1, 2...
 * up to and including the largest Fibonacci number that does not exceed N.
 */

function solve() {
    // Read the first argument provided after the node/tsx command
    const arg = process.argv[2];

    if (arg === undefined) {
        return;
    }

    const N = parseInt(arg, 10);

    // If N is not a number or N < 1, print nothing as per requirements
    if (isNaN(N) || N < 1) {
        return;
    }

    let a = 1;
    let b = 1;

    // The sequence starts with 1, 1...
    // First number in sequence
    console.log(a);

    // If N is 1, we only print the first two 1s (or just one if strictly following logic)
    // But the sequence provided in prompt is 1, 1, 2...
    // Let's handle the sequence generation iteratively.
    
    // Handle the second '1' specifically
    if (b <= N) {
        console.log(b);
    } else {
        // If N was 1, we've already printed the first 1. 
        // Actually, if N=1, we should print 1, 1.
        // Let's refine the loop.
    }

    // Generate subsequent numbers
    let next = a + b;
    while (next <= N) {
        console.log(next);
        a = b;
        b = next;
        next = a + b;
    }
}

// Overwriting the logic slightly to ensure perfect sequence starting 1, 1...
function solveCorrected() {
    const arg = process.argv[2];
    if (arg === undefined) return;

    const N = parseInt(arg, 10);
    if (isNaN(N) || N < 1) return;

    let curr = 1;
    let next = 1;

    // We use a loop that prints 'curr' and calculates the 'next'
    // The sequence is: 1, 1, 2, 3, 5...
    while (curr <= N) {
        console.log(curr);
        
        // Calculate next Fibonacci number
        const temp = curr + next;
        curr = next;
        next = temp;
    }
}

// To strictly follow "1, 1, 2, 3, 5", if N=1, output should be:
// 1
// 1
// If N=2:
// 1
// 1
// 2

// The corrected version:
// Iteration 1: curr=1, next=1. Print 1. New curr=1, next=2.
// Iteration 2: curr=1, next=2. Print 1. New curr=2, next=3.
// Iteration 3: curr=2, next=3. Print 2. New curr=3, next=5.
// This matches the requirement perfectly.

solveCorrected();