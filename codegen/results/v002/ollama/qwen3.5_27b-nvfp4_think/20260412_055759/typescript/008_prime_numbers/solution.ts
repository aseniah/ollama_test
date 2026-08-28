// Get the command line argument
const arg = process.argv[2];

// Validate input
if (arg === undefined || isNaN(parseInt(arg, 10))) {
  // Invalid or missing argument, exit silently
  process.exit(0);
}

const N = parseInt(arg, 10);

// If N is less than 2, there are no primes up to N
if (N < 2) {
  process.exit(0);
}

// Use Sieve of Eratosthenes to find all primes up to N
const sieve = new Array<boolean>(N + 1).fill(true);
sieve[0] = false;
sieve[1] = false;

for (let i = 2; i * i <= N; i++) {
  if (sieve[i]) {
    for (let j = i * i; j <= N; j += i) {
      sieve[j] = false;
    }
  }
}

// Print all primes, one per line
for (let i = 2; i <= N; i++) {
  if (sieve[i]) {
    console.log(i);
  }
}