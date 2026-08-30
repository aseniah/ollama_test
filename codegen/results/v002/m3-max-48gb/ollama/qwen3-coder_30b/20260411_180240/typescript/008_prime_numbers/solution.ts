function isPrime(num: number): boolean {
    if (num < 2) return false;
    if (num === 2) return true;
    if (num % 2 === 0) return false;
    
    for (let i = 3; i <= Math.sqrt(num); i += 2) {
        if (num % i === 0) return false;
    }
    return true;
}

function printPrimes(n: number): void {
    for (let i = 2; i <= n; i++) {
        if (isPrime(i)) {
            console.log(i);
        }
    }
}

// Read command line argument
const arg = process.argv[2];
if (!arg) {
    process.exit(1);
}

const n = parseInt(arg, 10);
if (isNaN(n)) {
    process.exit(1);
}

printPrimes(n);