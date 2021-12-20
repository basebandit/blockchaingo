package blockchaingo

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

//arbitrary number. Our goal is to have a target that takes less than 256 bits in memory.
const targetBits = 24

//ProofOfWork contains our proof data.Represents a proof-of-work
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

//NewProofOfWork builds and returns a ProofOfWork.
// For our ProofOfWork we are going to use the following algorithm:
// 1.Take some publicly known data (in case of email, it’s receiver’s email address; in case of Bitcoin, it’s block headers).
// 2. Add a counter to it. The counter starts at 0.
// 3. Get a hash of the data + counter combination.
// 4. Check that the hash meets certain requirements.
//		a. If it does, you’re done.
//		b. If it doesn’t, increase the counter and repeat the steps 3 and 4.
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits)) //we use256 bits bcoz the blocks are hashed using SHA-256 algo.

	pow := &ProofOfWork{b, target}

	return pow
}

//merges block fields with the target and nonce. nonce here is the counter from our pow algorithm description.(nonce is a cryptographic term)
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

//Run performs a proof-of-work.First, we initialize variables: hashInt is the integer representation of hash; nonce is the counter.
//Next, we run an “infinite” loop: it’s limited by maxNonce, which equals to math.MaxInt64; this is done to avoid a possible overflow of nonce.Although the difficuly of this PoW implementation is too low for the counter to overflow.
//In the loop we:
// 1.Prepare data.
// 2.Hash it with SHA-256.
// 3.Convert the hash to a big integer.
// 4.Compare the integer with the target.

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("\n\n")

	return nonce, hash[:]
}

//Validate validates proof-of-work
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
