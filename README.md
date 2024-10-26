# Keywords

A list and count of keywords in programming languages. Based on work originally
compiled by [@leighmcculloch](https://github.com/leighmcculloch) and contributors.

_2022 update (WIP):_ I took the latest versions of several popularity lists and am ensuring the top 10 in all three lists were on this list. There have been some pretty significant changes in the past couple of years, including the addition of some _soft_ or _contextual_ keywords in Python 3 and Java, increasing their keyword counts significantly.

In addition to the above, I migrated CI to GitHub Actions and am working to ensure multiple language versions with different keyword sets are represented.

![Languages by keyword](chart.png)

## Why does it matter?

The number of keywords in a programming language _can_ be an indication to it's simplicity/complexity, and that can impact the simplicity/complexity of the solutions that developers produce using it. Complex solutions can be more expensive to maintain and difficult to hire for. However, this is dependent on many factors and keyword count is only one; language idioms also play a massive part.

## License

Source code in this repository is licensed under the MIT License.

Compiled data in this repository is licensed under the Creative Commons Attribution 4.0 International Public License.

`SPDX-License-Identifier: MIT AND CC-BY-4.0`

## Contribute

Don't see a language here? Please open a pull request adding it!

## Keyword List

* [C (ANSI (C89)) (32 keywords)](#c-ansi-c89-32-keywords)
* [C (C11) (44 keywords)](#c-c11-44-keywords)
* [C (C17) (44 keywords)](#c-c17-44-keywords)
* [C (C99) (37 keywords)](#c-c99-37-keywords)
* [C# (8.0) (107 keywords)](#c-80-107-keywords)
* [C++ (C++03) (74 keywords)](#c-c-03-74-keywords)
* [C++ (C++11) (84 keywords)](#c-c-11-84-keywords)
* [C++ (C++14) (84 keywords)](#c-c-14-84-keywords)
* [C++ (C++17) (84 keywords)](#c-c-17-84-keywords)
* [C++ (C++20) (92 keywords)](#c-c-20-92-keywords)
* [C++ (C++98) (74 keywords)](#c-c-98-74-keywords)
* [Dart (2.2) (33 keywords)](#dart-22-33-keywords)
* [Elixir (1.10) (15 keywords)](#elixir-110-15-keywords)
* [Erlang (23) (27 keywords)](#erlang-23-27-keywords)
* [Fortran (Fortran 2008) (103 keywords)](#fortran-fortran-2008-103-keywords)
* [Go (1.18) (25 keywords)](#go-118-25-keywords)
* [Java (SE 11 LTS) (51 keywords)](#java-se-11-lts-51-keywords)
* [Java (SE 17 LTS) (67 keywords)](#java-se-17-lts-67-keywords)
* [JavaScript (1st edition) (35 keywords)](#javascript-1st-edition-35-keywords)
* [JavaScript (2nd edition) (59 keywords)](#javascript-2nd-edition-59-keywords)
* [JavaScript (3rd edition) (59 keywords)](#javascript-3rd-edition-59-keywords)
* [JavaScript (5th edition) (45 keywords)](#javascript-5th-edition-45-keywords)
* [JavaScript (6th edition) (46 keywords)](#javascript-6th-edition-46-keywords)
* [Kotlin (1.4) (79 keywords)](#kotlin-14-79-keywords)
* [Lua (5.3) (22 keywords)](#lua-53-22-keywords)
* [MATLAB (R2020a) (20 keywords)](#matlab-r2020a-20-keywords)
* [Objective-C (2.0) (85 keywords)](#objective-c-20-85-keywords)
* [PHP (7.4) (69 keywords)](#php-74-69-keywords)
* [Python 2 (2.7) (31 keywords)](#python-2-27-31-keywords)
* [Python 3 (3.10) (38 keywords)](#python-3-310-38-keywords)
* [R (4.0) (21 keywords)](#r-40-21-keywords)
* [Ruby (2.7) (41 keywords)](#ruby-27-41-keywords)
* [Rust (1.46) (53 keywords)](#rust-146-53-keywords)
* [Scala (2.13) (40 keywords)](#scala-213-40-keywords)
* [Swift (5.3) (97 keywords)](#swift-53-97-keywords)
* [Visual Basic (2019) (217 keywords)](#visual-basic-2019-217-keywords)

### C (ANSI (C89)) (32 keywords)
| | | | |
|---|---|---|---|
|  auto | break | case | char |
|  const | continue | default | do |
|  double | else | enum | extern |
|  float | for | goto | if |
|  int | long | register | return |
|  short | signed | sizeof | static |
|  struct | switch | typedef | union |
|  unsigned | void | volatile | while |


#### Sources:

[http://port70.net/~nsz/c/c89/c89-draft.html#3.1.1](http://port70.net/~nsz/c/c89/c89-draft.html#3.1.1)
### C (C11) (44 keywords)
| | | | |
|---|---|---|---|
|  auto | break | case | char |
|  const | continue | default | do |
|  double | else | enum | extern |
|  float | for | goto | if |
|  inline | int | long | register |
|  restrict | return | short | signed |
|  sizeof | static | struct | switch |
|  typedef | union | unsigned | void |
|  volatile | while | _Alignas | _Alignof |
|  _Atomic | _Bool | _Complex | _Generic |
|  _Imaginary | _Noreturn | _Static_assert | _Thread_local |


#### Sources:

[https://www.open-std.org/JTC1/SC22/WG14/www/docs/n1570.pdf](https://www.open-std.org/JTC1/SC22/WG14/www/docs/n1570.pdf)
### C (C17) (44 keywords)
| | | | |
|---|---|---|---|
|  auto | break | case | char |
|  const | continue | default | do |
|  double | else | enum | extern |
|  float | for | goto | if |
|  inline | int | long | register |
|  restrict | return | short | signed |
|  sizeof | static | struct | switch |
|  typedef | union | unsigned | void |
|  volatile | while | _Alignas | _Alignof |
|  _Atomic | _Bool | _Complex | _Generic |
|  _Imaginary | _Noreturn | _Static_assert | _Thread_local |


#### Sources:

[https://www.open-std.org/JTC1/SC22/WG14/www/docs/n2310.pdf](https://www.open-std.org/JTC1/SC22/WG14/www/docs/n2310.pdf)
### C (C99) (37 keywords)
| | | | |
|---|---|---|---|
|  auto | break | case | char |
|  const | continue | default | do |
|  double | else | enum | extern |
|  float | for | goto | if |
|  inline | int | long | register |
|  restrict | return | short | signed |
|  sizeof | static | struct | switch |
|  typedef | union | unsigned | void |
|  volatile | while | _Bool | _Complex |
|  _Imaginary |


#### Sources:

[https://www.open-std.org/JTC1/SC22/WG14/www/docs/n1256.pdf](https://www.open-std.org/JTC1/SC22/WG14/www/docs/n1256.pdf)
### C# (8.0) (107 keywords)
| | | | |
|---|---|---|---|
|  abstract | as | base | bool |
|  break | byte | case | catch |
|  char | checked | class | const |
|  continue | decimal | default | delegate |
|  do | double | else | enum |
|  event | explicit | extern | false |
|  finally | fixed | float | for |
|  foreach | goto | if | implicit |
|  in | int | interface | internal |
|  is | lock | long | namespace |
|  new | null | object | operator |
|  out | override | params | private |
|  protected | public | readonly | ref |
|  return | sbyte | sealed | short |
|  sizeof | stackalloc | static | string |
|  struct | switch | this | throw |
|  true | try | typeof | uint |
|  ulong | unchecked | unsafe | ushort |
|  using | virtual | void | volatile |
|  while | add | alias | ascending |
|  async | await | by | descending |
|  dynamic | equals | from | get |
|  global | group | into | join |
|  let | nameof | notnull | on |
|  orderby | partial | remove | select |
|  set | unmanaged | value | var |
|  when | where | yield |


#### Sources:

[https://standards.iso.org/ittf/PubliclyAvailableStandards/c075178_ISO_IEC_23270_2018.zip](https://standards.iso.org/ittf/PubliclyAvailableStandards/c075178_ISO_IEC_23270_2018.zip)
### C++ (C++03) (74 keywords)
| | | | |
|---|---|---|---|
|  asm | auto | bool | break |
|  case | catch | char | class |
|  const | const_cast | continue | default |
|  delete | do | double | dynamic_cast |
|  else | enum | explicit | export |
|  extern | false | float | for |
|  friend | goto | if | inline |
|  int | long | mutable | namespace |
|  new | operator | private | protected |
|  public | register | reinterpret_cast | return |
|  short | signed | sizeof | static |
|  static_cast | struct | switch | template |
|  this | throw | true | try |
|  typedef | typeid | typename | union |
|  unsigned | using | virtual | void |
|  volatile | wchar_t | while | and |
|  and_eq | bitand | bitor | compl |
|  not | not_eq | or | or_eq |
|  xor | xor_eq |


#### Sources:

[http://staff.ustc.edu.cn/~zhuang/cpp/specs/ISO_IEC%2014882%202003.pdf](http://staff.ustc.edu.cn/~zhuang/cpp/specs/ISO_IEC%2014882%202003.pdf)
### C++ (C++11) (84 keywords)
| | | | |
|---|---|---|---|
|  alignas | alignof | asm | auto |
|  bool | break | case | catch |
|  char | char_16t | char_32t | class |
|  const | constexpr | const_cast | continue |
|  decltype | default | delete | do |
|  double | dynamic_cast | else | enum |
|  explicit | export | extern | false |
|  float | for | friend | goto |
|  if | inline | int | long |
|  mutable | namespace | new | noexcept |
|  nullptr | operator | private | protected |
|  public | register | reinterpret_cast | return |
|  short | signed | sizeof | static |
|  static_assert | static_cast | struct | switch |
|  template | this | thread_local | throw |
|  true | try | typedef | typeid |
|  typename | union | unsigned | using |
|  virtual | void | volatile | wchar_t |
|  while | and | and_eq | bitand |
|  bitor | compl | not | not_eq |
|  or | or_eq | xor | xor_eq |


#### Sources:

[https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2011/n3242.pdf](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2011/n3242.pdf)
### C++ (C++14) (84 keywords)
| | | | |
|---|---|---|---|
|  alignas | alignof | asm | auto |
|  bool | break | case | catch |
|  char | char_16t | char_32t | class |
|  const | constexpr | const_cast | continue |
|  decltype | default | delete | do |
|  double | dynamic_cast | else | enum |
|  explicit | export | extern | false |
|  float | for | friend | goto |
|  if | inline | int | long |
|  mutable | namespace | new | noexcept |
|  nullptr | operator | private | protected |
|  public | register | reinterpret_cast | return |
|  short | signed | sizeof | static |
|  static_assert | static_cast | struct | switch |
|  template | this | thread_local | throw |
|  true | try | typedef | typeid |
|  typename | union | unsigned | using |
|  virtual | void | volatile | wchar_t |
|  while | and | and_eq | bitand |
|  bitor | compl | not | not_eq |
|  or | or_eq | xor | xor_eq |


#### Sources:

[https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2011/n3242.pdf](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2011/n3242.pdf)
### C++ (C++17) (84 keywords)
| | | | |
|---|---|---|---|
|  alignas | alignof | asm | auto |
|  bool | break | case | catch |
|  char | char16_t | char32_t | class |
|  const | constexpr | const_cast | continue |
|  decltype | default | delete | do |
|  double | dynamic_cast | else | enum |
|  explicit | export | extern | false |
|  float | for | friend | goto |
|  if | inline | int | long |
|  mutable | namespace | new | noexcept |
|  nullptr | operator | private | protected |
|  public | register | reinterpret_cast | return |
|  short | signed | sizeof | static |
|  static_assert | static_cast | struct | switch |
|  template | this | thread_local | throw |
|  true | try | typedef | typeid |
|  typename | union | unsigned | using |
|  virtual | void | volatile | wchar_t |
|  while | and | and_eq | bitand |
|  bitor | compl | not | not_eq |
|  or | or_eq | xor | xor_eq |


#### Sources:

[http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2017/n4659.pdf](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2017/n4659.pdf)
### C++ (C++20) (92 keywords)
| | | | |
|---|---|---|---|
|  alignas | alignof | asm | auto |
|  bool | break | case | catch |
|  char | char8_t | char16_t | char32_t |
|  class | concept | const | consteval |
|  constexpr | constinit | const_cast | continue |
|  co_await | co_return | co_yield | decltype |
|  default | delete | do | double |
|  dynamic_cast | else | enum | explicit |
|  export | extern | false | float |
|  for | friend | goto | if |
|  inline | int | long | mutable |
|  namespace | new | noexcept | nullptr |
|  operator | private | protected | public |
|  register | reinterpret_cast | requires | return |
|  short | signed | sizeof | static |
|  static_assert | static_cast | struct | switch |
|  template | this | thread_local | throw |
|  true | try | typedef | typeid |
|  typename | union | unsigned | using |
|  virtual | void | volatile | wchar_t |
|  while | and | and_eq | bitand |
|  bitor | compl | not | not_eq |
|  or | or_eq | xor | xor_eq |


#### Sources:

[https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/n4849.pdf](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/n4849.pdf)
### C++ (C++98) (74 keywords)
| | | | |
|---|---|---|---|
|  asm | auto | bool | break |
|  case | catch | char | class |
|  const | const_cast | continue | default |
|  delete | do | double | dynamic_cast |
|  else | enum | explicit | export |
|  extern | false | float | for |
|  friend | goto | if | inline |
|  int | long | mutable | namespace |
|  new | operator | private | protected |
|  public | register | reinterpret_cast | return |
|  short | signed | sizeof | static |
|  static_cast | struct | switch | template |
|  this | throw | true | try |
|  typedef | typeid | typename | union |
|  unsigned | using | virtual | void |
|  volatile | wchar_t | while | and |
|  and_eq | bitand | bitor | compl |
|  not | not_eq | or | or_eq |
|  xor | xor_eq |


#### Sources:

[https://www.lirmm.fr/~ducour/Doc-objets/ISO+IEC+14882-1998.pdf](https://www.lirmm.fr/~ducour/Doc-objets/ISO+IEC+14882-1998.pdf)
### Dart (2.2) (33 keywords)
| | | | |
|---|---|---|---|
|  assert | break | case | catch |
|  class | const | continue | default |
|  do | else | enum | extends |
|  false | final | finally | for |
|  if | in | is | new |
|  null | rethrow | return | super |
|  switch | this | throw | true |
|  try | var | void | while |
|  with |


#### Sources:

[https://dart.dev/guides/language/specifications/DartLangSpec-v2.2.pdf](https://dart.dev/guides/language/specifications/DartLangSpec-v2.2.pdf)
### Elixir (1.10) (15 keywords)
| | | | |
|---|---|---|---|
|  true | false | nil | when |
|  and | or | not | in |
|  fn | do | end | catch |
|  rescue | after | else |


#### Sources:

[https://hexdocs.pm/elixir/syntax-reference.html#reserved-words](https://hexdocs.pm/elixir/syntax-reference.html#reserved-words)
### Erlang (23) (27 keywords)
| | | | |
|---|---|---|---|
|  after | and | andalso | band |
|  begin | bnot | bor | bsl |
|  bsr | bxor | case | catch |
|  cond | div | end | fun |
|  if | let | not | of |
|  or | orelse | receive | rem |
|  try | when | xor |


#### Sources:

[http://erlang.org/doc/reference_manual/introduction.html#reserved-words](http://erlang.org/doc/reference_manual/introduction.html#reserved-words)
### Fortran (Fortran 2008) (103 keywords)
| | | | |
|---|---|---|---|
|  abstract | allocatable | allocate | assign |
|  associate | asynchronous | backspace | bind |
|  block | block data | call | case |
|  class | close | codimension | common |
|  contains | contiguous | continue | critical |
|  cycle | data | deallocate | deferred |
|  dimension | do | do concurrent | elemental |
|  else | else if | elsewhere | end |
|  endfile | endif | entry | enum |
|  enumerator | equivalence | error stop | exit |
|  extends | external | final | flush |
|  forall | format | function | generic |
|  goto | if | implicit | import |
|  include | inquire | intent | interface |
|  intrinsic | lock | module | namelist |
|  non_overridable | nopass | nullify | only |
|  open | operator | optional | parameter |
|  pass | pause | pointer | print |
|  private | procedure | program | protected |
|  public | pure | read | recursive |
|  result | return | rewind | rewrite |
|  save | select | sequence | stop |
|  submodule | subroutine | sync all | sync images |
|  sync memory | target | then | unlock |
|  use | value | volatile | wait |
|  where | while | write |


#### Sources:

[http://fortranwiki.org/fortran/show/Keywords](http://fortranwiki.org/fortran/show/Keywords)
### Go (1.18) (25 keywords)
| | | | |
|---|---|---|---|
|  break | case | chan | const |
|  continue | default | defer | else |
|  fallthrough | for | func | go |
|  goto | if | import | interface |
|  map | package | range | return |
|  select | struct | switch | type |
|  var |


#### Sources:

[https://golang.org/ref/spec#Keywords](https://golang.org/ref/spec#Keywords)
### Java (SE 11 LTS) (51 keywords)
| | | | |
|---|---|---|---|
|  abstract | assert | boolean | break |
|  byte | case | catch | char |
|  class | const | continue | default |
|  do | double | else | enum |
|  extends | final | finally | float |
|  for | if | goto | implements |
|  import | instanceof | int | interface |
|  long | native | new | package |
|  private | protected | public | return |
|  short | static | strictfp | super |
|  switch | synchronized | this | throw |
|  throws | transient | try | void |
|  volatile | while | _ |


#### Sources:

[https://docs.oracle.com/javase/specs/jls/se11/html/jls-3.html#jls-3.9](https://docs.oracle.com/javase/specs/jls/se11/html/jls-3.html#jls-3.9)
### Java (SE 17 LTS) (67 keywords)
| | | | |
|---|---|---|---|
|  abstract | assert | boolean | break |
|  byte | case | catch | char |
|  class | const | continue | default |
|  do | double | else | enum |
|  extends | final | finally | float |
|  for | if | goto | implements |
|  import | instanceof | int | interface |
|  long | native | new | package |
|  private | protected | public | return |
|  short | static | strictfp | super |
|  switch | synchronized | this | throw |
|  throws | transient | try | void |
|  volatile | while | _ | exports |
|  module | non-sealed | open | opens |
|  permits | provides | record | requires |
|  sealed | to | transitive | uses |
|  var | with | yield |


#### Sources:

[https://docs.oracle.com/javase/specs/jls/se17/html/jls-3.html#jls-3.9](https://docs.oracle.com/javase/specs/jls/se17/html/jls-3.html#jls-3.9)
### JavaScript (1st edition) (35 keywords)
| | | | |
|---|---|---|---|
|  break | continue | delete | else |
|  for | function | if | in |
|  new | return | this | typeof |
|  var | void | while | with |
|  case | catch | class | const |
|  debugger | default | do | enum |
|  export | extends | finally | import |
|  super | switch | throw | try |
|  null | true | false |


#### Sources:

[https://www.ecma-international.org/wp-content/uploads/ECMA-262_1st_edition_june_1997.pdf](https://www.ecma-international.org/wp-content/uploads/ECMA-262_1st_edition_june_1997.pdf)
### JavaScript (2nd edition) (59 keywords)
| | | | |
|---|---|---|---|
|  break | continue | delete | else |
|  for | function | if | in |
|  new | return | this | typeof |
|  var | void | while | with |
|  abstract | boolean | byte | case |
|  catch | char | class | const |
|  debugger | default | do | double |
|  enum | export | extends | final |
|  finally | float | goto | implements |
|  import | instanceof | int | interface |
|  long | native | package | private |
|  protected | public | short | static |
|  super | switch | synchronized | throw |
|  throws | transient | try | volatile |
|  null | true | false |


#### Sources:

[https://www.ecma-international.org/wp-content/uploads/ECMA-262_2nd_edition_august_1998.pdf](https://www.ecma-international.org/wp-content/uploads/ECMA-262_2nd_edition_august_1998.pdf)
### JavaScript (3rd edition) (59 keywords)
| | | | |
|---|---|---|---|
|  break | case | catch | continue |
|  default | delete | do | else |
|  finally | for | function | if |
|  in | instanceof | new | return |
|  switch | this | throw | try |
|  typeof | var | void | while |
|  with | abstract | boolean | byte |
|  char | class | const | debugger |
|  double | enum | export | extends |
|  final | float | goto | implements |
|  import | int | interface | long |
|  native | package | private | protected |
|  public | short | static | super |
|  synchronized | throws | transient | volatile |
|  null | true | false |


#### Sources:

[https://www.ecma-international.org/wp-content/uploads/ECMA-262_3rd_edition_december_1999.pdf](https://www.ecma-international.org/wp-content/uploads/ECMA-262_3rd_edition_december_1999.pdf)
### JavaScript (5th edition) (45 keywords)
| | | | |
|---|---|---|---|
|  break | case | catch | continue |
|  debugger | default | delete | do |
|  else | finally | for | function |
|  if | in | instanceof | new |
|  return | switch | this | throw |
|  try | typeof | var | void |
|  while | with | class | const |
|  enum | export | extends | import |
|  super | implements | interface | let |
|  package | private | protected | public |
|  static | yield | null | true |
|  false |


#### Sources:

[https://www.ecma-international.org/wp-content/uploads/ECMA-262_5th_edition_december_2009.pdf](https://www.ecma-international.org/wp-content/uploads/ECMA-262_5th_edition_december_2009.pdf)
### JavaScript (6th edition) (46 keywords)
| | | | |
|---|---|---|---|
|  break | case | catch | class |
|  const | continue | debugger | default |
|  delete | do | else | export |
|  extends | finally | for | function |
|  if | import | in | instanceof |
|  new | return | super | switch |
|  this | throw | try | typeof |
|  var | void | while | with |
|  yield | let | static | enum |
|  await | implements | interface | package |
|  private | protected | public | null |
|  true | false |


#### Sources:

[https://www.ecma-international.org/wp-content/uploads/ECMA-262_5th_edition_december_2009.pdf](https://www.ecma-international.org/wp-content/uploads/ECMA-262_5th_edition_december_2009.pdf)
### Kotlin (1.4) (79 keywords)
| | | | |
|---|---|---|---|
|  as | as? | break | class |
|  continue | do | else | false |
|  for | fun | if | in |
|  !in | interface | is | !is |
|  null | object | package | return |
|  super | this | throw | true |
|  try | typealias | typeof | val |
|  var | when | while | by |
|  catch | constructor | delegate | dynamic |
|  field | file | finally | get |
|  import | init | param | property |
|  receiver | set | setparam | where |
|  actual | abstract | annotation | companion |
|  const | crossinline | data | enum |
|  expect | external | final | infix |
|  inline | inner | internal | lateinit |
|  noinline | open | operator | out |
|  override | private | protected | public |
|  reified | sealed | suspend | tailrec |
|  vararg | field | it |


#### Sources:

[https://kotlinlang.org/docs/reference/keyword-reference.html](https://kotlinlang.org/docs/reference/keyword-reference.html)
### Lua (5.3) (22 keywords)
| | | | |
|---|---|---|---|
|  and | break | do | else |
|  elseif | end | false | for |
|  function | goto | if | in |
|  local | nil | not | or |
|  repeat | return | then | true |
|  until | while |


#### Sources:

[https://www.lua.org/manual/5.3/manual.html#3.1](https://www.lua.org/manual/5.3/manual.html#3.1)
### MATLAB (R2020a) (20 keywords)
| | | | |
|---|---|---|---|
|  break | case | catch | classdef |
|  continue | else | elseif | end |
|  for | function | global | if |
|  otherwise | parfor | persistent | return |
|  spmd | switch | try | while |


#### Sources:

[https://www.mathworks.com/help/matlab/ref/iskeyword.html](https://www.mathworks.com/help/matlab/ref/iskeyword.html)
### Objective-C (2.0) (85 keywords)
| | | | |
|---|---|---|---|
|  asm | auto | break | case |
|  char | const | continue | default |
|  do | double | else | enum |
|  extern | float | for | goto |
|  if | inline | int | long |
|  register | restrict | return | short |
|  signed | sizeof | static | struct |
|  switch | typedef | union | unsigned |
|  void | volatile | while | _Bool |
|  _Complex | __block | Imaginary | id |
|  Class | SEL | IMP | BOOL |
|  nil | Nil | YES | NO |
|  self | super | _cmd | __strong |
|  __weak | __autoreleasing | __unsafe_unretained | oneway |
|  In | out | inout | bycopy |
|  byref | @autoreleasepool | @interface | @implementation |
|  @protocol | @end | @private | @protected |
|  @public | @package | @try | @throw |
|  @catch() | @finally | @property | @synthesize |
|  @dynamic | @class | @selector() | @protocol() |
|  @required | @optional | @encode | @"string" |
|  @synchronized() |


#### Sources:

[https://learning.oreilly.com/library/view/learning-objective-c-20/9780321712110/](https://learning.oreilly.com/library/view/learning-objective-c-20/9780321712110/)
### PHP (7.4) (69 keywords)
| | | | |
|---|---|---|---|
|  __halt_compiler() | abstract | and | array() |
|  as | break | callable | case |
|  catch | class | clone | const |
|  continue | declare | default | die() |
|  do | echo | else | elseif |
|  empty() | enddeclare | endfor | endforeach |
|  endif | endswitch | endwhile | eval() |
|  exit() | extends | final | finally |
|  fn | for | foreach | function |
|  global | goto | if | implements |
|  include | include_once | instanceof | insteadof |
|  interface | isset() | list() | namespace |
|  new | or | print | private |
|  protected | public | require | require_once |
|  return | static | switch | throw |
|  trait | try | unset() | use |
|  var | while | xor | yield |
|  yield from |


#### Sources:

[http://php.net/manual/en/reserved.keywords.php](http://php.net/manual/en/reserved.keywords.php)
### Python 2 (2.7) (31 keywords)
| | | | |
|---|---|---|---|
|  and | as | assert | break |
|  class | continue | def | del |
|  elif | else | except | exec |
|  finally | for | from | global |
|  if | import | in | is |
|  lambda | not | or | pass |
|  print | raise | return | try |
|  while | with | yield |


#### Sources:

[https://docs.python.org/2.7/reference/lexical_analysis.html#keywords](https://docs.python.org/2.7/reference/lexical_analysis.html#keywords)
### Python 3 (3.10) (38 keywords)
| | | | |
|---|---|---|---|
|  False | None | True | and |
|  as | assert | async | await |
|  break | class | continue | def |
|  del | elif | else | except |
|  finally | for | from | global |
|  if | import | in | is |
|  lambda | nonlocal | not | or |
|  pass | raise | return | try |
|  while | with | yield | match |
|  case | _ |


#### Sources:

[https://docs.python.org/3.10/reference/lexical_analysis.html#keywords](https://docs.python.org/3.10/reference/lexical_analysis.html#keywords)
[https://docs.python.org/3.10/reference/lexical_analysis.html#keywords](https://docs.python.org/3.10/reference/lexical_analysis.html#keywords)
### R (4.0) (21 keywords)
| | | | |
|---|---|---|---|
|  ... | ..1 | FALSE | Inf |
|  NA | NA_character_ | NA_complex_ | NA_integer_ |
|  NA_real_ | NaN | NULL | TRUE |
|  break | else | for | function |
|  if | in | next | repeat |
|  while |


#### Sources:

[https://cran.r-project.org/doc/manuals/r-release/R-lang.html#Reserved-words](https://cran.r-project.org/doc/manuals/r-release/R-lang.html#Reserved-words)
### Ruby (2.7) (41 keywords)
| | | | |
|---|---|---|---|
|  _ENCODING_ | _LINE_ | _FILE_ | BEGIN |
|  END | alias | and | begin |
|  break | case | class | def |
|  defined? | do | else | elsif |
|  end | ensure | false | for |
|  if | in | module | next |
|  nil | not | or | redo |
|  rescue | retry | return | self |
|  super | then | true | undef |
|  unless | until | when | while |
|  yield |


#### Sources:

[https://docs.ruby-lang.org/en/2.7.0/keywords_rdoc.html](https://docs.ruby-lang.org/en/2.7.0/keywords_rdoc.html)
### Rust (1.46) (53 keywords)
| | | | |
|---|---|---|---|
|  as | break | const | continue |
|  crate | else | enum | extern |
|  false | fn | for | if |
|  impl | in | let | loop |
|  match | mod | move | mut |
|  pub | ref | return | self |
|  Self | static | struct | super |
|  trait | true | type | unsafe |
|  use | where | while | async |
|  await | dyn | abstract | become |
|  box | do | final | macro |
|  override | priv | typeof | unsized |
|  virtual | yield | try | union |
|  \'static |


#### Sources:

[https://doc.rust-lang.org/grammar.html#keywords](https://doc.rust-lang.org/grammar.html#keywords)
### Scala (2.13) (40 keywords)
| | | | |
|---|---|---|---|
|  abstract | case | catch | class |
|  def | do | else | extends |
|  false | final | finally | for |
|  forSome | if | implicit | import |
|  lazy | macro | match | new |
|  null | object | override | package |
|  private | protected | return | sealed |
|  super | this | throw | trait |
|  try | true | type | val |
|  var | while | with | yield |


#### Sources:

[https://scala-lang.org/files/archive/spec/2.13/01-lexical-syntax.html](https://scala-lang.org/files/archive/spec/2.13/01-lexical-syntax.html)
### Swift (5.3) (97 keywords)
| | | | |
|---|---|---|---|
|  associatedtype | class | deinit | enum |
|  extension | fileprivate | func | import |
|  init | inout | internal | let |
|  open | operator | private | protocol |
|  public | rethrows | static | struct |
|  subscript | typealias | var | break |
|  case | continue | default | defer |
|  do | else | fallthrough | for |
|  guard | if | in | repeat |
|  return | switch | where | while |
|  as | Any | catch | false |
|  is | nil | super | self |
|  Self | throw | throws | true |
|  try | _ | #available | #colorLiteral |
|  #column | #else | #elseif | #endif |
|  #error | #file | #filePath | #fileLiteral |
|  #function | #if | #imageLiteral | #line |
|  #selector | #sourceLocation | #warning | associativity |
|  convenience | dynamic | didSet | final |
|  get | infix | indirect | lazy |
|  left | mutating | none | nonmutating |
|  optional | override | postfix | precedence |
|  prefix | Protocol | required | right |
|  set | Type | unowned | weak |
|  willSet |


#### Sources:

[https://docs.swift.org/swift-book/ReferenceManual/LexicalStructure.html](https://docs.swift.org/swift-book/ReferenceManual/LexicalStructure.html)
### Visual Basic (2019) (217 keywords)
| | | | |
|---|---|---|---|
|  AddHandler | AddressOf | Alias | And |
|  AndAlso | As | Boolean | ByRef |
|  Byte | ByVal | Call | Case |
|  Catch | CBool | CByte | CChar |
|  CDate | CDbl | CDec | Char |
|  CInt | Class | CLng | CObj |
|  Const | Continue | CSByte | CShort |
|  CSng | CStr | CType | CUInt |
|  CULng | CUShort | Date | Decimal |
|  Declare | Default | Delegate | Dim |
|  DirectCast | Do | Double | Each |
|  Else | ElseIf | End | EndIf |
|  Enum | Erase | Error | Event |
|  Exit | False | Finally | For |
|  For Each | Friend | Function | Get |
|  GetType | GetXMLNamespace | Global | GoSub |
|  GoTo | Handles | If | If() |
|  Implements | Imports | In | Inherits |
|  Integer | Interface | Is | IsNot |
|  Let | Lib | Like | Long |
|  Loop | Me | Mod | Module |
|  MustInherit | MustOverride | MyBase | MyClass |
|  NameOf | Namespace | Narrowing | New |
|  Next | Not | Nothing | NotInheritable |
|  NotOverridable | Object | Of | On |
|  Operator | Option | Optional | Or |
|  OrElse | Out | Overloads | Overridable |
|  Overrides | ParamArray | Partial | Private |
|  Property | Protected | Public | RaiseEvent |
|  ReadOnly | ReDim | REM | RemoveHandler |
|  Resume | Return | SByte | Select |
|  Set | Shadows | Shared | Short |
|  Single | Static | Step | Stop |
|  String | Structure | Sub | SyncLock |
|  Then | Throw | To | True |
|  Try | TryCast | TypeOf | UInteger |
|  ULong | UShort | Using | Variant |
|  Wend | When | While | Widening |
|  With | WithEvents | WriteOnly | Xor |
|  #Const | #Else | #ElseIf | #End |
|  #If | = | & | &= |
|  * | *= | / | /= |
|  \ | \= | ^ | ^= |
|  + | += | - | -= |
|  >> | >>= | << | <<= |
|  Aggregate | Ansi | Assembly | Async |
|  Auto | Await | Binary | Compare |
|  Custom | Distinct | Equals | Explicit |
|  From | Group By | Group Join | Into |
|  IsFalse | IsTrue | Iterator | Join |
|  Key | Mid | Off | Order By |
|  Preserve | Skip | Skip While | Strict |
|  Take | Take While | Text | Unicode |
|  Until | Where | Yield | #ExternalSource |
|  #Region |


#### Sources:

[https://docs.microsoft.com/en-us/dotnet/visual-basic/language-reference/keywords/](https://docs.microsoft.com/en-us/dotnet/visual-basic/language-reference/keywords/)

