### Preface

togo is a command-line utility to easily review source code annotations, provided with the following gramatical syntax for comments: ``@Label: Context.``, all of the annotations will be presented to the user for review. 

togo will read the sourcefile and filter out unnecessary code, displaying: sourcefile name, line number of the comment followed by the label and context.

### Example

A basic example sourcefile could be as follows: 

```cpp
// file: example.cpp
#include <string>
#include <map> // @Cleanup: Remove this! We don't use maps.

// @Cleanup: Use a more appropiate class-name according to the Google guidelines.
// See here: https://google.github.io/styleguide/cppguide.html
// We need to be more consistent throughout our codebase.
class Example {
	public:
		// @Finish: Implement Foo() function.
		static void Foo() {}
		// @Comment: Here is a brief description of the issue.
		// Furthermore, we can add more context here if needed
		// to convey a better understanding of the issue.
		static void Bar() {} 
		// @Cleanup: Ideally, use std::size_t instead.
		static unsigned int Something() { return 1; }
		// @Hack: Rewrite this function, it uses *too* many hacks.
		static void Hacky() {}
		// @Review: Cleanup on isle 3, please check this.
};

int main(void) {
	// @Finish: Do something here.
	return (0);
}
```

Producing the following output:

![](https://s.wired.sh/misc/togo.png)

Whilst this may seem like a useless program, I have found much use for it despite it's simplity.
