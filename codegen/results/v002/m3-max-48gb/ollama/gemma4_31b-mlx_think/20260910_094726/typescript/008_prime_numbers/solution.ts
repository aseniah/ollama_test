/**
 * The program finds all prime numbers up to and including N.
 * It uses the Sieve of Eratosthenes algorithm for efficiency.
 */

function main() {
    // Get the argument N from the command line
    const arg = process.argv[2];

    if (arg === undefined) {
        return;
    }

    const N = parseInt(arg, 10);

    // If N is not a number or is less than 2, there are no primes to print
    if (isNaN(N) || N < 2) {
        return;
    }

    // Sieve of Eratosthenes algorithm
    // Create a boolean array "isPrime[0..N]" and initialize
    // all entries it as true. A value in isPrime[i] will
    // finally be false if i is Not a prime, else true.
    const isPrime = new Uint8Array(N + 1).fill(1);
    isPrime[0] = 0;
    isPrime[1] = 0;

    for (let p = 2; p * p <= N; p++) {
        // If isPrime[p] is not changed, then it is a prime
        if (isPrime[p] === 1) {
            // Update all multiples of p starting from p*p
            // because smaller multiples have already been marked
            for (let i = p * p; i <= N; i += p) {
                isPrime[i] = 0;
            }
        }
    }

    // Collect and print all prime numbers
    let output = "";
    for (let p = 2; p <= N; p++) {
        if (isPrime[p] === 1) {
            output += p + "\n";
        }
        
        // Periodically flush the output string to avoid memory issues 
        // with extremely large N and very long strings
        if (output.length > 8192) {
            process.stdout.write(output);
            output = "";
        }
    }
    process.stdout.write(output);
}

main();