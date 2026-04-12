/**
 * Sieve of Eratosthenes to find all prime numbers up to N.
 * 
 * @param n The upper limit (inclusive)
 * @returns An array of prime numbers up to n
 */
function getPrimesUpTo(n: number): number[] {
  if (n < 2) {
    return [];
  }

  // Initialize a boolean array to track prime status.
  // isPrime[i] will be false if i is prime, true if composite.
  const isComposite: boolean[] = new Array(n + 1).fill(false);
  const primes: number[] = [];

  for (let i = 2; i <= n; i++) {
    if (!isComposite[i]) {
      primes.push(i);

      // Mark multiples of i starting from i*i
      if (BigInt(i) * BigInt(i) <= BigInt(n)) {
        for (let j = i * i; j <= n; j += i) {
          isComposite[j] = true;
        }
      }
    }
  }

  return primes;
}

function main() {
  // Read command line argument
  if (process.argv.length < 3) {
    // No argument provided, print nothing
    return;
  }

  const input = process.argv[2];
  const n = parseInt(input, 10);

  if (isNaN(n)) {
    // If input is not a valid integer, print nothing or handle error
    // The problem implies N is an integer, so we assume valid integer input or print nothing
    return;
  }

  const primes = getPrimesUpTo(n);

  if (primes.length > 0) {
    console.log(primes.join('\n'));
  }
}

main();