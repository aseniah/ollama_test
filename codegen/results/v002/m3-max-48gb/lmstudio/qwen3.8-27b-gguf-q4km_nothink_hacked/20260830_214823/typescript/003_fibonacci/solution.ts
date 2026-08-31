// Read the integer N from command line arguments
const nArg = process.argv[2];

if (nArg === undefined) {
    process.exit(0);
}

let N: number;
try {
    N = parseInt(nArg, 10);
    if (isNaN(N)) {
        process.exit(0);
    }
} catch {
    process.exit(0);
}

if (N < 1) {
    process.exit(0);
}

const results: number[] = [];
let a = 1;
let b = 1;

// Start with 1
if (a <= N) {
    results.push(a);
    a = 1;
    b = 2;
} else {
    process.exit(0);
}

// Generate Fibonacci numbers
while (true) {
    if (b > N) break;
    results.push(b);
    const next = a + b;
    a = b;
    b = next;
}

// Print each result on a new line
for (const r of results) {
    console.log(r);
}