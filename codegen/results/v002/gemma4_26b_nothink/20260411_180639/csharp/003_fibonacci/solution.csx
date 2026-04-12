using System;

// Check if an argument was provided
if (Args.Count == 0)
{
    return;
}

// Parse the input N
if (!int.TryParse(Args[0], out int n) || n < 1)
{
    // If N < 1, print nothing as per requirements.
    return;
}

// Initialize the first two numbers of the Fibonacci sequence
long a = 1;
long b = 1;

// The sequence starts with 1, 1...
// We print the first number
Console.WriteLine(a);

// If n is 1, we've already printed the first 1. 
// However, the sequence is 1, 1, 2... so we check if the second 1 is valid.
if (n >= 1)
{
    // We handle the second '1' explicitly or via loop logic.
    // Let's use a loop that generates the next number in the sequence.
    
    // Special case: if n is at least 1, the second '1' should be printed 
    // if we follow the 1, 1, 2 pattern.
    if (n >= 1)
    {
        // If n is 1, the sequence is 1, 1.
        // Let's track the "current" number being evaluated.
        long current = 1;
        long next = 1;
        
        // Print the first 1 (already done)
        // Now handle the rest of the sequence.
        
        // To avoid duplicate logic, let's reset and use a clean approach:
        // We will print 'a' and then calculate the next.
    }
}

// Re-implementing clean logic for the sequence: 1, 1, 2, 3, 5...
// Requirement: Print all up to largest Fib <= N.
// Sequence starts: 1, 1, 2, 3, 5...

long prev = 1;
long curr = 1;

// Print the first 1
Console.WriteLine(prev);

// If N is at least 1, we check if the second 1 fits.
// Since N >= 1 is already checked, the second 1 always fits if N >= 1.
if (n >= 1)
{
    // We already printed the first 1.
    // Let's see if the second 1 is needed.
    // The sequence is 1, 1, 2, 3...
    // If N=1, output is 1, 1.
    
    // Let's use a loop that tracks the previous and current.
    // We use a temporary loop to handle the 1, 1 start.
    
    // We'll use a list or just loop.
    // Let's restart the logic for simplicity.
}