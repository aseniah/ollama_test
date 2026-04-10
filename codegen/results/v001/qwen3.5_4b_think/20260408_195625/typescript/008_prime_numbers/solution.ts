// Read command line argument N
const N = parseInt(process.argv[2]);

// Validate input
if (isNaN(N) || N < 0) {
  process.exit(0);
}

// Handle the case where N is less than the first prime number (2)
if (N < 2) {
  process.exit(0);
}

// Sieve of Eratosthenes to find primes up to N
const isPrime = new Array(N + 1).fill(true);

// 0 and 1 are not prime
isPrime[0] = false;
isPrime[1] = false;

// Mark multiples as non-prime
for (let i = 2; i * i <= N; i++) {
  if (isPrime[i]) {
    for (let j = i * i; j <= N; j += i) {
      isPrime[j] = false;
    }
  }
}

// Collect prime numbers
const primes: number[] = [];
for (let i = 2; i <= N; i++) {
  if (isPrime[i]) {
    primes.push(i);
  }
}

// Print all prime numbers, one per line
for (const prime of primes) {
  console.log(prime);
}