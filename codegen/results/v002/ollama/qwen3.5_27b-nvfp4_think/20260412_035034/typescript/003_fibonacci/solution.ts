const input: string | undefined = process.argv[2];
const N: number = parseInt(input || "", 10);

// If N is not a valid number or less than 1, print nothing and exit
if (isNaN(N) || N < 1) {
    process.exit(0);
}

let a: number = 0;
let b: number = 1;

// Generate Fibonacci sequence starting from 1, 1, 2...
while (b <= N) {
    console.log(b);
    [a, b] = [b, a + b];
}