using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

string filePath = "input/text.txt";
string[] lines = File.ReadAllLines(filePath);

Dictionary<string, int> wordCounts = new Dictionary<string, int>();

foreach (string line in lines)
{
    // Split line into tokens by whitespace
    string[] tokens = line.Split(new char[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);
    
    foreach (string token in tokens)
    {
        // Convert to lowercase
        string lowerToken = token.ToLower();
        
        // Strip all punctuation, keep only letters
        string cleanWord = new string(lowerToken.Where(char.IsLetter).ToArray());
        
        if (cleanWord.Length > 0)
        {
            if (wordCounts.ContainsKey(cleanWord))
            {
                wordCounts[cleanWord]++;
            }
            else
            {
                wordCounts[cleanWord] = 1;
            }
        }
    }
}

// Sort by count descending, then by word ascending
var sortedWords = wordCounts.OrderByDescending(kvp => kvp.Value)
                            .ThenBy(kvp => kvp.Key);

foreach (var kvp in sortedWords)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}