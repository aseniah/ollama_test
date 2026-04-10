/**
 * This script accepts a single integer N from the command line and
 * prints all prime numbers up to and including N, one per line.
 * It uses the Sieve of Eratosthenes algorithm for efficient computation.
 */

function main(): void {
  // Retrieve arguments from command line
  const args = process.argv.slice(2);

  if (args.length === 0) {
    return;
  }

  // Parse the first argument as an integer
  const n = parseInt(args[0], 10);

  // If N is not a valid number or is less than 2, there are no primes to print.
  if (isNaN(n) || n < 2) {
    return;
  }

  /**
   * Sieve of Eratosthenes
   * 
   * We use a Uint8Array instead of a standard Boolean array to be more
   * memory-efficient, especially for large values of N.
   * 1 represents prime, 0 represents not prime.
   */
  const isPrime = new Uint8Array(n + 1);
  // Initialize all indices as prime (1)
  isPrime.fill(1);

  // 0 and 1 are not prime numbers
  isPrime[0] = 0;
  isPrime[1] = 0;

  const limit = Math.sqrt(n);
  for (let p = 2; p <= limit; p++) {
    if (isPrime[p] === 1) {
      // Start marking multiples from p * p, as smaller multiples 
      // would have already been marked by smaller prime factors.
      for (let i = p * p; i <= n; i += p) {
        isPrime[i] = 0;
      }
    }
  }

  /**
   * Outputting results.
   * 
   * To optimize performance for large N, we buffer the output strings
   * and write to stdout in chunks. This is significantly faster than 
   * calling console.log or process.stdout.write for every single number.
   */
  let buffer = "";
  const BUFFER_THRESHOLD = 16384; // 16KB buffer

  for (let i = 2; i <= n; i++) {
    if (isPrime[i] === 1) {
      buffer += i + "\n";

      // If buffer exceeds threshold, flush to stdout
      if (buffer.length > BUFFER_THRESHOLD) {
        process.stdout.write(buffer);
        buffer = "";
      }
    }
  }

  // Final flush for any remaining content in the buffer
  if (buffer.length > 0) {
    process.stdout.write(buffer);
  }
}

// Execute the main function
main();