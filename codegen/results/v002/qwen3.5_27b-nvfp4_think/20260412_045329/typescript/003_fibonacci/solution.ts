#!/usr/bin/env tsx

const N = parseInt(process.argv[2], 10);

if (N < 1 || isNaN(N)) {
    process.exit(0);
}

let a = 1;
let b = 1;

// Print the first two Fibonacci numbers (both are 1)
console.log(a);
console.log(b);

// Generate and print remaining Fibonacci numbers up to N
while (true) {
    const next = a + b;
    if (next > N) {
        break;
    }
    console.log(next);
    a = b;
    b = next;
}