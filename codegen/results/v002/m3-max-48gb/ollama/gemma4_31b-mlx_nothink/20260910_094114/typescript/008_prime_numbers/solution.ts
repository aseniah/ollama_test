/**
 * This program takes a single integer N from the command line
 * and prints all prime numbers up to and including N.
 */

function isPrime(num: number): boolean {
  if (num <= 1) return false;
  if (num <= 3) return true;
  if (num % 2 === 0 || num % 3 === 0) return false;

  // Use 6k +/- 1 optimization
  for (let i = 5; i * i <= num; i += 6) {
    if (num % i === 0 || num % (i + 2) === 0) {
      return false;
    }
  }
  return true;
}

function main() {
  // Get the argument from the command line
  const arg = process.argv[2];

  if (arg === undefined) {
    return;
  }

  // Parse the argument as an integer
  const n = parseInt(arg, 10);

  // Ensure N is a valid number
  if (isNaN(n)) {
    return;
  }

  // Iterate from 2 up to N and print primes
  for (let i = 2; i <= n; i++) {
    if (isPrime(i)) {
      process.stdout.write(i.toString() + '\n');
    }
  }
}

main();