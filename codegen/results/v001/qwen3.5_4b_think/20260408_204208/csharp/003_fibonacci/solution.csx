using System;
using System.Collections.Generic;

var args = Args;
int n = int.Parse(args[0]);

if (n < 1)
{
    return;
}

long a = 1;
long b = 1;

Console.WriteLine(a);

while (b <= n)
{
    Console.WriteLine(b);
    long next = a + b;
    a = b;
    b = next;
}