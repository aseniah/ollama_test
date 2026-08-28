using System;

var n = int.Parse(Args[0]);

if (n < 1)
{
    return;
}

long a = 1;
long b = 1;

while (a <= n)
{
    Console.WriteLine(a);
    long temp = a + b;
    a = b;
    b = temp;
}