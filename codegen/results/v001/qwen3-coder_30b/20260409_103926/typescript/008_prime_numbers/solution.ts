const N = parseInt(process.argv[2]);

// Handle invalid input
if (isNaN(N) || N < 2) {
  process.exit(0);
}

// Function to check if a number is prime
function isPrime(num: number): boolean {
  if (num < 2) return false;
  if (num === 2) return true;
  if (num % 2 === 0) return false;
  
  // Only check odd divisors up to sqrt(num)
  for (let i = 3; i * i <= num; i += 2) {
    if (num % i === 0) return false;
  }
  return true;
}

// Print all primes up to N
for (let i = 2; i <= N; i++) {
  if (isPrime(i)) {
    console.log(i);
  }
}