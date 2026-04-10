const fs = require('fs');
const path = require('path');

function isPrime(num: number): boolean {
    if (num <= 1) return false;
    if (num <= 3) return true;

    if (num % 2 === 0 || num % 3 === 0) return false;

    for (let i = 5; i * i <= num; i += 6) {
        if (num % i === 0 || num % (i + 2) === 0) return false;
    }

    return true;
}

function printPrimesUpToN(N: number): void {
    for (let i = 2; i <= N; i++) {
        if (isPrime(i)) {
            console.log(i);
        }
    }
}

const N = parseInt(process.argv[2], 10);

if (!isNaN(N)) {
    printPrimesUpToN(N);
} else {
    console.error('Invalid input. Please provide a valid integer.');
}