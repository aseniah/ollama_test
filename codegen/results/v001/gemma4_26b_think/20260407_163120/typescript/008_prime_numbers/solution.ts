/**
 * This script finds all prime numbers up to a given integer N 
 * provided via command-line arguments using the Sieve of Eratosthenes algorithm.
 */

function main() {
  // Get the first argument from the command line
  const input = process.argv[2];

  // If no argument is provided, exit silently
  if (input === undefined) {
    return;
  }

  const N = parseInt(input, 10);

  // If N is not a number or is less than 2, there are no primes to print
  if (isNaN(N) || N < 2) {
    return;
  }

  /**
   * Sieve of Eratosthenes implementation.
   * We use a Uint8Array for memory efficiency. 
   * 1 represents prime, 0 represents not prime.
   */
  const isPrime = new Uint8Array(N + 1).fill(1);
  isPrime[0] = 0;
  isPrime[1] = 0;

  const limit = Math.sqrt(N);
  for (let p = 2; p <= limit; p++) {
    if (isPrime[p] === 1) {
      // Mark multiples of p starting from p*p as not prime
      for (let multiple = p * p; multiple <= N; multiple += p) {
        isPrime[multiple] = 0;
      }
    }
  }

  /**
   * Output the primes.
   * We use a buffer string to reduce the number of syscalls to stdout,
   * which is significantly faster than calling console.log for every number.
   */
  let outputBuffer = "";
  const BUFFER_THRESHOLD = 16384; // 16KB buffer

  for (let i = 2; i <= N; i++) {
    if (isPrime[i] === 1) {
      outputBuffer += i + "\n";

      // Periodically flush the buffer to stdout
      if (outputBuffer.length > BUFFER_THRESHOLD) {
        process.stdout.write(outputBuffer);
        outputBuffer = "";
      }
    }
  }

  // Final flush of any remaining content in the buffer
  if (outputBuffer.length > 0) {
    process.stdout.write(outputBuffer);
  }
}

// Execute the main function
main();