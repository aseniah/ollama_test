import * as fs from 'fs';

function isPrime(num: number): boolean {
    if (num < 2) return false;
    if (num === 2) return true;
    if (num % 2 === 0) return false;
    
    const limit = Math.sqrt(num);
    for (let i = 3; i <= limit; i += 2) {
        if (num % i === 0) return false;
    }
    return true;
}

function main(): void {
    const arg = process.argv[2];
    
    if (arg === undefined) {
        return;
    }

    const n = parseInt(arg, 10);

    if (isNaN(n)) {
        return;
    }

    for (let i = 2; i <= n; i++) {
        if (isPrime(i)) {
            console.log(i);
        }
    }
}

main();