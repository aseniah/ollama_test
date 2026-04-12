import { readdirSync } from 'fs';
import { join } from 'path';

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
    const args = process.argv.slice(2);
    if (args.length !== 1) {
        console.error('Usage: tsx script.ts <N>');
        return;
    }

    const N = parseInt(args[0], 10);

    for (let num = 2; num <= N; num++) {
        if (isPrime(num)) {
            console.log(num);
        }
    }
}

main().catch((error) => {
    console.error('Error:', error);
});