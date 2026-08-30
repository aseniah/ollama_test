// Read the integer N from command line
const nStr = process.argv[2];
const N = parseInt(nStr, 10);

if (isNaN(N) || N < 1) {
    // Print nothing
    process.exit(0);
}

// Generate Fibonacci numbers up to N
// Sequence: 1, 1, 2, 3, 5, ...
let a = 1;
let b = 1;

// Print the first two 1s if N >= 1
// We need to print all Fibonacci numbers <= N

// Use a loop to generate Fibonacci numbers
const fibs: number[] = [];
let f1 = 1;
let f2 = 1;
fibs.push(f1);
if (N >= f2) {
    fibs.push(f2);
}

// Continue generating
while (true) {
    const next = f1 + f2;
    if (next > N) {
        break;
    }
    fibs.push(next);
    f1 = f2;
    f2 = next;
}

// Print each Fibonacci number on its own line
for (const f of fibs) {
    console.log(f);
}