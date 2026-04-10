import fs from 'fs';
import path from 'path';
import { parseArgv } from 'util';
import { exec } from 'child_process';

// Function to generate Fibonacci numbers up to a specified limit
function generateFibonacci(limit: number): number[] {
  const fibonacciNumbers: number[] = [1, 1];

  while (true) {
    const nextFibonacci = fibonacciNumbers[fibonacciNumbers.length - 1] + fibonacciNumbers[fibonacciNumbers.length - 2];
    if (nextFibonacci > limit) {
      break;
    }
    fibonacciNumbers.push(nextFibonacci);
  }

  return fibonacciNumbers;
}

// Main function
async function main(): Promise<void> {
  const argv = parseArgv(process.argv, {
    numeric: true,
    default: 1,
  });

  const n = argv[0];

  if (n < 1) {
    console.log('No Fibonacci numbers printed.');
    return;
  }

  const fibonacciNumbers = generateFibonacci(n);

  console.log(fibonacciNumbers.join('\n'));
}

main().catch((error) => {
  console.error('Error:', error);
});