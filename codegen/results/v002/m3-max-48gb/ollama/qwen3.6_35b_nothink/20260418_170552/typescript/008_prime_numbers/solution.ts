import * as process from 'process';

function isPrime(n: number): boolean {
    if (n < 2) {
        return false;
    }
    if (n === 2) {
        return true;
    }
    if (n % 2 === 0) {
        return false;
    }
    const sqrt = Math.sqrt(n);
    for (let i = 3; i <= sqrt; i += 2) {
        if (n % i === 0) {
            return false;
        }
    }
    return true;
}

const N = parseInt(process.argv[2], 10);

if (isNaN(N) || N < 2) {
    // No primes to print
    process.exit(0);
}

for (let i = 2; i <= N; i++) {
    if (isPrime(i)) {
        console.log(i);
    }
}