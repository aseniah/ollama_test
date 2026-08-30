using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

string filePath = "input/text.txt";

if (!File.Exists(filePath))
{
    return;
}

string text = File.ReadAllText(filePath);

// Process the text: keep only letters and spaces to separate words
StringBuilder sb = new StringBuilder();
foreach (char c in text)
{
    if (char.IsLetter(c))
    {
        sb.Append(char.ToLower(c));
    }
    else
    {
        sb.Append(' ');
    }
}

// Split by whitespace and remove empty entries
string[] words = sb.ToString().Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

// Count frequencies
var wordCounts = words
    .GroupBy(word => word)
    .Select(group => new 
    { 
        Word = group.Key, 
        Count = group.Count() 
    })
    // Sort by count descending, then by word alphabetically ascending
    .OrderByDescending(x => x.Count)
    .ThenBy(x => x.Word)
    .ToList();

// Output the results in the specified format
foreach (var item in wordCounts)
{
    Console.WriteLine($"{item.Word}: {item.Count}");
}