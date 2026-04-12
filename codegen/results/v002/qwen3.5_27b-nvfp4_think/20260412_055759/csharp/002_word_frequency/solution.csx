#r "System.Linq" // Optional explicit reference if needed, but System.Linq is usually included in .NET Core contexts

using System;
using System.IO;
using System.Linq;

string content = File.ReadAllText("input/text.txt");
content = content.ToLower();

// Replace non-letters and non-whitespace with space to preserve word boundaries while stripping punctuation
string cleaned = string.Concat(content.Select(c => (char.IsLetter(c) || char.IsWhiteSpace(c)) ? c : ' '));

// Split on whitespace, removing empty entries from multiple spaces
var words = cleaned.Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

// Count frequencies and sort
var result = words
    .GroupBy(w => w)
    .Select(g => new { Word = g.Key, Count = g.Count() })
    .OrderByDescending(x => x.Count)
    .ThenBy(x => x.Word);

// Output results
foreach (var item in result)
{
    Console.WriteLine($"{item.Word}: {item.Count}");
}