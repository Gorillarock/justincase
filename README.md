# JustInCase 
## a toolset for easy regex generation.

### Functions:
GenerateCaseInsensitiveRegex
- takes string input param
- returns *regexp.Regexp (which matches on input word regardless of casing of any letters)


### Example Usage:
<pre><code>
package main

import (
	"fmt"

	justincase "github.com/Gorillarock/justincase"
)

func main() {
	reg := justincase.GenerateCaseInsensitiveRegex("pEopLe")
	fmt.Printf("people == %s (%v)\n", reg, reg.MatchString("people"))
	fmt.Printf("PEOPLE == %s (%v)\n", reg, reg.MatchString("PEOPLE"))
	fmt.Printf("PeOPLe == %s (%v)\n", reg, reg.MatchString("PeOPLe"))
	fmt.Printf("POPLE == %s (%v)\n", reg, reg.MatchString("POPLE"))
}
</code></pre>