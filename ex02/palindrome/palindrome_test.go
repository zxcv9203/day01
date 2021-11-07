package palindrome

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMemUse(t *testing.T) {
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(1))
	var st, ed runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&st)
	Palindrome("asdoifhjroeijgoejrgioejrgioerjgeriogjoeirgjeoirgjeriog", "dqwdqwd", "", "안녕하세요 정말 반가워요")
	runtime.ReadMemStats(&ed)
	alloc := ed.TotalAlloc - st.TotalAlloc
	limit := uint64(64 * 1000)
	fmt.Println(alloc, limit)
	if alloc > limit {
		t.Errorf("Runtime Error : too much using memory")
	}
}

func ExamplePalindrome_done() {
	Palindrome("level")
	// Output: level 5
}

func ExamplePalindrome_noString() {
	Palindrome("")
	//Output:
}

func ExamplePalindrome_alpha() {
	Palindrome("asdqw")
	//Output: asdqwqdsa 9
}

func ExamplePalindrome_bigger() {
	Palindrome("fasuiofhroawiefghiouetrhgojfkahfklrgaelfgyaweriugheruihgieurghfgsdfsdf")
	//Output: fasuiofhroawiefghiouetrhgojfkahfklrgaelfgyaweriugheruihgieurghfgsdfsdfdsfdsgfhgrueighiurehguirewaygfleagrlkfhakfjoghrteuoihgfeiwaorhfoiusaf 139
}
func ExamplePalindrome_kr() {
	Palindrome("안녕")
	//Output:안녕안 3
}
func ExamplePalindrome_multiCase() {
	Palindrome("asdasd", "level", "", "안녕")
	//Output:asdasdsadsa 11
	//level 5
	//안녕안 3
}
