import { parseArgs } from 'util';

function isPrime(n: number): boolean {
    if (n <= 1) return false;
    for (let i = 2; i * i <= n; i++) {
        if (n % i === 0) return false;
    }
    return true;
}

if (process.argv.length < 2) {
    process.stderr.write('Usage: node script.js <number>\n');
    process.exit(1);
}

const N = parseInt(process.argv[2], 10);

if (isNaN(N) || N < 2) {
    // If N < 2, there are no primes up to N.
    // Nothing to print.
    process.exit(0);
}

// Use a simple sieve for efficiency if N is large, though basic primality test is fine for moderate N.
const MAX = Math.max(Math.floor(N), 0);
if (MAX > 1e7) {
    console.error('N is too large for this implementation.');
    process.exit(1);
}

const primes: number[] = [];
// Sieve of Eratosthenes up to MAX
for (let i = 2; i <= MAX; i++) {
    if (!primes.includes(i)) {
        // i is prime, add it and mark multiples
        const start = Math.max(2 * i, i * i); // start from 4
        for (let j = start; j <= MAX; j += i) {
            primes.push(j);
        }
    }
}

// primes array now contains all non-primes that were not added to the sieve? 
// Wait, my logic above is flawed. Let's rewrite correctly.