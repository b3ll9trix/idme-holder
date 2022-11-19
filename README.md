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
  
  "type": "passport",
  "id": "P87237292",
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





