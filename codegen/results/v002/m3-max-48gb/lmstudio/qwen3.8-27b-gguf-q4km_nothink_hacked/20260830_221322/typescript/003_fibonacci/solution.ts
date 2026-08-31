// Accept a single integer argument N from the command line
// Print all Fibonacci numbers up to and including the largest Fibonacci number that does not exceed N.
// Start the sequence: 1, 1, 2, 3, 5, ...
// If N < 1, print nothing.

const N = parseInt(process.argv[2], 10);

if (isNaN(N) || N < 1) {
    // Print nothing
    process.exit(0);
}

// Generate Fibonacci numbers up to N
let prev = 1;
let curr = 1;

// Print the first two 1s
console.log(1);

if (N >= 1) {
    // We already printed the first 1
    if (N >= 1) {
        // Print the second 1 if N >= 1 (which it is)
        // Actually, let's restructure. Start with a = 1, b = 1
        // Print a, then compute next
    }
}

// Let me restructure more cleanly
let a = 1;
let b = 1;

// Print first Fibonacci number
if (a <= N) {
    console.log(a);
}

// Now generate subsequent Fibonacci numbers
while (b <= N) {
    console.log(b);
    const next = a + b;
    a = b;
    b = next;
}