package creational

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Received cache was nill")
	}
	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}
	if item1 == whitePrototype {
		t.Error("item1 cannot be equal to white prototype")
	}

	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("Type assertion not successfully")
	}
	shirt1.SKU = "abc"

	item2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal("Type ")
	}

	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt not successfully")
	}
	if shirt1.SKU == shirt2.SKU {
		t.Error("SKU's of shirt1 and shirt2 must be different")
	}
	if shirt1 == shirt2 {
		t.Error("Shirt 1 cannot be equal to Shirt 2")
	}
}
