/**
 * Checks if a number is prime.
 * Optimized to check divisibility only up to the square root of n.
 */
function isPrime(num: number): boolean {
  if (num <= 1) return false;
  if (num === 2) return true;
  if (num % 2 === 0) return false;

  const limit = Math.floor(Math.sqrt(num));
  for (let i = 3; i <= limit; i += 2) {
    if (num % i === 0) return false;
  }
  return true;
}

/**
 * Main execution function
 */
function main() {
  // Read the command line argument N
  // process.argv[0] is the node executable, process.argv[1] is the script name
  // process.argv[2] is the first user argument
  const arg = process.argv[2];

  if (!arg) {
    return; // No argument provided, exit gracefully
  }

  const n = parseInt(arg, 10);

  if (isNaN(n)) {
    return; // Invalid integer provided, exit gracefully
  }

  // Iterate from 2 up to and including N
  for (let i = 2; i <= n; i++) {
    if (isPrime(i)) {
      console.log(i);
    }
  }
}

main();