const N = parseInt(process.argv[2], 10);

if (isNaN(N) || N < 1) {
    // Print nothing
    process.exit(0);
}

if (N === 1) {
    // The Fibonacci sequence starts: 1, 1, 2, 3, 5, ...
    // Largest Fibonacci number that does not exceed 1 is 1.
    // But we need to print all Fibonacci numbers up to and including the largest that doesn't exceed N.
    // The sequence is 1, 1, 2, 3, 5, ...
    // So for N=1, the Fibonacci numbers are 1, 1. Both are <= 1.
    // Wait, the problem says "Start the sequence: 1, 1, 2, 3, 5, ..."
    // And "Print all Fibonacci numbers up to and including the largest Fibonacci number that does not exceed N."
    // So for N=1, we print 1 and 1.
    console.log("1");
    console.log("1");
    process.exit(0);
}

let a = 1;
let b = 1;
console.log(a); // 1
console.log(b); // 1

a = 1;
b = 1;
let c = a + b;
while (c <= N) {
    console.log(c);
    a = b;
    b = c;
    c = a + b;
}