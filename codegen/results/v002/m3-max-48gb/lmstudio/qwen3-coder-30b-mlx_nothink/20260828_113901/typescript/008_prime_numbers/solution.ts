// Function to check if a number is prime
function isPrime(num: number): boolean {
    if (num < 2) return false;
    if (num === 2) return true;
    if (num % 2 === 0) return false;
    
    // Check odd divisors up to sqrt(num)
    for (let i = 3; i * i <= num; i += 2) {
        if (num % i === 0) return false;
    }
    return true;
}

// Main function
function main() {
    // Read the command line argument
    const input = process.argv[2];
    
    // Validate that input is a number
    if (!input || isNaN(parseInt(input))) {
        return;
    }
    
    const N = parseInt(input);
    
    // Print all primes up to N
    for (let i = 2; i <= N; i++) {
        if (isPrime(i)) {
            console.log(i);
        }
    }
}

// Run the program
main();