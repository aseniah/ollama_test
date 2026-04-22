using System;

// Get N from command line
if (Args.Count < 1)
{
    Console.WriteLine("Usage: Please provide a single integer argument N.");
    return;
}

int n = int.Parse(Args[0]);

// If N < 1, print nothing
if (n < 1)
{
    return;
}

// Generate Fibonacci numbers up to n
int a = 1;
int b = 1;

while (a <= n)
{
    Console.WriteLine(a);
    int temp = a + b;
    a = b;
    b = temp;
}