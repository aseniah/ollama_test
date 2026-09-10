using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;

string text = File.ReadAllText("input/text.txt");

var words = text
    .Split(new char[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries)
    .Select(w => new string(w.Where(char.IsLetter).ToArray()).ToLower())
    .Where(w => w.Length > 0);

var counts = words
    .GroupBy(w => w)
    .Select(g => new { Word = g.Key, Count = g.Count() })
    .OrderByDescending(x => x.Count)
    .ThenBy(x => x.Word)
    .ToList();

foreach (var item in counts)
{
    Console.WriteLine($"{item.Word}: {item.Count}");
}