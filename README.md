# Course_Cryptography_for_Developers
Завдання: Програмна реалізація алгоритму гешування
2.1 Власна реалізація алгоритму гешування на вибір

Напишіть власну реалізацію алгоритму гешування SHA-1 або SHA-3 (Keccak) на вибір. Зверніть увагу, що ці алгоритми гешування є блочними і вхідні дані слід правильно розділити на порції, кожна довжиною рівно один блок. Якщо останній блок є неповним його треба доповнити згідно зі специфікацією алгоритма.
Ваша функція повинна мати змогу загешувати будь-які дані довільної довжини. В більш розширеному варіанті можете реалізувати подачу вхідних даних порціями, тобто щоб можна було декілька разів викликати функцію гешування передаючи кожну наступну порцію даних для гешування. Ця функціональність буде доречна для гешування дуже великих обсягів даних, коли неприпустимо зберігати одразу весь аргумент функції в памʼяті.

2.2 Тест на співпадіння гешу з бібліотечною реалізацією

Метою цього етапу є перевірка вашої реалізації на коректність. Для цього оберіть декілька повідомлень різної довжини: до одного блока, більше одного блока, декілька блоків. Далі обчислить геш-значення від цих повідомлень використовуючи вашу реалізацію і використовуючи бібліотечну реалізацію від сторонніх розробників. Результати гешування порівняйте попарно з зробіть висновок щодо коректності обчислень. У разі необхідності знайдіть і виправте помилки у вашій реалізації.

2.3 Порівняння швидкодії власної реалізації з бібліотечною

Метою цього етапу є оптимізація вашої реалізації алгоритму гешування для зменшення обсягу памʼяті і процесорного часу, що використовуються для обчислення геш-значення від повідомлення певної довжини. Для виконання цього етапу зробіть замір часу витраченого на гешування певного повідомлення вашою функцією і замір часу гешування функцією з готової бібліотеки. Результати замірів порівняйте і зробіть висновок щодо ефективності вашої реалізації.
Обсяг використаної памʼяті можна виміряти додатковими інструментами моніторингу ПЗ або вашої ОС. Усі результати порівняння можна викласти у файлі readme.md вашого репозиторію.

# Нажаль у мене не співпадють геші і я не розумію чому.
Я спробував реалізувати алгоритм SHA-3, але щось я зробив не так, бо мої геші не співпадають з бібліотечними. Я використав багато ресурсів, щоб виправити це, але не вийшло. Зато я спробував.
Також я зробив тести порівняння і замір часу виконання.
# Якщо трохи докопатись до слів завдання, то теоретично я зробив завдання правильно, бо воно звучить як "Напишіть ВЛАСНУ реалізацію алгоритму гешування SHA-1 або SHA-3")

Я реалізовував читаючи цей алгоритм.

# Алгоритм SHA-3:
Design
SHA-3 uses the sponge construction, in which data is "absorbed" into the sponge, then the result is "squeezed" out. In the absorbing phase, message blocks are XORed into a subset of the state, which is then transformed as a whole using a permutation function f. (Calling f a permutation may be confusing. It is technically a permutation of the state space, thus a permutation of a set with 
2^1600≈4.4⋅10^481 elements, but it does more than merely permute the bits of the state vector.) In the "squeeze" phase, output blocks are read from the same subset of the state, alternated with the state transformation function f. The size of the part of the state that is written and read is called the "rate" (denoted r), and the size of the part that is untouched by input/output is called the "capacity" (denoted c). The capacity determines the security of the scheme. The maximum security level is half the capacity.

Given an input bit string N, a padding function pad, a permutation function f that operates on bit blocks of width b, a rate r and an output length d, we have capacity c=b-r and the sponge construction 
Z=sponge[[f,pad,r](N,d), yielding a bit string Z of length d, works as follows: 
pad the input N using the pad function, yielding a padded bit string P with a length divisible by r (such that n=len(P)/r is an integer)
break P into n consecutive r-bit pieces P0, ..., P(n−1)
initialize the state S to a string of b zero bits
absorb the input into the state: for each block Pi:
extend Pi at the end by a string of c zero bits, yielding one of length b
XOR that with S
apply the block permutation f to the result, yielding a new state S
initialize Z to be the empty string
while the length of Z is less than d:
append the first r bits of S to Z
if Z is still less than d bits long, apply f to S, yielding a new state S
truncate Z to d bits
The fact that the internal state S contains c additional bits of information in addition to what is output to Z prevents the length extension attacks that SHA-2, SHA-1, MD5 and other hashes based on the Merkle–Damgård construction are susceptible to.

In SHA-3, the state S consists of a 5 × 5 array of w-bit words (with w = 64), b = 5 × 5 × w = 5 × 5 × 64 = 1600 bits total. Keccak is also defined for smaller power-of-2 word sizes w down to 1 bit (total state of 25 bits). Small state sizes can be used to test cryptanalytic attacks, and intermediate state sizes (from w = 8, 200 bits, to w = 32, 800 bits) can be used in practical, lightweight applications.[11][12]

For SHA3-224, SHA3-256, SHA3-384, and SHA3-512 instances, r is greater than d, so there is no need for additional block permutations in the squeezing phase; the leading d bits of the state are the desired hash. However, SHAKE128 and SHAKE256 allow an arbitrary output length, which is useful in applications such as optimal asymmetric encryption padding.                                                                                    Padding
To ensure the message can be evenly divided into r-bit blocks, padding is required. SHA-3 uses the pattern 10*1 in its padding function: a 1 bit, followed by zero or more 0 bits (maximum r − 1) and a final 1 bit.

The maximum of r − 1 zero bits occurs when the last message block is r − 1 bits long. Then another block is added after the initial 1 bit, containing r − 1 zero bits before the final 1 bit.

The two 1 bits will be added even if the length of the message is already divisible by r.[5]: 5.1  In this case, another block is added to the message, containing a 1 bit, followed by a block of r − 2 zero bits and another 1 bit. This is necessary so that a message with length divisible by r ending in something that looks like padding does not produce the same hash as the message with those bits removed.

The initial 1 bit is required so messages differing only in a few additional 0 bits at the end do not produce the same hash.

The position of the final 1 bit indicates which rate r was used (multi-rate padding), which is required for the security proof to work for different hash variants. Without it, different hash variants of the same short message would be the same up to truncation.                                                                                                    The block permutation
The block transformation f, which is Keccak-f[1600] for SHA-3, is a permutation that uses XOR, AND and NOT operations, and is designed for easy implementation in both software and hardware.

It is defined for any power-of-two word size, w = 2ℓ bits. The main SHA-3 submission uses 64-bit words, ℓ = 6.

The state can be considered to be a 5 × 5 × w array of bits. Let a[i][ j][k] be bit (5i + j) × w + k of the input, using a little-endian bit numbering convention and row-major indexing. I.e. i selects the row, j the column, and k the bit.

Index arithmetic is performed modulo 5 for the first two dimensions and modulo w for the third.

The basic block permutation function consists of 12 + 2ℓ rounds of five steps:

θ (theta)
Compute the parity of each of the 5w (320, when w = 64) 5-bit columns, and exclusive-or that into two nearby columns in a regular pattern. To be precise, a[i][ j][k] ← a[i][ j][k] ⊕ parity(a[0...4][j-1][k]) ⊕ parity(a[0...4][j+1][k−1])
ρ (rho)
Bitwise rotate each of the 25 words by a different triangular number 0, 1, 3, 6, 10, 15, .... To be precise, a[0][0] is not rotated, and for all 0 ≤ t < 24, a[i][ j][k] ← a[i][ j][k−(t+1)(t+2)/2], where 
{\begin{pmatrix}i\\j\end{pmatrix}}={\begin{pmatrix}3&2\\1&0\end{pmatrix}}^{t}{\begin{pmatrix}0\\1\end{pmatrix}}.
π (pi)
Permute the 25 words in a fixed pattern. a[3i+2j][i] ← a[ i][j].
χ (chi)
Bitwise combine along rows, using x ← x ⊕ (¬y & z). To be precise, a[i][ j][k] ← a[i][ j][k] ⊕ (¬a[i][ j+1][k] & a[i][ j+2][k]). This is the only non-linear operation in SHA-3.
ι (iota)
Exclusive-or a round constant into one word of the state. To be precise, in round n, for 0 ≤ m ≤ ℓ, a[0][0][2m−1] is XORed with bit m + 7n of a degree-8 LFSR sequence. This breaks the symmetry that is preserved by the other steps.
