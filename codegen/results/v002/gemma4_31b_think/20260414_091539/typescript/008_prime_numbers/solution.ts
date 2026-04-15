/**
 * This program finds and prints all prime numbers up to and including a 
 * given integer N provided via command line arguments.
 * 
 * It uses the Sieve of Eratosthenes algorithm for efficiency.
 */

function main() {
  // Retrieve the argument from the command line
  const arg = process.argv[2];

  if (arg === undefined) {
    return;
  }

  const n = parseInt(arg, 10);

  // Check if the input is a valid number
  if (isNaN(n)) {
    return;
  }

  // Primes are defined as integers greater than 1
  if (n < 2) {
    return;
  }

  // Sieve of Eratosthenes implementation
  // Create a boolean array "isPrime[0..n]" and initialize
  // all entries it as true. A value in isPrime[i] will
  // finally be false if i is Not a prime, else true.
  const isPrime = new Uint8Array(n + 1).fill(1);
  isPrime[0] = 0;
  isPrime[1] = 0;

  const limit = Math.sqrt(n);
  for (let p = 2; p <= limit; p++) {
    // If isPrime[p] is not changed, then it is a prime
    if (isPrime[p] === 1) {
      // Update all multiples of p starting from p*p
      // because smaller multiples have already been marked
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
    // To prevent memory issues with extremely large N when building a string,
    // we can print in chunks or use process.stdout.write
    if (output.length > 10000) {
      process.stdout.write(output);
      output = "";
    }
  }
  process.stdout.write(output);
}

main();