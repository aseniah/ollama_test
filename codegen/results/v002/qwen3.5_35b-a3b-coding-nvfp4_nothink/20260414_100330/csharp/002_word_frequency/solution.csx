using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

string[] lines = File.ReadAllLines("input/text.txt");
var wordCounts = new Dictionary<string, int>();

foreach (var line in lines)
{
    // Split by whitespace
    var tokens = line.Split((char[])null, StringSplitOptions.RemoveEmptyEntries);
    
    foreach (var token in tokens)
    {
        // Keep only letters (strip punctuation)
        var cleanWord = new string(token.Where(char.IsLetter).ToArray());
        
        if (cleanWord.Length > 0)
        {
            var lowerWord = cleanWord.ToLower();
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
}

// Sort by count descending, then alphabetically ascending
var sorted = wordCounts
    .OrderByDescending(x => x.Value)
    .ThenBy(x => x.Key)
    .ToList();

foreach (var item in sorted)
{
    Console.WriteLine($"{item.Key}: {item.Value}");
}