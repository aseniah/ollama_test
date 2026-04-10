// Check if the required argument is provided
if (process.argv.length !== 3) {
    console.error("Usage: tsx fibonacci.ts <integer>");
    process.exit(1);
}

// Parse the input argument as an integer
const N = parseInt(process.argv[2], 10);

// Function to generate Fibonacci numbers up to N
function* fibonacciUpToN(n: number): Generator<number> {
    let a = 1, b = 1;
    while (a <= n) {
        yield a;
        [a, b] = [b, a + b];
    }
}

// Generate and print Fibonacci numbers up to N
for (const num of fibonacciUpToN(N)) {
    console.log(num);
}