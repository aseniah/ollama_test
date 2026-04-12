import { argv } from 'process';

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

async function main(): Promise<void> {
    const n = parseInt(argv[2] || '0');

    if (isNaN(n)) {
        console.error('Please provide a valid integer argument.');
        return;
    }

    for (let i = 2; i <= n; i++) {
        if (isPrime(i)) {
            console.log(i);
        }
    }
}

main().catch((error) => {
    console.error('An error occurred:', error);
});