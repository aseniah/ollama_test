using System;

// Check if there's an argument provided
if (Args.Count == 0)
{
    return;
}

// Parse the input argument N. We use long to handle larger values of N and avoid overflow.
if (!long.TryParse(Args[0], out long n) || n < 1)
{
    // If N < 1, print nothing as per requirements.
    return;
}

// The Fibonacci sequence starts with 1, 1, 2, 3, 5...
// We initialize the first two numbers of the sequence.
long prev = 1;
long curr = 1;

// We first print the first '1' in the sequence.
Console.WriteLine(prev);

// The loop will handle all subsequent numbers in the sequence as long as they do not exceed N.
// The variable 'curr' represents the next number to be printed and evaluated.
while (curr <= n)
{
    Console.WriteLine(curr);
    
    // Calculate the next number in the Fibonacci sequence.
    long next = prev + curr;
    
    // Update the previous and current numbers for the next iteration.
    prev = curr;
    curr = next;
}