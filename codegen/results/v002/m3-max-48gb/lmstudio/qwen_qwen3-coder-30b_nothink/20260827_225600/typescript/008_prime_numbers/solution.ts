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

// Read the command line argument
const input = process.argv[2];

// Validate that input is a number
if (!input || isNaN(parseInt(input))) {
    console.error("Please provide a valid integer argument");
    process.exit(1);
}

const N = parseInt(input);

// Generate and print all primes up to N
for (let i = 2; i <= N; i++) {
    if (isPrime(i)) {
        console.log(i);
    }
}