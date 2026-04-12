// Get N from command line arguments
const input = process.argv[2];
if (!input) {
    process.exit(0);
}

const n = parseInt(input, 10);

// Function to check if a number is prime
function isPrime(num: number): boolean {
    if (num < 2) return false;
    if (num === 2) return true;
    if (num % 2 === 0) return false;

    for (let i = 3; i * i <= num; i += 2) {
        if (num % i === 0) return false;
    }
    return true;
}

// Find and print all primes up to N
if (n >= 2) {
    for (let i = 2; i <= n; i++) {
        if (isPrime(i)) {
            console.log(i);
        }
    }
}