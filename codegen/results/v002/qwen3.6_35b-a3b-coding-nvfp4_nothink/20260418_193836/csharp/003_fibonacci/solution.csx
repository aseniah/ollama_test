using System;

// Get N from command line arguments
if (Args.Count == 0)
{
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    return;
}

if (n < 1)
{
    return;
}

// Generate Fibonacci numbers up to n
int a = 1;
int b = 1;

// Print the first 1
Console.WriteLine(a);

// If n >= 1, print the second 1 as well (since we already printed the first one)
// But we need to be careful: if n is 1, we should only print 1 once? 
// The problem says "Start the sequence: 1, 1, 2, 3, 5, ..."
// So the sequence is 1, 1, 2, 3, 5, 8, ...
// If N=1, we should print both 1s? Let's check: "up to and including the largest Fibonacci number that does not exceed N"
// The Fibonacci numbers are 1, 1, 2, 3, 5, ...
// If N=1, the largest Fibonacci number <= 1 is 1. But there are two 1s in the sequence.
// The problem says "Print all Fibonacci numbers up to and including the largest Fibonacci number that does not exceed N."
// This could mean print all terms in the sequence that are <= N.
// So for N=1, we print 1, 1 (both terms that are <= 1).

// Let me re-read: "Print all Fibonacci numbers up to and including the largest Fibonacci number that does not exceed N."
// This is a bit ambiguous. It could mean:
// 1. Print all distinct Fibonacci numbers <= N
// 2. Print all terms in the Fibonacci sequence (with duplicates) that are <= N

// Given the sequence starts 1, 1, 2, 3, 5, ..., and the instruction says "One number per line", I think it means print each term in the sequence that is <= N.

// Let's generate the sequence and print each term that is <= n.

// Reset and do it properly
a = 1;
b = 1;

// We need to print terms as we generate them
// First term: 1
if (a <= n)
{
    Console.WriteLine(a);
}

// Second term: 1
if (b <= n)
{
    Console.WriteLine(b);
}

// Generate subsequent terms
while (true)
{
    int next = a + b;
    if (next > n)
    {
        break;
    }
    Console.WriteLine(next);
    a = b;
    b = next;
}