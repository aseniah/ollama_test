// Read N from command line
const N = parseInt(process.argv[2]);

// Validate input
if (isNaN(N)) {
  console.error('Please provide a valid integer as argument');
  process.exit(1);
}

// Function to check if a number is prime
function isPrime(n: number): boolean {
  if (n < 2) return false;
  if (n === 2) return true;
  if (n % 2 === 0) return false;

  for (let i = 3; i <= Math.sqrt(n); i += 2) {
    if (n % i === 0) return false;
  }
  return true;
}

// Find and print all primes up to N
for (let i = 2; i <= N; i++) {
  if (isPrime(i)) {
    console.log(i);
  }
}