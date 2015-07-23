# entropy-mnemonics

mnemonics is a golang package that converts byte slices into human-friendly
phrases. The primary purpose is to assist with the generation of
cryptographically secure passwords. The threshold for a cryptographically
secure password is between 128 and 256 bits, which when converted to base64 is
22-43 random characters. Random characters are both difficult to remember and
subject to error when being written down - smudging or sloppy handwriting can
make it difficult to recover a password.

mnemonics solves these problems by converting byte slices into simple and
common words. Take the following 128 bit example:

```
Hex:      a26a4821e36c7f7dccaa5484c080cefa
Base64:   ompIIeNsf33MqlSEwIDO+g==
Mnemonic: austere sniff aching hiding pact damp focus tacit timber pram left wonders
```

Though more verbose, the mnemonic phrase is less prone to errors when being
handled by humans.

The words are chosen from a dictionary of size 1626. (12 words is almost
exactly 128 bits of entropy). Each dictionary features a unique prefix length,
meaning all words in the dictionary have a prefix of length 'n' that is unique.
When decoding a passphrase, only the prefixes are checked. For the English
dictionary, the unique prefix length is 3. This means that passphrases can be
altered to make them more understandable or more easily memorized. For example,
the phrase "austere sniff aching" could safely be changed to "austere sniff
achoo" and the phrase would still decode correctly.

Full UTF-8 support is available for dictionaries, including input normalization
for inputs with (canonical equivalence)[https://en.wikipedia.org/wiki/Unicode_equivalence].

Supported Dictionaries:

+ English, Prefix Size 3
+ German,  Prefix Size 4
