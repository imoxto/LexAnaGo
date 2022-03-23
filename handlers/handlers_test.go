package handlers

import "testing"

func TestValidateAddress(t *testing.T) {
	var email, web bool
	email, web = ValidateAddress("www.9anime.to")
	if email != false || web != true {
		t.Errorf("ValidateAddress(\"www.9anime.to\") = %v, %v ; want = false, true", email, web)
	}

	email, web = ValidateAddress("a1v@gmc.net")
	if email != true || web != false {
		t.Errorf("ValidateAddress(\"a1v@gmc.net\") = %v, %v ; want = true, false", email, web)
	}

	email, web = ValidateAddress("www.co1de.net")
	if email != false || web != true {
		t.Errorf("ValidateAddress(\"www.co1de.net\") = %v, %v ; want = false, true", email, web)
	}

	email, web = ValidateAddress("vip69@mail.gov.bd")
	if email != true || web != false {
		t.Errorf("ValidateAddress(\"vip69@mail.gov.bd\") = %v, %v ; want = true, false", email, web)
	}

	// this should be accepted but the dfa becomes too complex with "(www.)?"
	email, web = ValidateAddress("me.com")
	if email != false || web != false {
		t.Errorf("ValidateAddress(\"me.com\") = %v, %v ; want = false, false", email, web)
	}

	email, web = ValidateAddress("1ct2o")
	if email != false || web != false {
		t.Errorf("ValidateAddress(\"1ct2o\") = %v, %v ; want = false, false", email, web)
	}

	email, web = ValidateAddress("b2dg5s")
	if email != false || web != false {
		t.Errorf("ValidateAddress(\"b2dg5s\") = %v, %v ; want = false, false", email, web)
	}

	email, web = ValidateAddress("abc")
	if email != false || web != false {
		t.Errorf("ValidateAddress(\"b2dg5s\") = %v, %v ; want = false, false", email, web)
	}

}
