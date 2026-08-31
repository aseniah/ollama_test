// Read command-line argument
const n = parseInt(process.argv[2], 10);

// If N < 1, print nothing
if (isNaN(n) || n < 1) {
    process.exit(0);
}

// Generate Fibonacci numbers up to and including the largest Fibonacci number that does not exceed N
// Sequence: 1, 1, 2, 3, 5, ...
const fibs: number[] = [];

// If N >= 1, at least the first two 1s should be considered if they don't exceed N
if (n >= 1) {
    fibs.push(1);
    if (n >= 1) {
        fibs.push(1);
    }
}

// Continue generating Fibonacci numbers while the next number doesn't exceed N
let a = 1;
let b = 1;
while (true) {
    const next = a + b;
    if (next > n) break;
    fibs.push(next);
    a = b;
    b = next;
}

// Print each Fibonacci number on its own line
for (const f of fibs) {
    console.log(f);
}