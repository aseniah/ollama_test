import { readline } from 'readline';
import fs from 'fs';

async function printPrimes(n: number): Promise<void> {
    const primes = [];

    for (let i = 2; i <= n; i++) {
        if (isPrime(i)) {
            primes.push(i);
        }
    }

    for (const prime of primes) {
        console.log(prime);
    }
}

function isPrime(num: number): boolean {
    if (num <= 1) return false;
    if (num <= 3) return true;
    if (num % 2 === 0 || num % 3 === 0) return false;

    let i = 5;
    while (i * i <= num) {
        if (num % i === 0 || num % (i + 2) === 0) return false;
        i += 6;
    }
    return true;
}

async function main(): Promise<void> {
    if (process.argv.length < 2) {
        console.error('Usage: node prime.tsx <number>');
        process.exit(1);
    }

    const n = parseInt(process.argv[2], 10);
    await printPrimes(n);
}

main().catch(console.error);