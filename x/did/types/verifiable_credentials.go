package types

/*
This represents a minimal verifiable credential
definition: https://www.w3.org/TR/vc-data-model/#zero-knowledge-proofs
{
  "@context": [
    "https://www.w3.org/2018/credentials/v1",
    "https://www.w3.org/2018/credentials/examples/v1"
  ],
  "type": ["VerifiableCredential", "UniversityDegreeCredential"],
  "credentialSchema": {
    "id": "did:example:cdf:35LB7w9ueWbagPL94T9bMLtyXDj9pX5o",
    "type": "did:example:schema:22KpkXgecryx9k7N6XN1QoN3gXwBkSU8SfyyYQG"
  },
  "issuer": "did:example:Wz4eUg7SetGfaUVCn8U9d62oDYrUJLuUtcy619",
  "credentialSubject": {
    "givenName": "Jane",
    "familyName": "Doe",
    "degree": {
      "type": "BachelorDegree",
      "name": "Bachelor of Science and Arts",
      "college": "College of Engineering"
    }
  },
  "proof": {
    "type": "CLSignature2019",
    "issuerData": "5NQ4TgzNfSQxoLzf2d5AV3JNiCdMaTgm...BXiX5UggB381QU7ZCgqWivUmy4D",
    "attributes": "pPYmqDvwwWBDPNykXVrBtKdsJDeZUGFA...tTERiLqsZ5oxCoCSodPQaggkDJy",
    "signature": "8eGWSiTiWtEA8WnBwX4T259STpxpRKuk...kpFnikqqSP3GMW7mVxC4chxFhVs",
    "signatureCorrectnessProof": "SNQbW3u1QV5q89qhxA1xyVqFa6jCrKwv...dsRypyuGGK3RhhBUvH1tPEL8orH"
  }
}
*/

type VerifiableCredential struct {
	Context           string            `json:"@context"`
	ID                string            `json:"id"`
	Type              string            `json:"type"`
	Issuer            string            `json:"issuer"`
	CredentialSubject CredentialSubject `json:"credentialsubject"`
	Proof             Proof             `json:"proof"`
}

func NewVerifiableCredential(context string, id string, vctype string, issuer string, credentialSubject CredentialSubject, proof Proof) VerifiableCredential {
	return VerifiableCredential{
		Context:           context,
		ID:                id,
		Type:              vctype,
		Issuer:            issuer,
		CredentialSubject: credentialSubject,
		Proof:             proof,
	}
}

type CredentialSubject struct {
	Role       string `json:"role"`
	IsVerified bool   `json:"isverified"`
}

func NewCredentialSubject(role string, isVerified bool) CredentialSubject {
	return CredentialSubject{
		Role:       role,
		IsVerified: isVerified,
	}
}

type Proof struct {
	Type                      string `json:"type"`
	IssuerData                string `json:"issuerdata"`
	Attributes                string `json:"attributes"`
	Signature                 string `json:"signature"`
	SignatureCorrectnessProof string `json:"signatureproof"`
}

func NewProof(ptype string, issuerData string, attributes string, signature string, signatureProof string) Proof {
	return Proof{
		Type:                      ptype,
		IssuerData:                issuerData,
		Attributes:                attributes,
		Signature:                 signature,
		SignatureCorrectnessProof: signatureProof,
	}
}
