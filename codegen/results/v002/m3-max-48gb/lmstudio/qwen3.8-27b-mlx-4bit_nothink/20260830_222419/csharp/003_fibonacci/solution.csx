using System;
using System.Collections.Generic;

if (Args.Count != 1) {
    return;
}

if (!int.TryParse(Args[0], out int N)) {
    return;
}

if (N < 1) {
    return;
}

List<string> results = new List<string>();
long a = 1, b = 1;

// Handle the two 1s
if (a <= N) results.Add(a.ToString());
if (b <= N) results.Add(b.ToString());

long next;
while (true) {
    next = a + b;
    if (next > N) break;
    results.Add(next.ToString());
    a = b;
    b = next;
}

foreach (var r in results) {
    Console.WriteLine(r);
}