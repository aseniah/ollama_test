using System;
using System.Collections.Generic;

int N = Args.Count > 0 ? int.Parse(Args[0]) : 0;

if (N < 1)
{
    return;
}

List<int> fibonacci = new List<int>();
int a = 1, b = 1;

// Add the first Fibonacci number
if (a <= N)
{
    fibonacci.Add(a);
}

// Add the second Fibonacci number
if (b <= N)
{
    fibonacci.Add(b);
}

// Generate subsequent Fibonacci numbers
while (true)
{
    int next = a + b;
    if (next > N)
        break;
    
    fibonacci.Add(next);
    a = b;
    b = next;
}

// Print all Fibonacci numbers up to N
foreach (int num in fibonacci)
{
    Console.WriteLine(num);
}