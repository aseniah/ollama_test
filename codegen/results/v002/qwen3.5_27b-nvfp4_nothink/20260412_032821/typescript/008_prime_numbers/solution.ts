/**
 * This program accepts a single integer N from the command line
 * and prints all prime numbers up to and including N.
 * 
 * Usage: tsx index.ts <N>
 */

function isPrime(n: number): boolean {
  if (n <= 1) return false;
  if (n === 2) return true;
  if (n % 2 === 0) return false;
  
  // Check odd divisors up to the square root of n
  const limit = Math.sqrt(n);
  for (let i = 3; i <= limit; i += 2) {
    if (n % i === 0) return false;
  }
  
  return true;
}

function main(): void {
  // Get the command line argument
  const input = process.argv[2];
  
  // Validate that an argument was provided
  if (!input) {
    console.error("Usage: tsx index.ts <N>");
    process.exit(1);
  }
  
  const N = Number.parseInt(input, 10);
  
  // Validate that the input is a valid integer
  if (Number.isNaN(N)) {
    console.error("Error: Argument must be an integer.");
    process.exit(1);
  }
  
  // Iterate from 2 up to N and print primes
  for (let i = 2; i <= N; i++) {
    if (isPrime(i)) {
      console.log(i);
    }
  }
}

main();