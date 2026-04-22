import * as process from 'process';

function isPrime(n: number): boolean {
    if (n < 2) return false;
    if (n === 2) return true;
    if (n % 2 === 0) return false;
    for (let i = 3; i <= Math.sqrt(n); i += 2) {
        if (n % i === 0) return false;
    }
    return true;
}

const input = process.argv[2];
if (input === undefined) {
    process.exit(0);
}

const n = parseInt(input, 10);
if (isNaN(n)) {
    process.exit(0);
}

for (let i = 2; i <= n; i++) {
    if (isPrime(i)) {
        console.log(i);
    }
}