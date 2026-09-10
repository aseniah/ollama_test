/**
 * This program accepts a single integer N from the command line
 * and prints all prime numbers up to and including N.
 */

function isPrime(num: number): boolean {
    if (num <= 1) return false;
    if (num <= 3) return true;
    if (num % 2 === 0 || num % 3 === 0) return false;

    // Optimization: Check divisors up to the square root of num
    // All primes are of the form 6k ± 1
    for (let i = 5; i * i <= num; i += 6) {
        if (num % i === 0 || num % (i + 2) === 0) {
            return false;
        }
    }
    return true;
}

function main() {
    // process.argv[0] is node executable, process.argv[1] is the script path
    const arg = process.argv[2];

    if (arg === undefined) {
        return;
    }

    const n = parseInt(arg, 10);

    // If the argument is not a valid number, do nothing
    if (isNaN(n)) {
        return;
    }

    // Iterate from 2 up to N and print if the number is prime
    for (let i = 2; i <= n; i++) {
        if (isPrime(i)) {
            process.stdout.write(i.toString() + '\n');
        }
    }
}

main();