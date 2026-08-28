/**
 * This program prints all prime numbers up to and including a given integer N.
 * It uses the Sieve of Eratosthenes algorithm for efficient prime calculation.
 */

function main() {
    // Get the argument from the command line
    const arg = process.argv[2];

    if (arg === undefined) {
        return;
    }

    const n = parseInt(arg, 10);

    // Validate that the input is a number and at least 2 (since primes start at 2)
    if (isNaN(n) || n < 2) {
        return;
    }

    // Sieve of Eratosthenes
    // Create a boolean array "isPrime[0..n]" and initialize
    // all entries it as true. A value in isPrime[i] will
    // finally be false if i is Not a prime, else true.
    const isPrime = new Uint8Array(n + 1).fill(1);
    isPrime[0] = 0;
    isPrime[1] = 0;

    for (let p = 2; p * p <= n; p++) {
        // If isPrime[p] is not changed, then it is a prime
        if (isPrime[p] === 1) {
            // Update all multiples of p starting from p*p
            for (let i = p * p; i <= n; i += p) {
                isPrime[i] = 0;
            }
        }
    }

    // Print all prime numbers
    let output = "";
    for (let p = 2; p <= n; p++) {
        if (isPrime[p] === 1) {
            output += p + "\n";
            
            // To prevent memory overflow with extremely large N 
            // when building the string, we can print in chunks.
            if (output.length > 10000) {
                process.stdout.write(output);
                output = "";
            }
        }
    }
    process.stdout.write(output);
}

main();