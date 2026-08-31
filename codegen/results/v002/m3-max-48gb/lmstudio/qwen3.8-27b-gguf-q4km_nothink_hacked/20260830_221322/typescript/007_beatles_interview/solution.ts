Looking at the expected output, I need to:
1. Parse the CSV file
2. Split names into First and Last (first word = FirstName, rest = LastName)
3. Calculate ages as of July 1, 2025 (for deceased people, use their death date instead)
4. Build the Relatives array from Father, Mother, Brother, Sister columns
5. Output as JSON array

Let me check the age calculations:
- John Lennon: born 1940, died 1980 → age at death = 40 ✓
- James McCartney: born 1942, alive → 2025-1942 = 83 ✓
- Ringo Starr: born 1940, alive → 2025-1940 = 85... but expected is 84. Wait, born 7/7/1940, as of 7/1/2025 he hasn't turned 85 yet (birthday is July 7), so age = 84 ✓
- George Harrison: born 1943, died 2001 → age at death = 58 ✓

So for deceased people, age is calculated at their death date. For living people, age is calculated as of July 1, 2025.

Now let me write the TypeScript code: