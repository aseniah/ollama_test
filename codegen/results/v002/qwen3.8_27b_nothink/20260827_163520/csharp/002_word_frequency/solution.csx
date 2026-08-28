using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

string text = File.ReadAllText("input/text.txt");

// Split into words
string[] words = text.Split(new[] { ' ', '\t', '\n', '\r', ',' }, StringSplitOptions.RemoveEmptyEntries);

// Process each word: lowercase, keep only letters
var wordList = new List<string>();
foreach (string rawWord in words)
{
    string word = rawWord.ToLower();
    // Keep only letters
    word = new string(word.Where(c => char.IsLetter(c)).ToArray());
    if (!string.IsNullOrEmpty(word))
    {
        wordList.Add(word);
    }
}

// Count frequency
var wordCounts = wordList.GroupBy(w => w)
                         .ToDictionary(g => g.Key, g => g.Count());

// Sort by count descending, then alphabetically ascending
var sortedWords = wordCounts
    .OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key)
    .ToList();

// Output
foreach (var kv in sortedWords)
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}