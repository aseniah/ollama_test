// Function to generate all prime numbers up to a given integer N
function sieveOfEratosthenes(n: number): number[] {
    const isPrime = new Array(n + 1).fill(true);
    isPrime[0] = isPrime[1] = false; // 0 and 1 are not prime numbers

    for (let p = 2; p * p <= n; p++) {
        if (isPrime[p]) {
            for (let i = p * p; i <= n; i += p) {
                isPrime[i] = false;
            }
        }
    }

    const primes: number[] = [];
    for (let i = 2; i <= n; i++) {
        if (isPrime[i]) {
            primes.push(i);
        }
    }

    return primes;
}

// Main function to handle command-line arguments and print prime numbers
function main() {
    if (process.argv.length !== 3) {
        console.error("Usage: tsx program.ts <integer>");
        process.exit(1);
    }

    const n = parseInt(process.argv[2], 10);
    if (isNaN(n) || n < 2) {
        console.error("Argument must be a positive integer greater than or equal to 2.");
        process.exit(1);
    }

    const primes = sieveOfEratosthenes(n);
    primes.forEach(prime => console.log(prime));
}

// Run the main function
main();