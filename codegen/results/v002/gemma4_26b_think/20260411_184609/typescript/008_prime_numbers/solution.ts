/**
 * This program accepts a single integer N from the command line and 
 * prints all prime numbers up to and including N, one per line.
 * It uses the Sieve of Eratosthenes algorithm for efficient prime calculation.
 */

function main() {
    // Extract the first command-line argument
    const input = process.argv[2];

    // If no argument is provided, exit
    if (input === undefined) {
        return;
    }

    // Parse the argument as an integer
    const n = parseInt(input, 10);

    // If the input is not a number or less than 2, there are no primes to print.
    // Per requirements, if there are no primes, we print nothing.
    if (isNaN(n) || n < 2) {
        return;
    }

    /**
     * Sieve of Eratosthenes Algorithm
     * 
     * We use a Uint8Array to represent primality to be memory-efficient.
     * Uint8Array(n + 1) creates an array where indices correspond to integers.
     * 1 indicates a potential prime, 0 indicates a composite number.
     */
    const isPrime = new Uint8Array(n + 1).fill(1);

    // 0 and 1 are not prime numbers
    isPrime[0] = 0;
    isPrime[1] = 0;

    // Iterate through the array starting from the first prime, 2
    const limit = Math.sqrt(n);
    for (let p = 2; p <= limit; p++) {
        // If isPrime[p] is still 1, then it is a prime
        if (isPrime[p] === 1) {
            // Update all multiples of p starting from p * p
            // Multiples less than p * p would have already been marked by smaller primes
            for (let i = p * p; i <= n; i += p) {
                isPrime[i] = 0;
            }
        }
    }

    /**
     * Output the results.
     * We iterate through the sieve and print every index marked as prime.
     * Using console.log for each prime as required (one per line).
     */
    for (let i = 2; i <= n; i++) {
        if (isPrime[i] === 1) {
            process.stdout.write(i.toString() + '\n');
        }
    }
}

// Execute the main function
main();