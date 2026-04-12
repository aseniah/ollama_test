/**
 * Sieve of Eratosthenes implementation to find all prime numbers up to N.
 */
function sieveOfEratosthenes(n: number): number[] {
  if (n < 2) {
    return [];
  }

  // Create a boolean array "isPrime[0..n]" and initialize
  // all entries as true. A value in isPrime[i] will be false
  // if i is not a prime, else true.
  const isPrime: boolean[] = new Array(n + 1).fill(true);

  isPrime[0] = false;
  isPrime[1] = false;

  // Start from 2 (the first prime)
  for (let p = 2; p * p <= n; p++) {
    // If isPrime[p] is not changed, then it is a prime
    if (isPrime[p]) {
      // Update all multiples of p starting from p*p
      for (let i = p * p; i <= n; i += p) {
        isPrime[i] = false;
      }
    }
  }

  // Collect all prime numbers
  const primes: number[] = [];
  for (let p = 2; p <= n; p++) {
    if (isPrime[p]) {
      primes.push(p);
    }
  }

  return primes;
}

/**
 * Main execution function.
 */
function main(): void {
  // Get command line arguments
  const args = process.argv;
  
  // Skip the first two arguments (node/tsx and script name)
  if (args.length < 3) {
    // If no argument is provided, we print nothing as per instructions
    // (or we could throw an error, but the spec implies handling valid input or doing nothing)
    return;
  }

  const input = args[2];
  const n = parseInt(input, 10);

  if (isNaN(n)) {
    return;
  }

  const primes = sieveOfEratosthenes(n);

  // Print each prime on a new line
  primes.forEach((prime) => {
    console.log(prime);
  });
}

main();