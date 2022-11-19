package main

import (
    "net/http"
    "fmt"
    "encoding/json"
    "io"
    "os"
    "strconv"
    "os/exec"
    "math/rand"
    "time"
    "strings"
)



type request struct {
	TypeID int `json:"typeID"`
}

type proof struct {
        Context []string `json:"@context"`
        Type string `json:"type"`
        Created string `json:"created"`
        Domain string `json:"domain"`
        Nonce string `json:"nonce"`
        ProofOfPurpose string `json:"proofPurpose"`
        VerificationMethod string `json:"verificationMethod"`
        ProofValue string `json:"proofValue"`
}

type VC struct {
	TypeID int `json:"typeID"`
	Type string `json:"type"`
        ID string `json:"id"`
        Proof proof `json:"proof"`
}

type HolderSignedVC struct {
	VC VC `json`
	Proof proof `json:"proof"`
}



func fetchVC(id int) []byte  {
	content, err := os.ReadFile("./cert/vc_"+strconv.Itoa(id)+".cert")
	if (err != nil) {
		fmt.Println(err)
	}
	return content
}

func createDID(){
	s1 := rand.NewSource(time.Now().UnixNano())
    	r1 := rand.New(s1)
	randomNum := r1.Intn(100)
	refName := "user-key-"+strconv.Itoa(randomNum)
	cmd := exec.Command("algoid", "create", refName)
        out, err := cmd.Output()
        if err != nil {
                fmt.Println(err)
        }
	f, err := os.Create("did.ref")

        if err != nil {
                fmt.Println(err)
        }
        defer f.Close()

        _, err = f.WriteString(refName)
	if err != nil {
                fmt.Println(err)
        }

	//sync or publish
	cmd = exec.Command("algoid", "sync", refName)
        out, err = cmd.Output()
	if err != nil {
                fmt.Println(err)
        }
	fmt.Println(string(out))
}

func createSignature() {
	
	content, err := os.ReadFile("did.ref")
	if (err != nil){
		//create did
		createDID()
		content, err = os.ReadFile("did.ref")
	}
	if (err != nil){
		fmt.Println(err)
	}
	//did.key contains referencename
	referenceName := string(content)
	//get signature using algoid - algoid sign <referencename> -i "vp"
	///command := "algoid sign "+referencename+" -i \"vp\""
	cmd := exec.Command("algoid", "sign", referenceName, "-i", "vp")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.Create("user.sign")
	
    	if err != nil {
        	fmt.Println(err)
    	}
	defer f.Close()
	
	_, err = f.WriteString(string(out))
    	if err != nil {
        	fmt.Println(err)
    	}
}

func signVC(vc VC) HolderSignedVC {
	var signedVC HolderSignedVC
	var sign proof

	content, err := os.ReadFile("user.sign")
	if (err != nil){
		//create signature 
		createSignature()
		content, err = os.ReadFile("user.sign")
	}
	if (err != nil) {
		fmt.Println(err)
	}
	err = json.Unmarshal(content, &sign)
	if (err != nil){
		fmt.Println(err)
	}
	signedVC.VC = vc
	signedVC.Proof = sign
	return signedVC


}

func RequestVP(w http.ResponseWriter, req *http.Request) {
	var r request
	var vcWithoutSign VC
	b, _ := io.ReadAll(req.Body)
	err := json.Unmarshal(b, &r);
	if (err != nil) {
		fmt.Println(err)
	}
	vc := fetchVC(r.TypeID)
	err = json.Unmarshal(vc, &vcWithoutSign)
	//signing
	hsvc := signVC(vcWithoutSign)
	if (err != nil) {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hsvc)
}

func RequestAllVC(w http.ResponseWriter, req *http.Request) {
	certs, err := os.ReadDir("./cert")
	if err != nil {
		fmt.Println(err)
	}
	var vcs  []VC
	for _, cert := range certs {
		var vcWithoutSign VC
		fileName := cert.Name()
		splitFileName := strings.Split(fileName, "_")
		splitFileName = strings.Split(splitFileName[1], ".")
		docID := splitFileName[0]
		docIDInt, _ := strconv.Atoi(docID)
		vc := fetchVC(docIDInt)
		err := json.Unmarshal(vc, &vcWithoutSign)
		if err != nil {
			fmt.Println(err)
		}
		vcs = append(vcs, vcWithoutSign)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vcs)
}

func main() {
    //Handlers
    http.HandleFunc("/idme/holder/request/v1/vp", RequestVP)
    http.HandleFunc("/idme/holder/request/v1/allvc", RequestAllVC)
    fmt.Printf("Running on port 8080...");
    http.ListenAndServe("131.159.209.212:8080", nil)
}
