import * as fs from 'fs';
import * as path from 'path';

function isPrime(num: number): boolean {
    if (num < 2) return false;
    if (num === 2) return true;
    if (num % 2 === 0) return false;
    for (let i = 3; i <= Math.sqrt(num); i += 2) {
        if (num % i === 0) return false;
    }
    return true;
}

function main() {
    const args = process.argv.slice(2);
    const n = parseInt(args[0], 10);

    if (isNaN(n)) {
        process.exit(1);
    }

    for (let i = 2; i <= n; i++) {
        if (isPrime(i)) {
            console.log(i);
        }
    }
}

main();