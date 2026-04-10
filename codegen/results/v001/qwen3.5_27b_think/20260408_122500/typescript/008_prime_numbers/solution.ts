// Read command line argument
const nStr = process.argv[2];
const n = parseInt(nStr, 10);

// Validate input
if (!nStr || isNaN(n)) {
    process.exit(0);
}

// Helper function to check if a number is prime
function isPrime(num: number): boolean {
    if (num < 2) return false;
    if (num === 2) return true;
    if (num % 2 === 0) return false;

    // Check odd divisors up to sqrt(num)
    const limit = Math.floor(Math.sqrt(num));
    for (let i = 3; i <= limit; i += 2) {
        if (num % i === 0) return false;
    }
    return true;
}

// Find and print all primes up to n
for (let i = 2; i <= n; i++) {
    if (isPrime(i)) {
        console.log(i);
    }
}