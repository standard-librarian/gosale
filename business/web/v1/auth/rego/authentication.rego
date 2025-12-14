package mdht.rego

default auth := false

auth if {
	jwt_valid
}

jwt_valid := valid if {
	[valid, header, payload] := verify_jwt
}

verify_jwt := io.jwt.decode_verify(input.Token, {"cert": input.Key})
