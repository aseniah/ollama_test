/**
 * This program accepts a single integer N from the command line 
 * and prints all prime numbers up to and including N, one per line.
 */

function main() {
    // Extract the argument from the command line
    const arg = process.argv[2];

    // If no argument is provided, we exit silently
    if (arg === undefined) {
        return;
    }

    // Parse the argument as an integer
    const N = parseInt(arg, 10);

    // Validate that the input is a number and at least 2 (the smallest prime)
    if (isNaN(N) || N < 2) {
        return;
    }

    /**
     * We use the Sieve of Eratosthenes algorithm to find all primes up to N.
     * This is highly efficient with a time complexity of O(N log log N).
     */
    
    // Create a boolean array (using Uint8Array for memory efficiency)
    // initialized to 1 (true), where index represents the number.
    const isPrime = new Uint8Array(N + 1).fill(1);
    
    // 0 and 1 are not prime numbers
    isPrime[0] = 0;
    isPrime[1] = 0;

    // Iterate from 2 up to the square root of N
    for (let p = 2; p * p <= N; p++) {
        // If isPrime[p] is still 1, then it is a prime
        if (isPrime[p] === 1) {
            // Mark all multiples of p starting from p*p as not prime
            for (let i = p * p; i <= N; i += p) {
                isPrime[i] = 0;
            }
        }
    }

    // Print all indices that remained marked as prime
    for (let p = 2; p <= N; p++) {
        if (isPrime[p] === 1) {
            process.stdout.write(p.toString() + '\n');
        }
    }
}

// Execute the main function
main();