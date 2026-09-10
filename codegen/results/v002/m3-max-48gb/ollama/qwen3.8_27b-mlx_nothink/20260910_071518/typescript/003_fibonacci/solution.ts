const N = parseInt(process.argv[2], 10);

if (isNaN(N) || N < 1) {
    // print nothing
    process.exit(0);
}

// Start with 1, 1
console.log(1);

if (N === 1) {
    process.exit(0);
}

console.log(1);

let a = 1;
let b = 1;

while (true) {
    const next = a + b;
    if (next > N) {
        break;
    }
    console.log(next);
    a = b;
    b = next;
}