import { readline } from 'readline';

// Function to check if a number is prime
function isPrime(num: number): boolean {
    if (num <= 1) return false;
    if (num <= 3) return true;
    if (num % 2 === 0 || num % 3 === 0) return false;
    
    let i = 5;
    while (i * i <= num) {
        if (num % i === 0 || num % (i + 2) === 0) return false;
        i += 6;
    }
    return true;
}

// Main function to process command line argument
function main() {
    if (process.argv.length < 2) {
        console.error('Usage: node prime-printer.tsx <N>');
        return;
    }

    const N = parseInt(process.argv[2], 10);

    for (let i = 2; i <= N; i++) {
        if (isPrime(i)) {
            console.log(i);
        }
    }
}

// Execute the main function
main();