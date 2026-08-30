// Read the command-line argument N
const nStr = process.argv[2];
const N = parseInt(nStr, 10);

// If N < 1 or invalid, print nothing
if (isNaN(N) || N < 1) {
    process.exit(0);
}

// Generate Fibonacci numbers up to and including the largest that does not exceed N
let a = 1;
let b = 1;

// Print the first Fibonacci number
if (a <= N) {
    console.log(a);
}

// If N is 1, we only have the first 1
if (N <= 1) {
    process.exit(0);
}

// Print the second Fibonacci number if it doesn't exceed N
if (b <= N) {
    console.log(b);
}

// Generate subsequent Fibonacci numbers
let next = a + b;
while (next <= N) {
    console.log(next);
    a = b;
    b = next;
    next = a + b;
}