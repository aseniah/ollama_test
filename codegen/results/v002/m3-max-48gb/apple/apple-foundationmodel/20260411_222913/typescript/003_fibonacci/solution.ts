import { readline } from 'readline';
import fs from 'fs';
import path from 'path';
import { argv } from 'process';

// Function to generate Fibonacci numbers up to a given limit
const generateFibonacciUpTo = (limit: number): number[] => {
  const fibonacciNumbers = [1, 1];
  let nextFibonacci = 2; // The next Fibonacci number after 1, 1

  while (nextFibonacci <= limit) {
    fibonacciNumbers.push(nextFibonacci);
    nextFibonacci = fibonacciNumbers[fibonacciNumbers.length - 1] + fibonacciNumbers[fibonacciNumbers.length - 2];
  }

  return fibonacciNumbers;
};

// Main function
async function main() {
  try {
    // Read command-line argument
    const N = parseInt(argv[2]);
    if (isNaN(N) || N < 1) {
      console.log(''); // Print nothing if input is invalid or less than 1
      return;
    }

    // Generate Fibonacci numbers up to N
    const fibonacciNumbers = generateFibonacciUpTo(N);

    // Print Fibonacci numbers
    console.log(...fibonacciNumbers);
  } catch (error) {
    console.error('Error:', error);
  }
}

// Run the main function
main();