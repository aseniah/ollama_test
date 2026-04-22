import { readFileSync } from 'fs';
import { dirname } from 'path';

// Get N from command line arguments
const args = process.argv.slice(2);
const N = parseInt(args[0], 10);

if (isNaN(N)) {
    // No valid input, print nothing
    process.exit(0);
}

function isPrime(num: number): boolean {
    if (num < 2) {
        return false;
    }
    if (num === 2) {
        return true;
    }
    if (num % 2 === 0) {
        return false;
    }
    for (let i = 3; i <= Math.sqrt(num); i += 2) {
        if (num % i === 0) {
            return false;
        }
    }
    return true;
}

if (N >= 2) {
    for (let i = 2; i <= N; i++) {
        if (isPrime(i)) {
            console.log(i);
        }
    }
}