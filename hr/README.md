# Hacker Rank Problems

This lesson picks up two challenges from HackerRank.

[**Camel Problem**](https://www.hackerrank.com/challenges/camelcase/problem) - Find the number of words in a given camelcase variable, e.g. `hackerRankProblem`.

[**Cipher Problem**](https://www.hackerrank.com/challenges/caesar-cipher-1/problem) - Generate encrypted text given a key.


## Test Results

```
$ go test -v
=== RUN   TestCountWord
=== RUN   TestCountWord/santoshKumar
=== RUN   TestCountWord/endOfTheWorld
=== RUN   TestCountWord/coronaVirusPandemic
--- PASS: TestCountWord (0.00s)
    --- PASS: TestCountWord/santoshKumar (0.00s)
    --- PASS: TestCountWord/endOfTheWorld (0.00s)
    --- PASS: TestCountWord/coronaVirusPandemic (0.00s)
=== RUN   TestCaesarCipher
=== RUN   TestCaesarCipher/test_with_single_word
=== RUN   TestCaesarCipher/test_with_spaces_in_between
=== RUN   TestCaesarCipher/test_with_other_special_characters
=== RUN   TestCaesarCipher/test_for_rotating_letter
--- PASS: TestCaesarCipher (0.00s)
    --- PASS: TestCaesarCipher/test_with_single_word (0.00s)
    --- PASS: TestCaesarCipher/test_with_spaces_in_between (0.00s)
    --- PASS: TestCaesarCipher/test_with_other_special_characters (0.00s)
    --- PASS: TestCaesarCipher/test_for_rotating_letter (0.00s)
PASS
ok      github.com/santosh/gophercises/hr       0.004s
```
