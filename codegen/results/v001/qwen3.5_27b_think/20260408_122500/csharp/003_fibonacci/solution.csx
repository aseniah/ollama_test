using System;

if (Args.Count < 1 || !int.TryParse(Args[0], out int n)) {
    return;
}

if (n < 1) {
    return;
}

int a = 1;
int b = 1;

while (a <= n) {
    Console.WriteLine(a);
    int next = a + b;
    a = b;
    b = next;
}