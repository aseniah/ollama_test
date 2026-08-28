using System;
using System.Text;

if (Args.Count < 1)
    return;

if (!long.TryParse(Args[0], out long n))
    return;

if (n < 1)
    return;

long a = 1, b = 1;
while (true)
{
    Console.WriteLine(a);
    if (a > n / b)  // would overflow
        break;
    long next = a + b;
    if (next > n)
        break;
    a = b;
    b = next;
    // Wait, I need to restructure. Let me redo this properly.
}