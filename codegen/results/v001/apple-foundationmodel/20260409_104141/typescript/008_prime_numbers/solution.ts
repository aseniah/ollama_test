import fs from 'fs';
import path from 'path';

function isPrime(n: number): boolean {
    if (n <= 1) return false;
    if (n <= 3) return true;
    if (n % 2 === 0 || n % 3 === 0) return false;
    let i = 5;
    while (i * i <= n) {
        if (n % i === 0 || n % (i + 2) === 0) return false;
        i += 6;
    }
    return true;
}

async function main(): Promise<void> {
    const args = process.argv;
    if (args.length < 2) {
        console.error('Usage: n [number]');
        return;
    }

    const N = parseInt(args[2], 10);

    if (isNaN(N) || N < 2) {
        console.error('Please provide a positive integer greater than or equal to 2.');
        return;
    }

    console.log('Prime numbers up to and including', N, ':');
    for (let i = 2; i <= N; i++) {
        if (isPrime(i)) {
            console.log(i);
        }
    }
}

main().catch(console.error);