/**
 * Problem: Print all prime numbers up to and including N.
 * Input: Command-line argument N.
 * Output: Prime numbers one per line.
 */

function solve() {
    // Read the command line argument
    const arg = process.argv[2];
    if (arg === undefined) {
        return;
    }

    // Parse the input to an integer
    const n = parseInt(arg, 10);

    // If the input is not a number or less than 2, there are no primes to print
    if (isNaN(n) || n < 2) {
        return;
    }

    /**
     * Sieve of Eratosthenes algorithm to find all primes up to N.
     * This is more efficient than checking each number individually for larger N.
     */
    const isPrime = new Uint8Array(n + 1).fill(1);
    isPrime[0] = 0; // 0 is not prime
    isPrime[1] = 0; // 1 is not prime

    const limit = Math.sqrt(n);
    for (let p = 2; p <= limit; p++) {
        if (isPrime[p] === 1) {
            // Mark all multiples of p as non-prime starting from p*p
            for (let i = p * p; i <= n; i += p) {
                isPrime[i] = 0;
            }
        }
    }

    // Output the prime numbers
    let output = "";
    for (let p = 2; p <= n; p++) {
        if (isPrime[p] === 1) {
            output += p + "\n";
        }
        
        // To prevent memory issues with extremely large output strings, 
        // we flush the buffer periodically.
        if (output.length > 10000) {
            process.stdout.write(output);
            output = "";
        }
    }
    process.stdout.write(output);
}

solve();