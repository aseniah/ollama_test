using System;

if (Args.Count < 1)
    return;

long N = long.Parse(Args[0]);

if (N < 1)
    return;

long a = 1, b = 1;
while (a <= N)
{
    Console.WriteLine(a);
    // If a + b would exceed N (or overflow), stop after printing a
    if (a > N - b)
        break;
    long temp = a + b;
    a = b;
    b = temp;
}