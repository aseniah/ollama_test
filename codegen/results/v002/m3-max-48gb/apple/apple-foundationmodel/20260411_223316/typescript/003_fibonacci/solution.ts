import { readline } from 'readline';
import fs from 'fs';
import path from 'path';

function printFibonacciNumbers(upTo n: number): void {
    const fibonacciNumbers = [1, 1]; // Start with the first two Fibonacci numbers
    while (fibonacciNumbers[fibonacciNumbers.length - 1] <= n) {
        fibonacciNumbers.push(fibonacciNumbers[fibonacciNumbers.length - 1] + fibonacciNumbers[fibonacciNumbers.length - 2]);
    }
    for (const number of fibonacciNumbers) {
        console.log(number);
    }
}

async function main(): Promise<void> {
    if (process.argv.length < 2) {
        console.log('Usage: tsx fibonacci <N>');
        return;
    }

    try {
        const n = Number(process.argv[2]);
        if (n < 1) {
            console.log('N must be a positive integer.');
            return;
        }

        printFibonacciNumbers(n);
    } catch (error) {
        console.error('Invalid input. Please ensure you provide a valid integer.');
    }
}

main().catch((error) => {
    console.error('Error:', error);
});