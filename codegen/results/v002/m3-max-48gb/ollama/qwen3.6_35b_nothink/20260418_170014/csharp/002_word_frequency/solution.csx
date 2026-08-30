using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;

// Read the file
var text = File.ReadAllText("input/text.txt");

// Split into words, convert to lowercase
var words = text.ToLower().Split(new char[] { ' ', ',', '.', ';', ':', '!', '?', '\n', '\r', '\t' }, StringSplitOptions.RemoveEmptyEntries);

// Strip punctuation (keep only letters)
var cleanedWords = new List<string>();
foreach (var word in words)
{
    var cleaned = new string(word.Where(char.IsLetter).ToArray());
    if (cleaned.Length > 0)
    {
        cleanedWords.Add(cleaned);
    }
}

// Count frequency
var frequency = new Dictionary<string, int>();
foreach (var word in cleanedWords)
{
    if (frequency.ContainsKey(word))
    {
        frequency[word]++;
    }
    else
    {
        frequency[word] = 1;
    }
}

// Sort by count descending, then alphabetically ascending for ties
var sorted = frequency.OrderByDescending(x => x.Value).ThenBy(x => x.Key);

// Output
foreach (var entry in sorted)
{
    Console.WriteLine($"{entry.Key}: {entry.Value}");
}