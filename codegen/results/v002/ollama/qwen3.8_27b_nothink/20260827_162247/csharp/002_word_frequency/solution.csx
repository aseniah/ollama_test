using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

string text = File.ReadAllText("input/text.txt");

// Strip all punctuation (keep only letters), convert to lowercase
string cleaned = string.Concat(text.ToLower().Where(char.IsLetter));

// Split into words based on non-letter boundaries (original non-letters acted as separators)
// Actually we need to split by non-letter characters first, then filter
var words = text
    .Split(new char[] { '\r', '\n', ' ', '\t', '\r', '\n', ',', '.', '!', '?', ';', ':', '-', '"', '\'', '(', ')' }, StringSplitOptions.RemoveEmptyEntries)
    .Select(w => w.ToLower())
    .Where(w => w.Length > 0 && w.Any(char.IsLetter))
    .ToList();

// Count frequency
var freq = new Dictionary<string, int>();
foreach (var word in words)
{
    if (freq.ContainsKey(word))
        freq[word]++;
    else
        freq[word] = 1;
}

// Sort by count descending, then alphabetically ascending
var sorted = freq
    .OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key)
    .ToList();

foreach (var kv in sorted)
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}