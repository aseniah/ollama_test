// Read N from command line
const n = parseInt(process.argv[2], 10);

if (isNaN(n)) {
    // No valid argument, print nothing
    process.exit(0);
}

// If N < 2, there are no primes
if (n < 2) {
    process.exit(0);
}

// Sieve of Eratosthenes
const sieve = new Array(n + 1).fill(true);
sieve[0] = false;
if (n >= 1) sieve[1] = false;

for (let i = 2; i * i <= n; i++) {
    if (sieve[i]) {
        for (let j = i * i; j <= n; j += i) {
            sieve[j] = false;
        }
    }
}

// Print all primes
let output = '';
for (let i = 2; i <= n; i++) {
    if (sieve[i]) {
        output += i + '\n';
    }
}

if (output.length > 0) {
    process.stdout.write(output);
}