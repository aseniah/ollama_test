// Function to check if a number is prime
function isPrime(num: number): boolean {
    if (num <= 1) return false;
    if (num <= 3) return true;

    if (num % 2 === 0 || num % 3 === 0) return false;

    for (let i = 5; i * i <= num; i += 6) {
        if (num % i === 0 || num % (i + 2) === 0) return false;
    }
    return true;
}

// Main function to find and print prime numbers up to N
function findAndPrintPrimes(N: number): void {
    for (let i = 2; i <= N; i++) {
        if (isPrime(i)) {
            console.log(i);
        }
    }
}

// Entry point of the program
if (process.argv.length < 3) {
    console.error("Usage: tsx script.ts <N>");
    process.exit(1);
}

const N = parseInt(process.argv[2], 10);

if (isNaN(N)) {
    console.error("Invalid input. Please provide a valid integer.");
    process.exit(1);
}

findAndPrintPrimes(N);