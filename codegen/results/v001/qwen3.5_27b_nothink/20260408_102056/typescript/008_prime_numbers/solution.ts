// Accept N from command line arguments
const input = process.argv[2];

if (!input) {
  process.exit(0);
}

const n = parseInt(input, 10);

if (isNaN(n) || n < 2) {
  // No primes exist below 2
  process.exit(0);
}

const isPrime = (num: number): boolean => {
  if (num < 2) return false;
  if (num === 2) return true;
  if (num % 2 === 0) return false;

  const limit = Math.sqrt(num);
  for (let i = 3; i <= limit; i += 2) {
    if (num % i === 0) return false;
  }
  return true;
};

for (let i = 2; i <= n; i++) {
  if (isPrime(i)) {
    console.log(i);
  }
}