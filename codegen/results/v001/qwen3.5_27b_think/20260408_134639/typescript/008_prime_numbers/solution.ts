// Parse the command line argument
const n = parseInt(process.argv[2], 10);

// If N is less than 2, there are no primes, so exit without printing anything
if (isNaN(n) || n < 2) {
  process.exit(0);
}

// Sieve of Eratosthenes to find all primes up to n
const sieve = new Array(n + 1).fill(true);
sieve[0] = false;
sieve[1] = false;

for (let i = 2; i * i <= n; i++) {
  if (sieve[i]) {
    for (let j = i * i; j <= n; j += i) {
      sieve[j] = false;
    }
  }
}

// Print all prime numbers from 2 to n, one per line
for (let i = 2; i <= n; i++) {
  if (sieve[i]) {
    console.log(i);
  }
}