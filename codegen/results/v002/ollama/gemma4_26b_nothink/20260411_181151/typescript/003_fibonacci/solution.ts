/**
 * Problem: Print all Fibonacci numbers up to and including the largest 
 * Fibonacci number that does not exceed N.
 * Sequence starts: 1, 1, 2, 3, 5, ...
 * 
 * Execution: tsx script.ts <N>
 */

import process from 'node:process';

function solve(): void {
    // Extract command line arguments
    const args = process.argv.slice(2);

    if (args.length === 0) {
        return;
    }

    // Parse N from the first argument
    const nArg = args[0];
    const N = parseInt(nArg, 10);

    // Validation: If N is not a number or N < 1, print nothing as per instructions.
    if (isNaN(N) || N < 1) {
        return;
    }

    // Fibonacci sequence logic
    // We start with the first two numbers of the sequence: 1, 1
    let a = 1;
    let b = 1;

    // Handle the first 1 explicitly if N >= 1
    // Since the sequence starts 1, 1, we print the first 1.
    console.log(a);

    // If N is 1, we print the second 1 and stop.
    if (N >= 1) {
        // Check if we should print the second 1
        // Since the sequence is 1, 1, 2... and N >= 1, we always print at least the first 1.
        // If N=1, the loop below won't execute if we handle the second 1 carefully.
        
        // Let's use a more robust approach for the sequence:
        // Let's track the current value being printed.
        let current = 1;
        let next = 1;

        // We already printed the first '1' above.
        // Let's rewrite the loop logic for clarity.
    }
}

/**
 * Re-implementing the logic cleanly for the final output.
 */
function printFibonacci(N: number): void {
    if (N < 1) return;

    let first = 1;
    let second = 1;

    // Print the first element
    console.log(first);

    // If N is at least 1, the second element (also 1) is valid
    if (N >= 1) {
        // We check if the second element is <= N. 
        // Since N >= 1, the second 1 is always valid.
        // However, we must ensure we don't print the same instance 
        // if N were somehow less than 1 (but we already checked N < 1).
        
        // To avoid infinite loop and handle the "1, 1" start:
        let prev = 1;
        let curr = 1;
        
        // We already printed 'first'. Now print 'second' and then continue.
        console.log(curr);

        // Calculate subsequent numbers
        while (true) {
            let next = prev + curr;
            if (next > N) {
                break;
            }
            console.log(next);
            prev = curr;
            curr = next;
        }
    }
}

/**
 * Corrected single-pass approach
 */
function run(): void {
    const args = process.argv.slice(2);
    if (args.length === 0) return;

    const N = parseInt(args[0], 10);
    if (isNaN(N) || N < 1) return;

    let a = 1;
    let b = 1;

    // The sequence starts 1, 1, 2, 3...
    // We print 'a' as long as 'a' <= N
    // We need to handle the dual '1's at the start.
    
    // Sequence: F(1)=1, F(2)=1, F(3)=2...
    // We iterate through the indices.
    
    let currentVal = 1;
    let nextVal = 1;
    
    // We use a simple loop to generate the sequence.
    // We print the first 1.
    console.log(1);
    
    // If N >= 1, the second 1 is also part of the sequence.
    if (N >= 1) {
        // We only print the second 1 if we haven't finished.
        // But the sequence is 1, 1, 2... 
        // If N=1, output is 1 \n 1.
        
        // Let's track the sequence manually.
        let sequence: number[] = [1, 1];
        
        // If N is 1, we just print 1, 1.
        if (N === 1) {
            // We already printed the first 1.
            console.log(1);
            return;
        }

        // If N > 1, we continue the sequence.
        let p = 1;
        let c = 1;
        // We already printed the first 1. Now print the second 1.
        console.log(c);

        while (true) {
            let next = p + c;
            if (next > N) break;
            console.log(next);
            p = c;
            c = next;
        }
    }
}

// Final logic implementation
const main = () => {
    const args = process.argv.slice(2);
    if (args.length === 0) return;
    const N = parseInt(args[0], 10);
    if (isNaN(N) || N < 1) return;

    let a = 1;
    let b = 1;

    // Print the first 1
    console.log(a);
    
    // Check if second 1 is valid
    if (N >= 1) {
        // The problem states: 1, 1, 2, 3, 5...
        // If N=1, output should be:
        // 1
        // 1
        // (Because 1 <= 1)
        
        // We need to print the second 1.
        // But we must be careful not to print the first 1 twice in the same logic block.
        // Let's use a simple state-based approach.
        
        let current = 1; // The 2nd element
        let prev = 1;    // The 1st element
        
        // We already printed the 1st element.
        // Now check the 2nd element.
        if (current <= N) {
            // To avoid duplicating the logic for the very first 1:
            // We only enter here if we are handling elements starting from the 2nd.
            // But we must not print the first 1 AGAIN.
            // Let's reset logic.
        }
    }
}

// Simple, robust implementation:
function finalSolve(): void {
    const args = process.argv.slice(2);
    if (args.length === 0) return;
    const N = parseInt(args[0], 10);
    if (isNaN(N) || N < 1) return;

    let sequence: number[] = [1, 1];
    
    // We handle the base case of the first two elements manually
    // then expand the sequence.
    
    // First element
    console.log(1);
    
    // Second element (if it exists and is <= N)
    if (N >= 1) {
        // Since N >= 1 is already checked, the second 1 is always valid.
        // We need to ensure we don't print it if we already printed it.
        // Wait, the sequence is 1, 1, 2... 
        // If N=1, output is 1 then 1.
        
        // Let's track how many 1s we've printed.
        // But the loop logic is simpler:
        let p = 1;
        let c = 1;
        
        // We already printed the first 1.
        // Now print the second 1.
        console.log(c);
        
        // Now calculate the rest.
        while (true) {
            let next = p + c;
            if (next > N) break;
            console.log(next);
            p = c;
            c = next;
        }
    }
}

// Re-running logic to ensure precision:
// If N=1: 
//   Print 1 (first 1)
//   Print 1 (second 1)
//   Next is 2. 2 > 1, so stop.
// Result: 1, 1. Correct.

// If N=5:
//   Print 1 (first 1)
//   Print 1 (second 1)
//   Next is 2. 2 <= 5. Print 2.
//   Next is 3. 3 <= 5. Print 3.
//   Next is 5. 5 <= 5. Print 5.
//   Next is 8. 8 > 5. Stop.
// Result: 1, 1, 2, 3, 5. Correct.

// The only edge case: what if the user wants only one '1' if N=1? 
// "Print all Fibonacci numbers up to... largest... that does not exceed N."
// The sequence is 1, 1, 2, 3, 5...
// If N=1, both 1s are <= 1. So both should be printed.

function execute() {
    const args = process.argv.slice(2);
    if (args.length === 0) return;
    const N = parseInt(args[0], 10);
    if (isNaN(N) || N < 1) return;

    let p = 1;
    let c = 1;

    // Print first element
    console.log(p);

    // Print second element if it's <= N
    if (c <= N) {
        console.log(c);
        
        // Compute next elements
        while (true) {
            let next = p + c;
            if (next > N) break;
            console.log(next);
            p = c;
            c = next;
        }
    }
}

execute();