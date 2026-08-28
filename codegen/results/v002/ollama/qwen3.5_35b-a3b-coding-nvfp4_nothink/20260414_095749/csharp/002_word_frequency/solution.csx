using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

string content = File.ReadAllText("input/text.txt");
string[] lines = content.Split(new[] { '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

var wordCounts = new Dictionary<string, int>();

foreach (var line in lines)
{
    // Split by whitespace
    var words = line.Split(new[] { ' ', '\t' }, StringSplitOptions.RemoveEmptyEntries);
    
    foreach (var word in words)
    {
        // Filter only letters and convert to lowercase
        var cleaned = new StringBuilder();
        foreach (char c in word)
        {
            if (char.IsLetter(c))
            {
                cleaned.Append(char.ToLower(c));
            }
        }
        
        string cleanWord = cleaned.ToString();
        
        if (!string.IsNullOrEmpty(cleanWord))
        {
            if (!wordCounts.ContainsKey(cleanWord))
            {
                wordCounts[cleanWord] = 0;
            }
            wordCounts[cleanWord]++;
        }
    }
}

// Sort by count descending, then alphabetically ascending
var sorted = wordCounts
    .OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key)
    .ToList();

foreach (var item in sorted)
{
    Console.WriteLine($"{item.Key}: {item.Value}");
}