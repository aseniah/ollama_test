const args = process.argv.slice(2);
if (args.length < 1) {
    process.exit(0);
}

const N = parseInt(args[0], 10);
if (isNaN(N) || N < 1) {
    process.exit(0);
}

let a = 1;
let b = 1;

// Print the first Fibonacci number if it doesn't exceed N
if (a <= N) {
    console.log(a);
}

if (b <= N) {
    console.log(b);
}

let next = a + b;
while (next <= N) {
    console.log(next);
    a = b;
    b = next;
    next = a + b;
}