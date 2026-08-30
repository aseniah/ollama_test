// Read the command line argument
const input = process.argv[2];
if (!input || isNaN(Number(input))) {
    // No valid argument provided, exit without printing anything
    process.exit(0);
}

const N = parseInt(input, 10);

// Function to check if a number is prime
function isPrime(num: number): boolean {
    if (num < 2) return false;
    if (num === 2) return true;
    if (num % 2 === 0) return false;

    const limit = Math.floor(Math.sqrt(num));
    for (let i = 3; i <= limit; i += 2) {
        if (num % i === 0) return false;
    }
    return true;
}

// Print all primes up to and including N
for (let i = 2; i <= N; i++) {
    if (isPrime(i)) {
        console.log(i);
    }
}