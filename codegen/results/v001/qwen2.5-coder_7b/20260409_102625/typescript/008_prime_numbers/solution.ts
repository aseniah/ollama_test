const fs = require('fs');
const path = require('path');

// Helper function to check if a number is prime
function isPrime(num: number): boolean {
    if (num <= 1) return false;
    if (num <= 3) return true;

    if (num % 2 === 0 || num % 3 === 0) return false;

    for (let i = 5; i * i <= num; i += 6) {
        if (num % i === 0 || num % (i + 2) === 0) return false;
    }

    return true;
}

// Main function to print all prime numbers up to and including N
function printPrimesUpToN(N: number): void {
    for (let i = 2; i <= N; i++) {
        if (isPrime(i)) {
            console.log(i);
        }
    }
}

// Read the command-line argument
const N = parseInt(process.argv[2], 10);

if (!isNaN(N) && N >= 0) {
    printPrimesUpToN(N);
} else {
    console.error('Please provide a non-negative integer as an argument.');
}