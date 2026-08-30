import * as process from 'node:process';

/**
 * Finds all prime numbers up to N using the Sieve of Eratosthenes algorithm.
 * This algorithm is efficient for finding all primes in a range up to a large N.
 * 
 * @param n The upper limit (inclusive) for finding prime numbers.
 */
function printPrimesUpTo(n: number): void {
    // Prime numbers are integers greater than 1. 
    // If n < 2, there are no prime numbers to print.
    if (n < 2) {
        return;
    }

    // We use a Uint8Array instead of a standard boolean array for memory efficiency.
    // Uint8Array is much lighter in memory when dealing with large values of N.
    // The array size is n + 1 so that the index directly corresponds to the number.
    const isPrime = new Uint8Array(n + 1).fill(1);
    
    // 0 and 1 are not prime numbers.
    isPrime[0] = 0;
    isPrime[1] = 0;

    // The Sieve of Eratosthenes:
    // We only need to iterate up to the square root of n.
    const limit = Math.sqrt(n);
    for (let p = 2; p <= limit; p++) {
        // If isPrime[p] is not changed, then it is a prime.
        if (isPrime[p] === 1) {
            // Update all multiples of p starting from p * p.
            // Multiples smaller than p * p have already been marked by smaller primes.
            for (let i = p * p; i <= n; i += p) {
                isPrime[i] = 0;
            }
        }
    }

    /**
     * Output Management:
     * Calling process.stdout.write or console.log for every single prime can be extremely
     * slow due to the overhead of system calls. To optimize, we buffer the output
     * and write to stdout in chunks.
     */
    let buffer = '';
    const BUFFER_LIMIT = 10000; // Number of primes to buffer before flushing to stdout.
    let count = 0;

    for (let i = 2; i <= n; i++) {
        if (isPrime[i] === 1) {
            buffer += i + '\n';
            count++;

            // If the buffer reaches the limit, flush it to the terminal.
            if (count >= BUFFER_LIMIT) {
                process.stdout.write(buffer);
                buffer = '';
                count = 0;
            }
        }
    }

    // Flush any remaining content in the buffer.
    if (buffer.length > 0) {
        process.stdout.write(buffer);
    }
}

/**
 * Entry point of the program.
 * Reads the command-line argument and invokes the prime generator.
 */
function main() {
    // process.argv[0] is the node executable path.
    // process.argv[1] is the script file path.
    // process.argv[2] is the first user-provided argument.
    const arg = process.argv[2];
    
    // If no argument is provided, exit silently.
    if (arg === undefined) {
        return;
    }

    // Parse the input argument to an integer.
    const n = parseInt(arg, 10);

    // If the argument is not a valid number, exit silently.
    if (isNaN(n)) {
        return;
    }

    printPrimesUpTo(n);
}

// Run the program.
main();