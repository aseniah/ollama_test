/**
 * This program accepts a single integer argument N and prints all prime numbers
 * up to and including N, one per line.
 */

function main() {
  // Get the command line argument N
  const arg = process.argv[2];

  if (arg === undefined) {
    return;
  }

  const n = parseInt(arg, 10);

  // If N is not a number or is less than 2, there are no primes to print.
  if (isNaN(n) || n < 2) {
    return;
  }

  // Use the Sieve of Eratosthenes to find all primes up to N efficiently.
  // isPrime[i] will be true if i is a prime number.
  const isPrime = new Uint8Array(n + 1).fill(1);
  isPrime[0] = 0; // 0 is not prime
  isPrime[1] = 0; // 1 is not prime

  const sqrtN = Math.floor(Math.sqrt(n));
  for (let p = 2; p <= sqrtN; p++) {
    if (isPrime[p] === 1) {
      // Mark all multiples of p starting from p*p as not prime
      for (let i = p * p; i <= n; i += p) {
        isPrime[i] = 0;
      }
    }
  }

  // Print all numbers that remained marked as prime
  let output = "";
  for (let i = 2; i <= n; i++) {
    if (isPrime[i] === 1) {
      output += i + "\n";
    }
    
    // Periodically flush the output string to avoid memory issues with very large N
    if (output.length > 10000) {
      process.stdout.write(output);
      output = "";
    }
  }
  process.stdout.write(output);
}

main();