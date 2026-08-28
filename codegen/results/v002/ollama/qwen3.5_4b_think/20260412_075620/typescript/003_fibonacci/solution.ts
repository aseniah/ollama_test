// Get the command-line argument
const args = process.argv.slice(2);

if (args.length !== 1) {
    // Expected exactly one argument N
    process.exit(1);
}

// Parse the integer argument
const N = parseInt(args[0], 10);

if (isNaN(N) || N < 1) {
    // No Fibonacci numbers to print for N < 1
    process.exit(0);
}

// Generate Fibonacci numbers starting 1, 1, 2, 3, 5, ...
let a = 1;
let b = 1;

while (a <= N) {
    console.log(a);
    const next = a + b;
    a = b;
    b = next;
}