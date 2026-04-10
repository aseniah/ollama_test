// Parse the command-line argument N
const N = parseInt(process.argv[2] || "0", 10);

// Exit early if N is less than 1
if (N < 1) {
    process.exit(0);
}

// Start generating Fibonacci sequence starting with 1, 1
let a = 1;
let b = 1;

// Print the first two Fibonacci numbers
console.log(a);
console.log(b);

// Generate remaining Fibonacci numbers
while (true) {
    const next = a + b;
    if (next > N) {
        break;
    }
    console.log(next);
    a = b;
    b = next;
}