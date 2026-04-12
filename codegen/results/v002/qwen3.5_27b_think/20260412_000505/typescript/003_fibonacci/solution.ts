// Read the command line argument
const input = process.argv[2];

// Parse the integer
const N = parseInt(input, 10);

// If N < 1, print nothing
if (N < 1) {
    process.exit(0);
}

// Generate and print Fibonacci numbers up to N
let a = 1;
let b = 1;

while (a <= N) {
    console.log(a);
    const next = a + b;
    a = b;
    b = next;
}