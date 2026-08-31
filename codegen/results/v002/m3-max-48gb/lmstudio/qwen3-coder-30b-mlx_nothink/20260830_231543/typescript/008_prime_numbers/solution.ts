// Function to check if a number is prime
function isPrime(num: number): boolean {
    if (num < 2) return false;
    if (num === 2) return true;
    if (num % 2 === 0) return false;
    
    for (let i = 3; i <= Math.sqrt(num); i += 2) {
        if (num % i === 0) return false;
    }
    return true;
}

// Read command line argument
const input = process.argv[2];
if (!input) {
    process.exit(0);
}

const N = parseInt(input, 10);
if (isNaN(N)) {
    process.exit(0);
}

// Find and print all primes up to N
for (let i = 2; i <= N; i++) {
    if (isPrime(i)) {
        console.log(i);
    }
}