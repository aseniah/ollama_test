const input = process.argv[2];
if (!input) {
  process.exit(0);
}

const n = parseInt(input, 10);
if (isNaN(n) || n < 2) {
  process.exit(0);
}

// Sieve of Eratosthenes
const isPrime: boolean[] = new Array(n + 1).fill(true);
isPrime[0] = false;
isPrime[1] = false;

for (let p = 2; p * p <= n; p++) {
  if (isPrime[p]) {
    for (let i = p * p; i <= n; i += p) {
      isPrime[i] = false;
    }
  }
}

// Collect primes
const primes: string[] = [];
for (let p = 2; p <= n; p++) {
  if (isPrime[p]) {
    primes.push(p.toString());
  }
}

// Print results
if (primes.length > 0) {
  process.stdout.write(primes.join('\n') + '\n');
}