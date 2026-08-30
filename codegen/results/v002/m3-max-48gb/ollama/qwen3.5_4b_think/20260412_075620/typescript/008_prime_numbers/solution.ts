// Get the argument N from command line
const N = parseInt(process.argv[2]);

if (isNaN(N)) {
  process.exit(1);
}

// Sieve of Eratosthenes to find primes up to N
const isPrimeArray = new Array(N + 1).fill(true);
isPrimeArray[0] = false;
isPrimeArray[1] = false;

// Mark non-primes
for (let i = 2; i * i <= N; i++) {
  if (isPrimeArray[i]) {
    // Mark all multiples of i as composite
    for (let j = i * i; j <= N; j += i) {
      isPrimeArray[j] = false;
    }
  }
}

// Print all prime numbers up to and including N
for (let i = 2; i <= N; i++) {
  if (isPrimeArray[i]) {
    console.log(i);
  }
}