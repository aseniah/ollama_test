// fibonacci.ts

function generateFibonacciNumbers(n: number): void {
    if (n < 1) return;

    let a = 1, b = 1;
    while (a <= n) {
        console.log(a);
        [a, b] = [b, a + b];
    }
}

if (process.argv.length < 3) {
    console.error("Usage: tsx fibonacci.ts <N>");
    process.exit(1);
}

const N = parseInt(process.argv[2], 10);

if (isNaN(N)) {
    console.error("Invalid argument. Please provide a valid integer.");
    process.exit(1);
}

generateFibonacciNumbers(N);