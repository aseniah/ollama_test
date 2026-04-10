import { readline } from 'readline';
import { path, join } from 'path';
import { resolve } from 'path';
import { exec } from 'child_process';

// Function to read the Fibonacci sequence up to N
function printFibonacciSequence(n: number): void {
  const fibNumbers = [1, 1];

  while (true) {
    const nextFib = fibNumbers[fibNumbers.length - 1] + fibNumbers[fibNumbers.length - 2];
    if (nextFib > n) {
      break;
    }
    fibNumbers.push(nextFib);
  }

  for (const fib of fibNumbers) {
    console.log(fib);
  }
}

// Read command line argument N
const args = process.argv.slice(2);
if (args.length < 1) {
  console.log('Please provide a valid integer argument N.');
} else {
  try {
    const N = Number(args[0]);
    if (N < 1) {
      console.log('N must be at least 1.');
    } else {
      printFibonacciSequence(N);
    }
  } catch (e) {
    console.log('Invalid input. Please provide a valid integer.');
  }
}