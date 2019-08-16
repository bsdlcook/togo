# What is togo?

togo is a command-line utility to easily review source code annotations, provided with the following gramatical syntax for comments: ``@Label: Context.``, all of the annotations will be presented to the user for review. 

togo will read the sourcefile and filter out unnecessary code, displaying: sourcefile name, line number of the comment followed by the label and context.

A basic example could be as follows: 

```cpp
// file: example.cpp
#include <string>

class Example
{
public:
  // @Finish: Implement Foo() function.
  static void Foo() {}
  // @Cleanup: Make Bar() function private.
  static void Bar(const std::string& Foobar) { this->_Foobar = Foobar; }

private:
  std::string _Foobar;
};

int
main(void)
{
  // @Cleanup: Use a more appropiate lambda name.
  auto ReallyLongName = []() { return 1; };

  // @Cleanup: Use std::size_t instead.
  for (unsigned int i = 0; i < 10; ++i)
    Example::Bar(std::to_string(i));

  // @Review: Something else here.
  return (0);
}
```

Yielding the output:

```
	example.cpp
(L7) @Finish: Implement Foo() function. 
(L9) @Cleanup: Make Bar() function private. 
(L16) @Cleanup: : : : ; : : ; hello world. 
(L17) @Cleanup: lol. 
(L21) @Cleanup: Use a more appropiate lambda name. 
(L24) @Cleanup: Use std::size_t instead. 
(L28) @Review: Something else here. 
```

Whilst this may seem like a useless program, I have found much use for it despite it's simplity.
