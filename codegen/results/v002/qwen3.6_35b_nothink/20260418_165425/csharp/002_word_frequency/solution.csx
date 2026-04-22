using System;
using System.Collections.Generic;
using System.Linq;

string content = File.ReadAllText("input/text.txt");
string[] lines = content.Split(new[] { '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

Dictionary<string, int> wordCounts = new Dictionary<string, int>();

foreach (string line in lines)
{
    string[] words = line.Split(new[] { ' ', ',', '.', ';', ':', '!', '?', '-', '\t' }, StringSplitOptions.RemoveEmptyEntries);
    foreach (string word in words)
    {
        // Strip all punctuation (keep only letters)
        string cleaned = new string(word.Where(c => char.IsLetter(c)).ToArray());
        if (cleaned.Length == 0) continue;
        
        // Convert to lowercase
        string lowerWord = cleaned.ToLowerInvariant();
        
        if (wordCounts.ContainsKey(lowerWord))
        {
            wordCounts[lowerWord]++;
        }
        else
        {
            wordCounts[lowerWord] = 1;
        }
    }
}

// Sort by count descending, then alphabetically ascending
var sorted = wordCounts.OrderBy(kvp => kvp.Key).OrderByDescending(kvp => kvp.Count).ToList();

foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}