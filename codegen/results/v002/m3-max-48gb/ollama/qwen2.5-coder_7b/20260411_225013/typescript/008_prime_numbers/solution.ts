// prime-numbers.ts

function isPrime(num: number): boolean {
    if (num <= 1) return false;
    if (num <= 3) return true;

    if (num % 2 === 0 || num % 3 === 0) return false;

    for (let i = 5; i * i <= num; i += 6) {
        if (num % i === 0 || num % (i + 2) === 0) return false;
    }

    return true;
}

function printPrimesUpToN(n: number): void {
    for (let i = 2; i <= n; i++) {
        if (isPrime(i)) {
            console.log(i);
        }
    }
}

const args = process.argv.slice(2);

if (args.length !== 1) {
    console.error("Usage: tsx prime-numbers.ts <integer>");
    process.exit(1);
}

const n = parseInt(args[0]);

if (isNaN(n) || !Number.isInteger(n) || n < 0) {
    console.error("Please provide a valid non-negative integer.");
    process.exit(1);
}

printPrimesUpToN(n);