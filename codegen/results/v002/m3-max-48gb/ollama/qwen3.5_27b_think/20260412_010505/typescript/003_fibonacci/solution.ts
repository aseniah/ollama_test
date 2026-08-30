const N = parseInt(process.argv[2], 10);

if (isNaN(N) || N < 1) {
    // Print nothing if N is not a valid number or less than 1
    process.exit(0);
}

let a = 1;
let b = 1;

// Print the first Fibonacci number
console.log(a);

// Continue printing Fibonacci numbers while they do not exceed N
while (b <= N) {
    console.log(b);
    const next = a + b;
    a = b;
    b = next;
}