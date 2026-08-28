// Get N from command line argument
const n = parseInt(process.argv[2], 10);

// If N < 1, print nothing
if (n < 1 || isNaN(n)) {
    process.exit(0);
}

// Generate Fibonacci numbers starting: 1, 1, 2, 3, 5, ...
let a = 1;
let b = 1;

// Print first number (F(1) = 1)
console.log(a);

// Continue while b <= n
while (b <= n) {
    console.log(b);
    const next = a + b;
    a = b;
    b = next;
}