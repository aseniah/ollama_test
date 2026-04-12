/**
 * This program takes a single integer N from the command line
 * and prints all prime numbers up to and including N, one per line.
 * It uses the Sieve of Eratosthenes algorithm for efficiency.
 */

import process from 'node:process';

function solve(): void {
    // Get arguments from command line
    const args = process.argv.slice(2);

    if (args.length === 0) {
        return;
    }

    const nStr = args[0];
    const N = parseInt(nStr, 10);

    // Validate input: must be a number and not NaN
    if (isNaN(N)) {
        return;
    }

    // If N is less than 2, there are no prime numbers to print
    if (N < 2) {
        return;
    }

    // Sieve of Eratosthenes
    // Create a boolean array "isPrime[0..N]" and initialize
    // all entries it as true. A value in isPrime[i] will
    // finally be false if i is Not a prime, else true.
    const isPrime: boolean[] = new Array(N + 1).fill(true);
    isPrime[0] = false;
    isPrime[1] = false;

    const limit = Math.sqrt(N);

    for (let p = 2; p <= limit; p++) {
        // If isPrime[p] is not changed, then it is a prime
        if (isPrime[p]) {
            // Update all multiples of p starting from p*p
            // Numbers smaller than p*p have already been marked
            for (let i = p * p; i <= N; i += p) {
                isPrime[i] = false;
            }
        }
    }

    // Collect and print all primes
    let output = '';
    for (let p = 2; p <= N; p++) {
        if (isPrime[p]) {
            output += p + '\n';
            
            // Periodically flush output to stdout to manage large N 
            // and avoid huge string concatenation memory overhead
            if (output.length > 10000) {
                process.stdout.write(output);
                output = '';
            }
        }
    }
    
    process.stdout.write(output);
}

// Run the solver
solve();