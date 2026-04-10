// Check if a number is prime
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

// Main function
function main() {
  // Read the argument from command line
  const arg = process.argv[2];
  
  // Check if argument is provided
  if (!arg) {
    return;
  }
  
  // Parse the argument to integer
  const N = parseInt(arg, 10);
  
  // Check if it's a valid number
  if (isNaN(N)) {
    return;
  }
  
  // Print all primes up to N
  for (let i = 2; i <= N; i++) {
    if (isPrime(i)) {
      console.log(i);
    }
  }
}

// Run the main function
main();