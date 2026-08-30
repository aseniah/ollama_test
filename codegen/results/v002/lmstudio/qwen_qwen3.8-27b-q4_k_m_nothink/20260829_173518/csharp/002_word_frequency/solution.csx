using System;
using System.IO;
using System.Linq;
using System.Text;

string content = File.ReadAllText("input/text.txt");

// Get only letters (lowercase)
string lowercase = content.ToLower();
string lettersOnly = new string(lowercase.Where(c => char.IsLetter(c) || c == ' ' || c == '\n' || c == '\r' || c == '\t').ToArray());

// Split into words
var words = lettersOnly
    .Split(new[] { ' ', '\n', '\r', '\t' }, StringSplitOptions.RemoveEmptyEntries)
    .Where(w => w.Length > 0);

// Count frequencies
var freq = words
    .GroupBy(w => w)
    .ToDictionary(g => g.Key, g => g.Count());

// Sort by count descending, then alphabetically ascending
var sorted = freq
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key);

foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}