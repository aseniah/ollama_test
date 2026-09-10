using System;

if (Args.Count < 1) return;

int N = int.Parse(Args[0]);

if (N < 1) return;

long a = 1, b = 1;

Console.WriteLine(a);
Console.WriteLine(b);

while (true)
{
    long next = a + b;
    if (next > N) break;
    Console.WriteLine(next);
    a = b;
    b = next;
}