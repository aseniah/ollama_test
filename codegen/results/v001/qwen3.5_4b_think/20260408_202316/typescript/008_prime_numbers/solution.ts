// Read command-line argument
const N = parseInt(process.argv[2]);

if (isNaN(N) || N < 1) {
    process.exit(0); // Exit if argument is not valid or less than 1
}

// Function to check if a number is prime
function isPrime(num: number): boolean {
    if (num < 2) {
        return false;
    }
    if (num === 2) {
        return true;
    }
    if (num % 2 === 0) {
        return false;
    }
    for (let i = 3; i * i <= num; i += 2) {
        if (num % i === 0) {
            return false;
        }
    }
    return true;
}

// Collect primes from 2 up to N
let primes: number[] = [];
for (let i = 2; i <= N; i++) {
    if (isPrime(i)) {
        primes.push(i);
    }
}

// Print primes, one per line
primes.forEach(prime => {
    console.log(prime);
});