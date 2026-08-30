#r "System.Numerics"

using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

// Validate argument count
if (Args.Count != 1)
{
    // Exit if no arguments provided (per instructions, accept a single integer argument)
    return;
}

// Parse the integer argument
if (!int.TryParse(Args[0], out int n))
{
    return;
}

// If N < 1, print nothing
if (n < 1)
{
    return;
}

// Generate Fibonacci numbers
// Sequence: 1, 1, 2, 3, 5, 8...
// We use long to prevent overflow for N near int.MaxValue
long a = 1;
long b = 1;

// Print the first Fibonacci number (1)
if (a <= n)
{
    Console.WriteLine(a);
}

// Check if the second Fibonacci number (1) should be printed
// Note: The sequence is 1, 1, ... so if n >= 1, we print 1 twice.
// However, the logic below handles the generation. 
// We need to be careful not to print the second '1' if we already printed the first one 
// and the loop logic changes 'a'. 
// Let's stick to the specific sequence definition: 1, 1, 2, 3, 5...
// If N=1, output should be "1\n1".

// Let's reset and iterate strictly.
long current = 1;
long next = 1;

while (current <= n)
{
    Console.WriteLine(current);
    
    // Calculate next term
    long sum = current + next;
    current = next;
    next = sum;
}