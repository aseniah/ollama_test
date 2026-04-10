using System;
using System.Collections.Generic;

var n = int.Parse(Args[0]);

if (n < 1) {
    return;
}

int a = 1;
int b = 1;

Console.WriteLine(a);
Console.WriteLine(b);

while (b <= n) {
    int next = a + b;
    if (next <= n) {
        Console.WriteLine(next);
        a = b;
        b = next;
    } else {
        a = b;
        break;
    }
}