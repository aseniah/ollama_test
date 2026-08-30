using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

var lines = File.ReadAllLines("input/text.txt");
var wordCounts = new Dictionary<string, int>();

foreach (var line in lines)
{
    // Split by spaces and punctuation markers to get potential words
    var tokens = line.Split(new char[] { ' ', ',', ';', ':', '!', '?', '.', '\n', '\t', '\r' }, StringSplitOptions.RemoveEmptyEntries);
    
    foreach (var token in tokens)
    {
        // Build a clean word containing only letters
        var cleanWord = new StringBuilder();
        foreach (var c in token)
        {
            if (char.IsLetter(c))
            {
                cleanWord.Append(char.ToLower(c));
            }
        }

        var word = cleanWord.ToString();
        
        if (word.Length > 0)
        {
            if (wordCounts.ContainsKey(word))
            {
                wordCounts[word]++;
            }
            else
            {
                wordCounts[word] = 1;
            }
        }
    }
}

// Sort by count descending, then by word ascending
var sorted = wordCounts
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key);

foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}