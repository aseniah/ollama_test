// Read N from command line argument (Node.js compatible)
const nStr = process.argv[2];
if (!nStr || nStr.trim() === '') {
  console.error('Error: Missing required argument N');
  process.exit(1);
}

const N = parseInt(nStr, 10);
if (Number.isNaN(N) || N < 0) {
  console.error('Error: Invalid number, must be a non-negative integer');
  process.exit(1);
}

// Function to check if a number is prime
function isPrime(num: number): boolean {
  if (num <= 1) return false;
  if (num === 2) return true;
  if (num % 2 === 0) return false; // Even numbers > 2 are not prime

  for (let i = 3; i <= Math.sqrt(num); i += 2) {
    if (num % i === 0) return false;
  }
  return true;
}

// Collect all prime numbers up to and including N
const primes: number[] = [];
for (let i = 2; i <= N; i++) {
  if (isPrime(i)) {
    primes.push(i);
  }
}

// Print each prime on a separate line
if (primes.length > 0) {
  for (const prime of primes) {
    console.log(prime);
  }
} else {
  // No output if no primes found, even when N is small (e.g., 0 or 1)
}