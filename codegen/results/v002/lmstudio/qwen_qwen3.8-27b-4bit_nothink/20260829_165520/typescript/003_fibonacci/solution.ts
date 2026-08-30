#!/usr/bin/env tsx

const arg = process.argv[2];
if (!arg) {
    process.exit(0);
}

const N = parseInt(arg, 10);

if (isNaN(N) || N < 1) {
    process.exit(0);
}

// Print all Fibonacci numbers up to and including the largest <= N
// Sequence: 1, 1, 2, 3, 5, 8, ...

let a = 1;
let b = 1;

while (a <= N) {
    console.log(a);
    const next = a + b;
    a = b;
    b = next;
}