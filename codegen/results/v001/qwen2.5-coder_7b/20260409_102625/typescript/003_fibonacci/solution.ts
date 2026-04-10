// Importing necessary Node.js built-in modules
import { parseInt } from "lodash";

// Function to generate Fibonacci numbers up to a given limit
function fibonacciUpTo(limit: number): void {
    if (limit < 1) return;

    let a = 0, b = 1;
    while (a <= limit) {
        console.log(a);
        [a, b] = [b, a + b];
    }
}

// Reading the command-line argument
const args = process.argv.slice(2);
if (args.length !== 1) {
    console.error("Usage: tsx script.ts <integer>");
    process.exit(1);
}

const N = parseInt(args[0]);
fibonacciUpTo(N);