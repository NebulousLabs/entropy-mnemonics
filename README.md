# entropy-mnemonics

mnemonics is a golang package that converts []byte's into human-friendly
phrases, using common words pulled from a dictionary. The dictionary size is
1626, and multiple languages are supported.  Each dictionary supports modified
phrases. Only the first few characters of each word are important. These
characters form a unique prefix. For example, in the English dictionary, the
unique prefix len (EnglishUniquePrefixLen) is 3, which means the word 'abbey'
could be replaced with the word 'abbot', and the program would still run as
expected.

The primary purpose of this library is creating human-friendly
cryptographically secure passwords. A cryptographically secure password
needs to contain between 128 and 256 bits of entropy. Humans are typically
incapable of generating sufficiently secure passwords without a random
number generator, and 256-bit random numbers tend to difficult to memorize
and even to write down (a single mistake in the writing, or even a single
somewhat sloppy character can render the backup useless).

By using a small set of common words instead of random numbers, copying
errors are more easily spotted and memorization is also easier, without
sacrificing password strength.

The mnemonics package does not have any functions for actually generating
entropy, it just converts existing entropy into human-friendly phrases.
