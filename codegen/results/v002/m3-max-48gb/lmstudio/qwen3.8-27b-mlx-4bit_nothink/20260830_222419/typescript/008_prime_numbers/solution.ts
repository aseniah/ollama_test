#!/usr/bin/env tsx

function isPrime(n: number): boolean {
    if (n < 2) return false;
    if (n === 2) return true;
    if (n % 2 === 0) return false;
    for (let i = 3; i * i <= n; i += 2) {
        if (n % i === 0) return false;
    }
    return true;
}

function main(): void {
    const arg = process.argv[2];
    if (arg === undefined) {
        return;
    }

    const n = Number.parseInt(arg, 10);

    if (Number.isNaN(n) || n < 2) {
        return;
    }

    for (let i = 2; i <= n; i++) {
        if (isPrime(i)) {
            console.log(i);
        }
    }
}

main();