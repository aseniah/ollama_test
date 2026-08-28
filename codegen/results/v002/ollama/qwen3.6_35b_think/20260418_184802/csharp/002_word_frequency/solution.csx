using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

var text = File.ReadAllText("input/text.txt");
var counts = new Dictionary<string, int>();

foreach (var word in text.Split(' ', '\t', '\n', '\r', StringSplitOptions.RemoveEmptyEntries))
{
    var cleaned = new string(word.ToLowerInvariant().Where(c => c >= 'a' && c <= 'z').ToArray());
    if (cleaned.Length > 0)
    {
        counts.TryGetValue(cleaned, out var count);
        counts[cleaned] = count + 1;
    }
}

foreach (var kv in counts.OrderByDescending(x => x.Value).ThenBy(x => x.Key))
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}