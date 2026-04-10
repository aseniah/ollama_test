// Parse the integer argument from command line
const N = parseInt(process.argv[2], 10);

// If N is less than 1, print nothing
if (N < 1) {
    process.exit(0);
}

// Start with first two Fibonacci numbers
let a = 1;
let b = 1;

// Print the first two Fibonacci numbers
console.log(a);
console.log(b);

// Generate and print Fibonacci numbers up to N
while (true) {
    const next = a + b;
    if (next > N) {
        break;
    }
    console.log(next);
    a = b;
    b = next;
}