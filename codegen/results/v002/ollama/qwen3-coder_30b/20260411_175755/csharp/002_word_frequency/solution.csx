using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

var words = new Dictionary<string, int>();

foreach (var line in File.ReadAllLines("input/text.txt"))
{
    foreach (var word in line.Split(new char[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries))
    {
        var cleanWord = new string(word.Where(c => char.IsLetter(c)).ToArray()).ToLower();
        if (!string.IsNullOrEmpty(cleanWord))
        {
            if (words.ContainsKey(cleanWord))
                words[cleanWord]++;
            else
                words[cleanWord] = 1;
        }
    }
}

foreach (var kvp in words.OrderByDescending(x => x.Value).ThenBy(x => x.Key))
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}