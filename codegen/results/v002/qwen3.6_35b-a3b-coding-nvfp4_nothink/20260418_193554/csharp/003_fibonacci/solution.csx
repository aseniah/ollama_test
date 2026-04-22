using System;

// Get N from command line arguments
if (Args.Count < 1)
{
    return;
}

int N;
if (!int.TryParse(Args[0], out N))
{
    return;
}

if (N < 1)
{
    return;
}

// Generate Fibonacci numbers up to N
int a = 1;
int b = 1;

// Print the first 1
Console.WriteLine(a);

// If N >= 1, print the second 1 as well (since we start with 1, 1)
if (N >= 1)
{
    Console.WriteLine(b);
}

// Generate subsequent Fibonacci numbers
while (true)
{
    int next = a + b;
    if (next > N)
    {
        break;
    }
    Console.WriteLine(next);
    a = b;
    b = next;
}