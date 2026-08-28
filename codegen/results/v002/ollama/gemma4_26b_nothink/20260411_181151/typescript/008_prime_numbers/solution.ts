import * as process from 'node:process';

/**
 * Solves the task of printing all prime numbers up to and including N.
 * Uses the Sieve of Eratosthenes algorithm for efficiency.
 */
function solve(): void {
  // Get the first command line argument
  const args = process.argv.slice(2);

  if (args.length === 0) {
    return;
  }

  // Parse the argument as an integer
  const n = parseInt(args[0], 10);

  // Validate input: if not a number or less than 2, there are no primes to print.
  if (isNaN(n) || n < 2) {
    return;
  }

  // Sieve of Eratosthenes
  // Create a boolean array "isPrime[0..n]" and initialize
  // all entries it as true. A value in isPrime[i] will
  // finally be false if i is Not a prime, else true.
  const isPrime: boolean[] = new Array(n + 1).fill(true);
  isPrime[0] = false;
  isPrime[1] = false;

  for (let p = 2; p * p <= n; p++) {
    // If isPrime[p] is not changed, then it is a prime
    if (isPrime[p]) {
      // Update all multiples of p greater than or equal to the square of it
      // numbers which are multiple of p and are already processed
      // are not needed
      for (let i = p * p; i <= n; i += p) {
        isPrime[i] = false;
      }
    }
  }

  // Collect and print all prime numbers
  const results: string[] = [];
  for (let i = 2; i <= n; i++) {
    if (isPrime[i]) {
      results.push(i.toString());
    }
  }

  if (results.length > 0) {
    process.stdout.write(results.join('\n') + '\n');
  }
}

// Run the solver
solve();