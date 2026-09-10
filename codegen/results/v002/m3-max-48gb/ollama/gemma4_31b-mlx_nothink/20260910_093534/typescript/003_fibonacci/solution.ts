/**
 * This program calculates and prints the Fibonacci sequence 
 * up to the largest Fibonacci number that does not exceed N.
 */

function main() {
    // Get N from the command line arguments
    const arg = process.argv[2];

    if (arg === undefined) {
        return;
    }

    const N = parseInt(arg, 10);

    // Handle cases where N is not a number or N < 1
    if (isNaN(N) || N < 1) {
        return;
    }

    // Starting sequence: 1, 1, 2, 3, 5...
    let a = 1;
    let b = 1;

    // Print the first number
    console.log(a);

    // If N is 1, we are done after the first '1' (or both '1's based on sequence logic)
    // The prompt specifies the sequence starts 1, 1, 2...
    // So for N=1, it should print 1 then 1.
    if (N >= 1) {
        // The sequence starts with two 1s
        // We already printed the first 1.
        // Check if the second 1 fits.
        if (b <= N) {
            console.log(b);
        } else {
            return;
        }
    }

    // Calculate and print subsequent Fibonacci numbers
    while (true) {
        const next = a + b;
        if (next > N) {
            break;
        }
        console.log(next);
        a = b;
        b = next;
    }
}

main();