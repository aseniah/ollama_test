// Read N from command line
const N = parseInt(process.argv[2]);

// Check if N is valid
if (N < 1) {
    process.exit(0);
}

// Generate and print Fibonacci sequence up to N
let a = 1;
let b = 1;

while (a <= N) {
    console.log(a);
    let temp = a + b;
    a = b;
    b = temp;
}