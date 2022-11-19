#REST APIs for the app IDme!

### There are 3 actors:
1. Holder 
2. Issuer
3. Verifier

### 1. Holder
The Holder is the user whose ID issued by the Issuer and verified by the Verifier

### 1.1 APIs

 GET `/idme/holder/request/v1/vc`

### SUCCESS Response
Code: `200 OK` \
Body: 
```json
{
  
  "@context": [
    "https://www.w3.org/2018/credentials/v1",
    "https://www.w3.org/2018/credentials/examples/v1"
  ],
  
  "id": "http://example.edu/credentials/1872",
  
  "type": ["VerifiableCredential", "AlumniCredential"],
  
  "issuer": "https://example.edu/issuers/565049",
  
  "issuanceDate": "2010-01-01T19:23:24Z",
  
  "credentialSubject": {
    
    "id": "did:example:ebfeb1f712ebc6f1c276e12ec21",
    
    "alumniOf": {
      "id": "did:example:c276e12ec21ebfeb1f712ebc6f1",
      "name": [{
        "value": "Example University",
        "lang": "en"
      }, {
        "value": "Exemple d'Université",
        "lang": "fr"
      }]
    }
  },
  
  
  "proof": {
  "@context": [
      "https://w3id.org/security/v1"
    ],
    "type": "Ed25519Signature2020",
    "created": "2022-11-19T11:12:00Z",
    "domain": "did.algorand.foundation",
    "nonce": "5f8b64ef46fcdc7b5072257834336d30",
    "proofPurpose": "authentication",
    "verificationMethod": "did:algo:f08114f0-0438-466a-8044-6a45bbaf7ec9#master",
    "proofValue": "ut8ASooBKV98LQukfEmrykSHWzkLx5NMzSOWlQ0RSMcrCGu0mbNPOduCfTZ5mA5L90UJYplPpRk2VdoCk14HBw=="

  }
}
```





