using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

var counts = new Dictionary<string, int>();
foreach (var line in File.ReadAllLines("input/text.txt"))
{
    var currentWord = new StringBuilder();
    foreach (char c in line)
    {
        if (char.IsLetter(c))
        {
            currentWord.Append(char.ToLower(c));
        }
        else if (currentWord.Length > 0)
        {
            var word = currentWord.ToString();
            counts.TryGetValue(word, out int count);
            counts[word] = count + 1;
            currentWord.Clear();
        }
    }
    if (currentWord.Length > 0)
    {
        var word = currentWord.ToString();
        counts.TryGetValue(word, out int count);
        counts[word] = count + 1;
    }
}

foreach (var kvp in counts.OrderByDescending(kvp => kvp.Value).ThenBy(kvp => kvp.Key))
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}