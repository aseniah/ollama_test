using System;
using System.Collections.Generic;

var Args = Environment.GetCommandLineArgs();
if (Args.Length != 1) {
    Console.Error.WriteLine("Usage: dotnet-script <N>");
    Environment.Exit(1);
}

int N = int.Parse(Args[0]);

if (N < 1) {
    return;
}

var a = 1;
var b = 1;

Console.WriteLine(a);
Console.WriteLine(b);

while (true) {
    var next = a + b;
    if (next > N) {
        break;
    }
    Console.WriteLine(next);
    a = b;
    b = next;
}

return null;