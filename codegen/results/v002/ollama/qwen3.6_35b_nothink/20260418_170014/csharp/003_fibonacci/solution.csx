using System;

// Get N from command line argument
if (Args.Count < 1)
{
    return;
}

int n = int.Parse(Args[0]);

if (n < 1)
{
    return;
}

// Generate Fibonacci numbers up to and including the largest <= n
// Start: 1, 1, 2, 3, 5, ...
int a = 1;
int b = 1;

while (a <= n)
{
    Console.WriteLine(a);
    int next = a + b;
    a = b;
    b = next;
}