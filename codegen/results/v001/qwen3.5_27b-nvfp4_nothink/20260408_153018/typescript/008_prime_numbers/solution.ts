import { argv } from 'process';

/**
 * Checks if a given number n is prime.
 */
function isPrime(n: number): boolean {
  if (n < 2) return false;
  if (n === 2) return true;
  if (n % 2 === 0) return false;

  const limit = Math.sqrt(n);
  for (let i = 3; i <= limit; i += 2) {
    if (n % i === 0) return false;
  }
  
  return true;
}

/**
 * Generates all prime numbers up to N.
 */
function printPrimesUpToN(n: number): void {
  // Iterate from 2 to n
  for (let i = 2; i <= n; i++) {
    if (isPrime(i)) {
      console.log(i);
    }
  }
}

// Entry point
const main = () => {
  // Expecting a single argument at index 2
  if (argv.length !== 3) {
    return; 
  }

  const input = argv[2];
  const n = parseInt(input, 10);

  // Validate if parsing was successful
  if (isNaN(n)) {
    return;
  }

  printPrimesUpToN(n);
};

main();