// Get the command line argument
const arg = process.argv[2];

// Parse the integer
const N = parseInt(arg, 10);

// Check if valid
if (isNaN(N) || N < 2) {
    // No primes to print
    process.exit(0);
}

// Function to check if a number is prime
function isPrime(num: number): boolean {
    if (num === 2) return true;
    if (num % 2 === 0) return false;
    
    // Check odd divisors up to sqrt(num)
    const limit = Math.sqrt(num);
    for (let i = 3; i <= limit; i += 2) {
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