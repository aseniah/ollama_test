using System;
using System.Collections.Generic;

int N = int.Parse(Args[0]);

if (N < 1)
{
    return;
}

List<int> fibonacci = new List<int>();
int a = 1, b = 1;

// Add the first number if it doesn't exceed N
if (a <= N)
{
    fibonacci.Add(a);
}

// Add the second number if it doesn't exceed N
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

// Print each Fibonacci number on its own line
foreach (int num in fibonacci)
{
    Console.WriteLine(num);
}